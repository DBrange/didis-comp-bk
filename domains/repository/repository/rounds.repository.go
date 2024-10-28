package repository

import (
	"context"
	"fmt"
	"sort"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/round/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) CreateRound(ctx context.Context, roundDAO *dao.CreateRoundDAOReq) (string, error) {
	roundDAO.SetTimeStamp()

	result, err := r.roundColl.InsertOne(ctx, roundDAO)
	if err != nil {
		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error round scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting round: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetRoundByID(ctx context.Context, roundID string) (*dao.GetRoundByIDDAORes, error) {
	var round dao.GetRoundByIDDAORes

	roundOID, err := r.ConvertToObjectID(roundID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *roundOID}

	err = r.roundColl.FindOne(ctx, filter).Decode(&round)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for round: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the round: %w", err)
	}

	return &round, nil
}

func (r *Repository) UpdateRoundTotalPrize(ctx context.Context, roundOID *primitive.ObjectID, totalPrize float64) error {
	filter := bson.M{"_id": roundOID}

	update := bson.M{"$set": bson.M{"total_prize": totalPrize}}

	result, err := r.roundColl.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("error updating round: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no round found with id: %s", customerrors.ErrNotFound, roundOID.Hex())
	}

	return nil
}

func (r *Repository) UpdateRoundPoints(ctx context.Context, roundOID *primitive.ObjectID, points int) error {
	filter := bson.M{"_id": roundOID}

	update := bson.M{"$set": bson.M{"points": points}}

	result, err := r.roundColl.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("error updating round: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no round found with id: %s", customerrors.ErrNotFound, roundOID.Hex())
	}

	return nil
}

func (r *Repository) VerifyRoundExists(ctx context.Context, roundOID *primitive.ObjectID) error {
	var result struct{}

	filter := bson.M{"_id": roundOID}

	opts := options.FindOne().SetProjection(bson.M{"_id": 1})

	err := r.roundColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("%w: error when searching for round: %s", customerrors.ErrNotFound, err.Error())
		}
		return fmt.Errorf("error when searching for the round: %w", err)
	}

	return nil
}

func (r *Repository) GetRoundWithMatches(ctx context.Context, roundOID, categoryOID *primitive.ObjectID) (*dao.GetRoundWithMatchesDAORes, error) {
	pipeline := buildRoundWithMatchesPipeline(roundOID, categoryOID)

	var result dao.GetRoundWithMatchesDAORes
	cursor, err := r.roundColl.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error executing aggregate pipeline: %w", err)
	}
	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("error decoding result: %w", err)
		}
	} else {
		return nil, fmt.Errorf("round not found: %s", roundOID.Hex())
	}

	return processMatches(result), nil
}

func buildRoundWithMatchesPipeline(roundOID, categoryOID *primitive.ObjectID) mongo.Pipeline {
	return mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"_id": roundOID}}},
		bson.D{{Key: "$lookup", Value: buildLookupStage("matches", "_id", "round_id", "matches")}},
		bson.D{{Key: "$lookup", Value: buildLookupStage("competitor_matches", "matches._id", "match_id", "competitor_matches")}},
		bson.D{{Key: "$lookup", Value: buildLookupStage("competitor_users", "competitor_matches.competitor_id", "competitor_id", "competitor_users")}},
		bson.D{{Key: "$lookup", Value: buildLookupStage("users", "competitor_users.user_id", "_id", "users")}},
		bson.D{{Key: "$lookup", Value: buildLookupStage("guest_competitors", "competitor_matches.competitor_id", "competitor_id", "guest_competitors")}},
		bson.D{{Key: "$lookup", Value: buildLookupStage("guest_users", "guest_competitors.guest_user_id", "_id", "guest_users")}},
		bson.D{{Key: "$lookup", Value: buildLookupStage("category_registrations", "competitor_matches.competitor_id", "competitor_id", "category_registration")}},
		bson.D{{Key: "$project", Value: buildProjectStage(categoryOID)}},
	}
}

func buildLookupStage(from, localField, foreignField, as string) bson.D {
	return bson.D{{Key: "from", Value: from}, {Key: "localField", Value: localField}, {Key: "foreignField", Value: foreignField}, {Key: "as", Value: as}}
}

