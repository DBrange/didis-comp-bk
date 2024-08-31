package repository

import (
	"context"
	"fmt"

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
		fmt.Printf("por aca %v", err)
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
		"_id":    "$$match._id",
		"result": "$$match.result",
		"winner": "$$match.winner",
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
		fmt.Printf("por aca %v", err)
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("%w: error when searching for round: %s", customerrors.ErrNotFound, err.Error())
		}
		return fmt.Errorf("error when searching for the round: %w", err)
	}

	return nil
}

// func (r *Repository) UpdateCompetitorsPointsAndMoneyEarned(ctx context.Context, tournamentOID *primitive.ObjectID, tournamentWithCategory bool) error {
// 	pipeline := mongo.Pipeline{
// 		// Filtrar por el torneo
// 		bson.D{{Key: "$match", Value: bson.M{"tournament_id": tournamentOID}}},
// 		// Buscar las rondas
// 		bson.D{{Key: "$lookup", Value: bson.M{
// 			"from":         "matches",
// 			"localField":   "_id",
// 			"foreignField": "round_id",
// 			"as":           "match",
// 		}}},
// 		bson.D{{Key: "$unwind", Value: bson.M{
// 			"path":                       "$match",
// 			"preserveNullAndEmptyArrays": true, // para manejar documentos que podrían no tener coincidencias
// 		}}},
// 		// Buscar los competitor_matches
// 		bson.D{{Key: "$lookup", Value: bson.M{
// 			"from":         "competitor_matches",
// 			"localField":   "match._id",
// 			"foreignField": "match_id",
// 			"as":           "competitor_match",
// 		}}},
// 		bson.D{{Key: "$unwind", Value: bson.M{
// 			"path":                       "$competitor_match",
// 			"preserveNullAndEmptyArrays": true,
// 		}}},
// 		// Agrupar por ronda para obtener IDs de competidores, puntos y premios
// 		bson.D{{Key: "$group", Value: bson.M{
// 			"_id":            "$_id",
// 			"competitor_ids": bson.M{"$push": "$competitor_match.competitor_id"},
// 			"points":         bson.M{"$first": "$points"},
// 			"total_prize":    bson.M{"$first": "$total_prize"},
// 		}}},
// 		// Buscar en competitor_stats
// 		bson.D{{Key: "$lookup", Value: bson.M{
// 			"from":         "competitor_stats",
// 			"localField":   "competitor_ids",
// 			"foreignField": "competitor_id",
// 			"as":           "competitor_stats",
// 		}}},
// 		bson.D{{Key: "$unwind", Value: bson.M{
// 			"path":                       "$competitor_stats",
// 			"preserveNullAndEmptyArrays": true,
// 		}}},
// 		// Sumar el total del premio al dinero ganado
// 		bson.D{{Key: "$set", Value: bson.M{
// 			"competitor_stats.money_earned": bson.M{"$add": bson.A{"$competitor_stats.money_earned", "$total_prize"}},
// 		}}},
// 	}

// 	if tournamentWithCategory {
// 		pipeline = append(pipeline,
// 			// Buscar en category_registrations
// 			bson.D{{Key: "$lookup", Value: bson.M{
// 				"from":         "category_registrations",
// 				"localField":   "competitor_ids",
// 				"foreignField": "competitor_id",
// 				"as":           "category_registration",
// 			}}},
// 			bson.D{{Key: "$unwind", Value: bson.M{
// 				"path":                       "$category_registration",
// 				"preserveNullAndEmptyArrays": true,
// 			}}},
// 			// Sumar los puntos de la categoría
// 			bson.D{{Key: "$set", Value: bson.M{
// 				"category_registration.points": bson.M{"$add": bson.A{"$category_registration.points", "$points"}},
// 			}}})
// 	}

// 	cursor, err := r.roundColl.Aggregate(ctx, pipeline)
// 	if err != nil {
// 		return fmt.Errorf("error al ejecutar el pipeline de agregación: %v", err)
// 	}
// 	defer cursor.Close(ctx)

// 	// Iterar sobre los resultados y actualizar las colecciones
// 	for cursor.Next(ctx) {
// 		var result struct {
// 			CompetitorStats struct {
// 				ID          primitive.ObjectID `bson:"_id"`
// 				MoneyEarned float64            `bson:"money_earned"`
// 			} `bson:"competitor_stats"`
// 			CategoryRegistration struct {
// 				ID     primitive.ObjectID `bson:"_id"`
// 				Points int                `bson:"points"`
// 			} `bson:"category_registration"`
// 		}

// 		if err := cursor.Decode(&result); err != nil {
// 			return fmt.Errorf("error al decodificar el resultado: %v", err)
// 		}

// 		// Actualizar competitor_stats
// 		_, err = r.competitorStatsColl.UpdateOne(
// 			ctx,
// 			bson.M{"_id": result.CompetitorStats.ID},
// 			bson.M{"$set": bson.M{"money_earned": result.CompetitorStats.MoneyEarned}},
// 		)
// 		if err != nil {
// 			return fmt.Errorf("error al actualizar competitor_stats: %v", err)
// 		}

// 		// Actualizar category_registrations si corresponde
// 		if result.CategoryRegistration.ID != primitive.NilObjectID {
// 			_, err = r.categoryRegistrationColl.UpdateOne(
// 				ctx,
// 				bson.M{"_id": result.CategoryRegistration.ID},
// 				bson.M{"$set": bson.M{"points": result.CategoryRegistration.Points}},
// 			)
// 			if err != nil {
// 				return fmt.Errorf("error al actualizar category_registrations: %v", err)
// 			}
// 		}
// 	}

// 	if err := cursor.Err(); err != nil {
// 		return fmt.Errorf("error al iterar sobre los resultados: %v", err)
// 	}

// 	return nil
// }