func buildProjectStage(categoryOID *primitive.ObjectID) bson.M {
	return bson.M{
		"_id":         1,
		"round":       1,
		"total_prize": 1,
		"points":      1,
		"matches": bson.M{
			"$map": bson.M{
				"input": "$matches",
				"as":    "match",
				"in":    buildMatchProjection(categoryOID),
			},
		},
	}
}

func buildMatchProjection(categoryOID *primitive.ObjectID) bson.M {
	return bson.M{
		"_id":      "$$match._id",
		"result":   "$$match.result",
		"winner":   "$$match.winner",
		"date":     "$$match.date",
		"position": "$$match.position",
		"competitors": bson.M{
			"$map": bson.M{
				"input": bson.M{
					"$filter": bson.M{
						"input": "$competitor_matches",
						"as":    "cm",
						"cond":  bson.M{"$eq": bson.A{"$$cm.match_id", "$$match._id"}},
					},
				},
				"as": "cm",
				"in": buildCompetitorProjection(categoryOID),
			},
		},
	}
}

func buildCompetitorProjection(categoryOID *primitive.ObjectID) bson.M {
	return bson.M{
		"_id":              "$$cm.competitor_id",
		"current_position": buildCurrentPositionProjection(categoryOID),
		"position":         "$$cm.position",
		"users":            buildUsersProjection(),
		"guest_users":      buildGuestUsersProjection(),
	}
}

func buildCurrentPositionProjection(categoryOID *primitive.ObjectID) bson.M {
	return bson.M{
		"$let": bson.M{
			"vars": bson.M{
				"cr": bson.M{
					"$filter": bson.M{
						"input": "$category_registration",
						"as":    "cr",
						"cond": bson.M{
							"$and": bson.A{bson.M{"$eq": bson.A{"$$cr.competitor_id", "$$cm.competitor_id"}},
								bson.M{"$eq": bson.A{"$$cr.category_id", categoryOID}}}},
					},
				},
			},
			"in": bson.M{
				"$cond": bson.A{
					bson.M{"$gt": bson.A{bson.M{"$size": "$$cr"}, 0}},
					bson.M{"$arrayElemAt": bson.A{"$$cr.current_position", 0}},
					nil,
				},
			},
		},
	}
}

func buildUsersProjection() bson.M {
	return bson.M{
		"$filter": bson.M{
			"input": "$users",
			"as":    "user",
			"cond":  bson.M{"$in": bson.A{"$$user._id", buildUserIDArray()}},
		},
	}
}

func buildGuestUsersProjection() bson.M {
	return bson.M{
		"$filter": bson.M{
			"input": "$guest_users",
			"as":    "guest",
			"cond":  bson.M{"$in": bson.A{"$$guest._id", buildGuestUserIDArray()}},
		},
	}
}

func buildUserIDArray() bson.M {
	return bson.M{
		"$map": bson.M{
			"input": bson.M{
				"$filter": bson.M{
					"input": "$competitor_users",
					"as":    "cu",
					"cond":  bson.M{"$eq": bson.A{"$$cu.competitor_id", "$$cm.competitor_id"}},
				},
			},
			"as": "cu",
			"in": "$$cu.user_id",
		},
	}
}

func buildGuestUserIDArray() bson.M {
	return bson.M{
		"$map": bson.M{
			"input": bson.M{
				"$filter": bson.M{
					"input": "$guest_competitors",
					"as":    "gc",
					"cond":  bson.M{"$eq": bson.A{"$$gc.competitor_id", "$$cm.competitor_id"}},
				},
			},
			"as": "gc",
			"in": "$$gc.guest_user_id",
		},
	}
}

func processMatches(result dao.GetRoundWithMatchesDAORes) *dao.GetRoundWithMatchesDAORes {
	for i, match := range result.Matches {
		// Ordenar los competidores por su campo "position"
		sort.Slice(result.Matches[i].Competitors, func(a, b int) bool {
			return result.Matches[i].Competitors[a].Position < result.Matches[i].Competitors[b].Position
		})

		for j, competitor := range match.Competitors {
			if len(competitor.Users) > 0 {
				result.Matches[i].Competitors[j].Users = convertToUserDAORes(competitor.Users)
				result.Matches[i].Competitors[j].GuestUsers = []*dao.GetRoundWithMatchesUserDAORes{} // Slice vacío
			} else if len(competitor.GuestUsers) > 0 {
				result.Matches[i].Competitors[j].GuestUsers = convertToUserDAORes(competitor.GuestUsers)
				result.Matches[i].Competitors[j].Users = []*dao.GetRoundWithMatchesUserDAORes{} // Slice vacío
			} else {
				// Inicializar ambos como slices vacíos si no hay datos
				result.Matches[i].Competitors[j].Users = []*dao.GetRoundWithMatchesUserDAORes{}
				result.Matches[i].Competitors[j].GuestUsers = []*dao.GetRoundWithMatchesUserDAORes{}
			}
		}

		if match.Winner != nil {
			result.Matches[i].PositionWinner = findCompetitorPosition(result.Matches[i].Competitors, *match.Winner)
		}
	}

	return &result
}

func findCompetitorPosition(competitors []*dao.GetRoundWithMatchesCompetitorDAORes, winnerID primitive.ObjectID) *int {
	for _, competitor := range competitors {
		if competitor.ID != nil && *competitor.ID == winnerID {
			return &competitor.Position
		}
	}
	return nil
}

func convertToUserDAORes(users []*dao.GetRoundWithMatchesUserDAORes) []*dao.GetRoundWithMatchesUserDAORes {
	userDAORes := make([]*dao.GetRoundWithMatchesUserDAORes, len(users))
	for i, user := range users {
		userDAORes[i] = &dao.GetRoundWithMatchesUserDAORes{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Image:     user.Image,
		}
	}
	return userDAORes
}

func (r *Repository) GetRoundsWithCompetitors(ctx context.Context, tournamentOID *primitive.ObjectID) ([]*dao.GetRoundWithCompetitorsDAORes, error) {
	pipeline := mongo.Pipeline{
		// Filtrar por el torneo
		bson.D{{Key: "$match", Value: bson.M{"tournament_id": tournamentOID}}},
		// Buscar las rondas
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "matches",
			"localField":   "_id",
			"foreignField": "round_id",
			"as":           "match",
		}}},
		bson.D{{Key: "$unwind", Value: bson.M{
			"path":                       "$match",
			"preserveNullAndEmptyArrays": true, // para manejar documentos que podrían no tener coincidencias
		}}},
		// Buscar los competitor_matches
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "competitor_matches",
			"localField":   "match._id",
			"foreignField": "match_id",
			"as":           "competitor_match",
		}}},
		bson.D{{Key: "$unwind", Value: bson.M{
			"path":                       "$competitor_match",
			"preserveNullAndEmptyArrays": true,
		}}},
		// Agrupar por ronda para obtener el ID, total_prize, points, y competitor_ids
		bson.D{{Key: "$group", Value: bson.M{
			"_id":            "$_id",
			"competitor_ids": bson.M{"$addToSet": "$competitor_match.competitor_id"}, // Usamos $addToSet para evitar duplicados
			"total_prize":    bson.M{"$first": "$total_prize"},
			"points":         bson.M{"$first": "$points"},
		}}},
	}

	cursor, err := r.roundColl.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar el pipeline de agregación: %v", err)
	}
	defer cursor.Close(ctx)

	var roundsWithCompetitors []*dao.GetRoundWithCompetitorsDAORes

	// Iterar sobre los resultados y decodificarlos en una estructura
	for cursor.Next(ctx) {
		var result *dao.GetRoundWithCompetitorsDAORes
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("error al decodificar el resultado: %v", err)
		}
		roundsWithCompetitors = append(roundsWithCompetitors, result)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar sobre los resultados: %v", err)
	}

	return roundsWithCompetitors, nil
}

// func (r *Repository) GetDoubleElimMatches(ctx context.Context, tournamentOID *primitive.ObjectID) ([]*dao.GetRoundWithCompetitorsDAORes, error) {
// 	pipeline := mongo.Pipeline{
// 		// Filtrar por el torneo
// 		bson.D{{Key: "$match", Value: bson.M{"_id": tournamentOID}}},
// 		// Buscar el double_elimination_id
// 		bson.D{{Key: "$lookup", Value: bson.M{
// 			"from":         "double_eliminations",
// 			"localField":   "double_elimination_id",
// 			"foreignField": "_id",
// 			"as":           "double_elim",
// 		}}},
// 		bson.D{{Key: "$unwind", Value: bson.M{
// 			"path":                       "$double_elim",
// 			"preserveNullAndEmptyArrays": true,
// 		}}},
// 		// Buscar los matches en double_eliminations
// 		bson.D{{Key: "$lookup", Value: bson.M{
// 			"from":         "matches",
// 			"localField":   "double_elim.rounds",
// 			"foreignField": "_id",
// 			"as":           "double_elim_matches",
// 		}}},
// 		bson.D{{Key: "$unwind", Value: bson.M{
// 			"path":                       "$double_elim_matches",
// 			"preserveNullAndEmptyArrays": true,
// 		}}},
// 		// Agrupar para obtener la lista de matches
// 		bson.D{{Key: "$group", Value: bson.M{
// 			"_id":                   "$double_elim._id",
// 			"double_elim_matches":  bson.M{"$addToSet": "$double_elim_matches"},
// 		}}},
// 	}

// 	cursor, err := r.tournamentColl.Aggregate(ctx, pipeline)
// 	if err != nil {
// 		return nil, fmt.Errorf("error al ejecutar el pipeline de agregación: %v", err)
// 	}
// 	defer cursor.Close(ctx)

// 	var matches []*dao.GetRoundWithCompetitorsDAORes

// 	// Iterar sobre los resultados y decodificarlos en una estructura
// 	for cursor.Next(ctx) {
// 		var result struct {
// 			DoubleElimMatches []*dao.GetRoundWithCompetitorsDAORes `bson:"double_elim_matches"`
// 		}
// 		if err := cursor.Decode(&result); err != nil {
// 			return nil, fmt.Errorf("error al decodificar el resultado: %v", err)
// 		}
// 		matches = append(matches, result.DoubleElimMatches...)
// 	}

// 	if err := cursor.Err(); err != nil {
// 		return nil, fmt.Errorf("error al iterar sobre los resultados: %v", err)
// 	}

// 	return matches, nil
// }

func (r *Repository) VerifyRoundInTournament(ctx context.Context, roundOID, tournamentOID *primitive.ObjectID) error {
	var result struct{}

	filter := bson.M{"_id": roundOID, "tournament_id": tournamentOID}

	opts := options.FindOne().SetProjection(bson.M{"_id": 1})

	err := r.roundColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("%w: error when searching for round: %s", customerrors.ErrNotFound, err.Error())
		}
		return fmt.Errorf("error when searching for the round: %w", err)
	}

	return nil
}

func buildCompetitorsPipeline(roundOID, categoryOID *primitive.ObjectID) mongo.Pipeline {
	return mongo.Pipeline{
		// Initial match to find the specific round
		bson.D{{Key: "$match", Value: bson.M{"_id": roundOID}}},

		// Lookup competitors directly
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "competitors"},
			{Key: "localField", Value: "competitor_ids"}, // Cambia esto al campo correcto en tu documento de round
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "competitors"},
		}}},

		// Unwind competitors array
		bson.D{{Key: "$unwind", Value: "$competitors"}},

		// Lookup competitor users
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "competitor_users"},
			{Key: "localField", Value: "competitors._id"}, // Cambia esto al campo correcto en tus documentos de competidores
			{Key: "foreignField", Value: "competitor_id"},
			{Key: "as", Value: "competitor_users"},
		}}},

		// Lookup regular users
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "users"},
			{Key: "localField", Value: "competitor_users.user_id"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "users"},
		}}},

		// Lookup guest competitors
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "guest_competitors"},
			{Key: "localField", Value: "competitors._id"},
			{Key: "foreignField", Value: "competitor_id"},
			{Key: "as", Value: "guest_competitors"},
		}}},

		// Lookup guest users
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "guest_users"},
			{Key: "localField", Value: "guest_competitors.guest_user_id"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "guest_users"},
		}}},

		bson.D{{Key: "$lookup", Value: buildLookupStage("category_registrations", "competitor_users.competitor_id", "competitor_id", "category_registration")}},

		// Project final competitor structure
		bson.D{{Key: "$project", Value: bson.M{
			"_id":              "$competitors._id",
			"current_position": nil,
			"position":         "$competitors.position",
			"users":            "$users",
			"guest_users":      "$guest_users",
		}}},
	}
}

func (r *Repository) GetRoundGroups(ctx context.Context, roundOID, categoryOID *primitive.ObjectID) (*dao.GetRoundGroupsDAORes, error) {
	pipeline := buildRoundGroupsPipeline(roundOID, categoryOID)

	var result dao.GetRoundGroupsDAORes
	cursor, err := r.roundColl.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error executing aggregate pipeline: %w", err)
	}
	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("error decoding result: %w", err)
		}
	} else {
		return nil, fmt.Errorf("round not found: %s", roundOID.Hex())
	}

	return processGroups(result), nil
}

func buildGroupMatchProjection(categoryOID *primitive.ObjectID) bson.M {
	return bson.M{
		"_id":      "$$match._id",
		"result":   "$$match.result",
		"winner":   "$$match.winner",
		"date":     "$$match.date",
		"position": "$$match.position",
		"competitors": bson.M{
			"$map": bson.M{
				"input": bson.M{
					"$filter": bson.M{
						"input": "$competitor_matches",
						"as":    "cm",
						"cond":  bson.M{"$eq": bson.A{"$$cm.match_id", "$$match._id"}},
					},
				},
				"as": "cm",
				"in": buildCompetitorProjection(categoryOID),
			},
		},
	}
}

func processGroups(result dao.GetRoundGroupsDAORes) *dao.GetRoundGroupsDAORes {
	// Sort groups by position
	sort.Slice(result.Groups, func(i, j int) bool {
		return result.Groups[i].Position < result.Groups[j].Position
	})

	// Process matches and competitors within each group
	for i, group := range result.Groups {
		// Sort matches by position
		sort.Slice(result.Groups[i].Matches, func(a, b int) bool {
			return result.Groups[i].Matches[a].Position < result.Groups[i].Matches[b].Position
		})

		// Process matches
		for j, match := range group.Matches {
			// Sort competitors by position
			sort.Slice(result.Groups[i].Matches[j].Competitors, func(a, b int) bool {
				return result.Groups[i].Matches[j].Competitors[a].Position < result.Groups[i].Matches[j].Competitors[b].Position
			})

			processCompetitors(result.Groups[i].Matches[j].Competitors)

			if match.Winner != nil {
				result.Groups[i].Matches[j].PositionWinner = findCompetitorPosition(result.Groups[i].Matches[j].Competitors, *match.Winner)
			}
		}

		// Process group competitors
		processCompetitorsWithStats(result.Groups[i].Competitors)
	}

	return &result
}

func processCompetitors(competitors []*dao.GetRoundWithMatchesCompetitorDAORes) {
	for k, competitor := range competitors {
		if len(competitor.Users) > 0 {
			competitors[k].Users = convertToUserDAORes(competitor.Users)
			competitors[k].GuestUsers = []*dao.GetRoundWithMatchesUserDAORes{}
		} else if len(competitor.GuestUsers) > 0 {
			competitors[k].GuestUsers = convertToUserDAORes(competitor.GuestUsers)
			competitors[k].Users = []*dao.GetRoundWithMatchesUserDAORes{}
		} else {
			competitors[k].Users = []*dao.GetRoundWithMatchesUserDAORes{}
			competitors[k].GuestUsers = []*dao.GetRoundWithMatchesUserDAORes{}
		}
	}
}

func processCompetitorsWithStats(competitors []*dao.GetRoundGroupCompetitorWithStatsDAORes) {
	for k, competitor := range competitors {
		if len(competitor.Users) > 0 {
			competitors[k].Users = convertToUserDAORes(competitor.Users)
			competitors[k].GuestUsers = []*dao.GetRoundWithMatchesUserDAORes{}
		} else if len(competitor.GuestUsers) > 0 {
			competitors[k].GuestUsers = convertToUserDAORes(competitor.GuestUsers)
			competitors[k].Users = []*dao.GetRoundWithMatchesUserDAORes{}
		} else {
			competitors[k].Users = []*dao.GetRoundWithMatchesUserDAORes{}
			competitors[k].GuestUsers = []*dao.GetRoundWithMatchesUserDAORes{}
		}
	}
}

func buildRoundGroupsPipeline(roundOID, categoryOID *primitive.ObjectID) mongo.Pipeline {
	return mongo.Pipeline{
		// Coincidir el ID de la ronda
		bson.D{{Key: "$match", Value: bson.M{"_id": roundOID}}},

		// Lookup para el torneo
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "tournaments"},
			{Key: "localField", Value: "tournament_id"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "tournament"},
		}}},

		// Desenrollar el torneo ya que solo esperamos uno
		bson.D{{Key: "$unwind", Value: "$tournament"}},

		// Lookup para los grupos del torneo
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "tournament_groups"},
			{Key: "localField", Value: "tournament.groups"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "groups"},
		}}},

		// Lookup para los competidores en los grupos
		bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "competitors"},
				{Key: "let", Value: bson.D{
					{Key: "groupCompetitors", Value: bson.M{
						"$reduce": bson.M{
							"input":        "$groups.competitors",
							"initialValue": bson.A{},
							"in": bson.M{
								"$concatArrays": bson.A{
									"$$value",
									bson.M{
										"$map": bson.M{
											"input": "$$this",
											"as":    "comp",
											"in":    "$$comp.competitor_id", // Assuming this should point to the ID of the competitor
										},
									},
								},
							},
						},
					}},
				}},
				{Key: "pipeline", Value: bson.A{
					bson.D{{Key: "$match", Value: bson.D{
						{Key: "$expr", Value: bson.D{
							{Key: "$in", Value: bson.A{"$_id", "$$groupCompetitors"}},
						}},
					}}},
				}},
				{Key: "as", Value: "all_group_competitors"},
			}},
		},

		// Lookup para las partidas de cada grupo
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "matches"},
			{Key: "localField", Value: "groups.matches"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "all_matches"},
		}}},

		// Lookup para los competidores en las partidas
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "competitor_matches"},
			{Key: "localField", Value: "all_matches._id"},
			{Key: "foreignField", Value: "match_id"},
			{Key: "as", Value: "competitor_matches"},
		}}},

		// Lookup para usuarios competidores
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "competitor_users"},
			{Key: "let", Value: bson.D{
				{Key: "matchCompetitors", Value: "$competitor_matches.competitor_id"},
				{Key: "groupCompetitors", Value: "$all_group_competitors._id"},
			}},
			{Key: "pipeline", Value: bson.A{
				bson.D{{Key: "$match", Value: bson.D{
					{Key: "$expr", Value: bson.D{
						{Key: "$or", Value: bson.A{
							bson.D{{Key: "$in", Value: bson.A{"$competitor_id", "$$matchCompetitors"}}},
							bson.D{{Key: "$in", Value: bson.A{"$competitor_id", "$$groupCompetitors"}}},
						}},
					}},
				}}},
			}},
			{Key: "as", Value: "competitor_users"},
		}}},

		// Lookup para los usuarios regulares
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "users"},
			{Key: "localField", Value: "competitor_users.user_id"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "users"},
		}}},

		// Lookup para los competidores invitados
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "guest_competitors"},
			{Key: "let", Value: bson.D{
				{Key: "matchCompetitors", Value: "$competitor_matches.competitor_id"},
				{Key: "groupCompetitors", Value: "$all_group_competitors._id"},
			}},
			{Key: "pipeline", Value: bson.A{
				bson.D{{Key: "$match", Value: bson.D{
					{Key: "$expr", Value: bson.D{
						{Key: "$or", Value: bson.A{
							bson.D{{Key: "$in", Value: bson.A{"$competitor_id", "$$matchCompetitors"}}},
							bson.D{{Key: "$in", Value: bson.A{"$competitor_id", "$$groupCompetitors"}}},
						}},
					}},
				}}},
			}},
			{Key: "as", Value: "guest_competitors"},
		}}},

		// Lookup para los usuarios invitados
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "guest_users"},
			{Key: "localField", Value: "guest_competitors.guest_user_id"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "guest_users"},
		}}},

		// Lookup para registros de categorías
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "category_registrations"},
			{Key: "let", Value: bson.D{
				{Key: "matchCompetitors", Value: "$competitor_matches.competitor_id"},
				{Key: "groupCompetitors", Value: "$all_group_competitors._id"},
			}},
			{Key: "pipeline", Value: bson.A{
				bson.D{{Key: "$match", Value: bson.D{
					{Key: "$expr", Value: bson.D{
						{Key: "$or", Value: bson.A{
							bson.D{{Key: "$in", Value: bson.A{"$competitor_id", "$$matchCompetitors"}}},
							bson.D{{Key: "$in", Value: bson.A{"$competitor_id", "$$groupCompetitors"}}},
						}},
					}},
				}}},
			}},
			{Key: "as", Value: "category_registration"},
		}}},

		// Proyección final de la estructura
		bson.D{{Key: "$project", Value: buildGroupProjectStage(categoryOID)}},
	}
}

func buildGroupProjectStage(categoryOID *primitive.ObjectID) bson.M {
	return bson.M{
		"_id":              1,
		"round":            1,
		"total_prize":      1,
		"points":           1,
		"total_classified": 1,
		"best_third":       1,
		"groups": bson.M{
			"$map": bson.M{
				"input": "$groups",
				"as":    "group",
				"in": bson.M{
					"_id":      "$$group._id",
					"position": "$$group.position",
					"matches": bson.M{
						"$map": bson.M{
							"input": bson.M{
								"$filter": bson.M{
									"input": "$all_matches",
									"as":    "match",
									"cond": bson.M{
										"$in": bson.A{"$$match._id", "$$group.matches"},
									},
								},
							},
							"as": "match",
							"in": buildGroupMatchProjection(categoryOID),
						},
					},
					"competitors": bson.M{
						"$map": bson.M{
							"input": bson.M{
								"$filter": bson.M{
									"input": "$all_group_competitors",
									"as":    "competitor",
									"cond": bson.M{
										"$in": bson.A{
											"$$competitor._id",
											bson.M{
												"$map": bson.M{
													"input": "$$group.competitors",
													"as":    "groupComp",
													"in":    "$$groupComp.competitor_id",
												},
											},
										},
									},
								},
							},
							"as": "competitor",
							"in": bson.M{
								"_id":              "$$competitor._id",
								"current_position": buildCurrentPositionProjectiontt(categoryOID),
								"position":         0,
								"users":            buildUsersProjectiontt(),
								"guest_users":      buildGuestUsersProjectiontt(),
								"stats": bson.M{
									"$let": bson.M{
										"vars": bson.M{
											"competitorStats": bson.M{
												"$filter": bson.M{
													"input": "$$group.competitors",
													"as":    "gc",
													"cond": bson.M{
														"$eq": bson.A{"$$gc.competitor_id", "$$competitor._id"},
													},
												},
											},
										},
										"in": bson.M{
											"$arrayElemAt": bson.A{"$$competitorStats.stats", 0},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func buildCurrentPositionProjectiontt(categoryOID *primitive.ObjectID) bson.M {
	return bson.M{
		"$let": bson.M{
			"vars": bson.M{
				"cr": bson.M{
					"$filter": bson.M{
						"input": "$category_registration",
						"as":    "cr",
						"cond": bson.M{
							"$and": bson.A{bson.M{"$eq": bson.A{"$$cr.competitor_id", "$$competitor._id"}},
								bson.M{"$eq": bson.A{"$$cr.category_id", categoryOID}}}},
					},
				},
			},
			"in": bson.M{
				"$cond": bson.A{
					bson.M{"$gt": bson.A{bson.M{"$size": "$$cr"}, 0}},
					bson.M{"$arrayElemAt": bson.A{"$$cr.current_position", 0}},
					nil,
				},
			},
		},
	}
}

func buildUsersProjectiontt() bson.M {
	return bson.M{
		"$filter": bson.M{
			"input": "$users",
			"as":    "user",
			"cond":  bson.M{"$in": bson.A{"$$user._id", buildUserIDArraytt()}},
		},
	}
}

func buildGuestUsersProjectiontt() bson.M {
	return bson.M{
		"$filter": bson.M{
			"input": "$guest_users",
			"as":    "guest",
			"cond":  bson.M{"$in": bson.A{"$$guest._id", buildGuestUserIDArraytt()}},
		},
	}
}

func buildUserIDArraytt() bson.M {
	return bson.M{
		"$map": bson.M{
			"input": bson.M{
				"$filter": bson.M{
					"input": "$competitor_users",
					"as":    "cu",
					"cond":  bson.M{"$eq": bson.A{"$$cu.competitor_id", "$$competitor._id"}},
				},
			},
			"as": "cu",
			"in": "$$cu.user_id",
		},
	}
}

func buildGuestUserIDArraytt() bson.M {
	return bson.M{
		"$map": bson.M{
			"input": bson.M{
				"$filter": bson.M{
					"input": "$guest_competitors",
					"as":    "gc",
					"cond":  bson.M{"$eq": bson.A{"$$gc.competitor_id", "$$competitor._id"}},
				},
			},
			"as": "gc",
			"in": "$$gc.guest_user_id",
		},
	}
}
