package repository

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	api_assets "github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/match/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) CreateMatch(ctx context.Context, matchDAO *dao.CreateMatchDAOReq) (string, error) {
	matchDAO.SetTimeStamp()

	result, err := r.matchColl.InsertOne(ctx, matchDAO)
	if err != nil {
		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error match scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting match: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetMatchByID(ctx context.Context, matchID string) (*dao.GetMatchByIDDAORes, error) {
	var match dao.GetMatchByIDDAORes

	matchOID, err := r.ConvertToObjectID(matchID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *matchOID}

	err = r.matchColl.FindOne(ctx, filter).Decode(&match)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for match: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the match: %w", err)
	}

	return &match, nil
}

func (r *Repository) FindMatchID(ctx context.Context, position int, roundOID *primitive.ObjectID) (string, error) {
	var findID struct {
		ID *primitive.ObjectID `bson:"_id"`
	}

	filter := bson.M{"round_id": roundOID, "position": position}

	opts := options.FindOne().SetProjection(bson.M{"_id": 1})

	err := r.matchColl.FindOne(ctx, filter, opts).Decode(&findID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", fmt.Errorf("%w: error when searching for match: %s", customerrors.ErrNotFound, err.Error())
		}
		return "", fmt.Errorf("error when searching for the match: %w", err)
	}

	return findID.ID.Hex(), nil
}

func (r *Repository) UpdateMatch(ctx context.Context, matchID string, matchInfoDAO *dao.UpdateMatchDAOReq) error {
	matchOID, err := r.ConvertToObjectID(matchID)
	if err != nil {
		return err
	}

	matchInfoDAO.RenewUpdate()

	filter := bson.M{"_id": *matchOID}
	update, err := api_assets.StructToBsonMap(matchInfoDAO)
	if err != nil {
		return err
	}

	result, err := r.matchColl.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": update},
	)
	if err != nil {
		return fmt.Errorf("error updating match: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no match found with id: %s", customerrors.ErrNotFound, matchID)
	}

	return nil
}

func (r *Repository) DeleteMatch(ctx context.Context, matchID string) error {
	err := r.SetDeletedAt(ctx, r.matchColl, matchID, "match")
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) VerifyMatchExists(ctx context.Context, matchOID *primitive.ObjectID) error {
	var result struct{}

	filter := bson.M{"_id": matchOID}

	opts := options.FindOne().SetProjection(bson.M{"_id": 1})

	err := r.matchColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("%w: error when searching for match: %s", customerrors.ErrNotFound, err.Error())
		}
		return fmt.Errorf("error when searching for the match: %w", err)
	}

	return nil
}

func (r *Repository) VerifyMatchesExist(ctx context.Context, matchOIDs []*primitive.ObjectID) error {
	if len(matchOIDs) == 0 {
		return nil // No hay competidores para verificar
	}
	fmt.Printf("%+v", matchOIDs)
	filter := bson.M{"_id": bson.M{"$in": matchOIDs}}
	opts := options.Find().SetProjection(bson.M{"_id": 1})

	cursor, err := r.matchColl.Find(ctx, filter, opts)
	if err != nil {
		return fmt.Errorf("error when searching for matches: %w", err)
	}
	defer cursor.Close(ctx)

	var foundIDs []*primitive.ObjectID
	for cursor.Next(ctx) {
		var result struct {
			ID *primitive.ObjectID `bson:"_id"`
		}
		if err := cursor.Decode(&result); err != nil {
			return fmt.Errorf("error decoding match: %w", err)
		}
		foundIDs = append(foundIDs, result.ID)
	}

	if err := cursor.Err(); err != nil {
		return fmt.Errorf("error iterating cursor: %w", err)
	}

	if len(foundIDs) != len(matchOIDs) {
		missingIDs := r.getMissingIDs(matchOIDs, foundIDs)
		return fmt.Errorf("%w: the following matches were not found: %v", customerrors.ErrNotFound, missingIDs)
	}

	return nil
}

func (r *Repository) GetPositionsBracketMatch(ctx context.Context, roundOID *primitive.ObjectID, organizeType models.ORGANIZE_TYPE, organizeBracket models.ORGANIZE_BRACKET) ([]*dao.GetPositionsBracketMatchDAORes, error) {
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"round_id": roundOID}}},
		bson.D{{Key: "$lookup", Value: buildLookupStage("competitor_matches", "_id", "match_id", "competitor_match")}},
		bson.D{{Key: "$unwind", Value: "$competitor_match"}},
		bson.D{{Key: "$lookup", Value: buildLookupStage("competitor_registrations", "competitor_match.competitor_id", "competitor_id", "competitor_registration")}},
		bson.D{{Key: "$unwind", Value: "$competitor_registration"}},
		bson.D{{Key: "$group", Value: bson.M{
			"_id":            "$_id",
			"match_position": bson.M{"$first": "$position"},
			"competitors": bson.M{"$push": bson.M{
				"_id":              "$competitor_registration.competitor_id",
				"position":         "$competitor_match.position",
				"current_position": "$competitor_registration.position",
			}},
		}}},
	}

	var positions []*dao.GetPositionsBracketMatchDAORes
	cursor, err := r.matchColl.Aggregate(ctx, pipeline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for match: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the match: %w", err)
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &positions); err != nil {
		return nil, fmt.Errorf("error when decoding match: %w", err)
	}

	return positions, nil
}

func (r *Repository) SetWinnerInMatch(ctx context.Context, matchOID, competitorOID *primitive.ObjectID, matchResult string) error {
	filter := bson.M{"_id": matchOID}

	update := bson.M{"$set": bson.M{"winner": competitorOID, "result": matchResult}}

	result, err := r.matchColl.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return fmt.Errorf("error updating match: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no match found with id: %s or there is already a winner", customerrors.ErrNotFound, matchOID.Hex())
	}

	return nil
}

func (r *Repository) VerifyMatchesInRoundExits(ctx context.Context, roundOID *primitive.ObjectID) (bool, error) {
	var result struct{}

	filter := bson.M{"round_id": roundOID}

	opts := options.FindOne().SetProjection(bson.M{"_id": 1})

	err := r.matchColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Si no se encuentran documentos, retorna false y nil, ya que no es un error en sí.
			return false, nil
		}
		// Para otros errores, retorna false y el error.
		return false, fmt.Errorf("error when searching for the match: %w", err)
	}

	// Si se encuentra un documento, retorna true y nil.
	return true, nil
}

func (r *Repository) VerifyMatchPosition(ctx context.Context, matchOID *primitive.ObjectID, position int) error {
	var result struct{}

	filter := bson.M{"_id": matchOID, "position": position}

	opts := options.FindOne().SetProjection(bson.M{"_id": 1})

	err := r.matchColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("%w: error when searching for match: %s", customerrors.ErrNotFound, err.Error())
		}
		return fmt.Errorf("error when searching for the match: %w", err)
	}

	return nil
}

func (r *Repository) GetRoundQuantityMatches(ctx context.Context, roundOID *primitive.ObjectID) (int, error) {
	pipeline := mongo.Pipeline{
		// Filtrar los documentos por el round_id dado
		bson.D{{Key: "$match", Value: bson.M{"round_id": roundOID}}},

		// Agrupar por round_id y contar la cantidad de documentos en cada grupo
		bson.D{{Key: "$group", Value: bson.M{
			"_id":   "$round_id",       // Agrupar por round_id
			"count": bson.M{"$sum": 1}, // Contar el número de documentos
		}}},
	}

	cursor, err := r.matchColl.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, fmt.Errorf("error executing aggregation: %w", err)
	}
	defer cursor.Close(ctx)

	var result struct {
		Count int `bson:"count"`
	}

	// Si hay resultados, decodificar y retornar el conteo
	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return 0, fmt.Errorf("error decoding result: %w", err)
		}
		return result.Count, nil
	}

	// Si no hay resultados, significa que no hay documentos con el round_id dado
	return 0, nil
}

func (r *Repository) GetMatchPosition(ctx context.Context, matchOID *primitive.ObjectID) (int, error) {
	filter := bson.M{"_id": matchOID}

	var position struct {
		Position int `bson:"position"`
	}

	opts := options.FindOne().SetProjection(bson.M{"position": 1})

	err := r.matchColl.FindOne(ctx, filter, opts).Decode(&position)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, fmt.Errorf("%w: error when searching for match: %s", customerrors.ErrNotFound, err.Error())
		}
		return 0, fmt.Errorf("error when searching for the match: %w", err)
	}

	return position.Position, nil
}

func (r *Repository) VerifyMatchAndRoundCoincidence(ctx context.Context, matchOID, roundOID *primitive.ObjectID, round models.ROUND) error {
	// Definir el pipeline
	pipeline := bson.A{
		// Coincidencia inicial en la colección de "match" para verificar el matchOID y roundOID
		bson.M{"$match": bson.M{"_id": matchOID, "round_id": roundOID}},

		// Lookup para unir con la colección de "round" basada en roundOID
		bson.M{
			"$lookup": bson.M{
				"from":         "rounds", // Nombre de la colección de "rounds"
				"localField":   "round_id",
				"foreignField": "_id",
				"as":           "round",
			},
		},

		// Desenredar el arreglo resultante del lookup
		bson.M{"$unwind": "$round"},

		// Filtrar para asegurarse de que el nombre del round coincida
		bson.M{"$match": bson.M{"round.round": round}},
	}

	// Ejecutar el pipeline
	cursor, err := r.matchColl.Aggregate(ctx, pipeline)
	if err != nil {
		return fmt.Errorf("error running aggregation: %w", err)
	}
	defer cursor.Close(ctx)

	// Verificar si se encontró algún resultado
	if !cursor.Next(ctx) {
		return fmt.Errorf("%w: error when searching for match or round", customerrors.ErrNotFound)
	}

	// Si no hay más errores, todo está en orden
	return nil
}

func (r *Repository) DeleteMultipleMatches(ctx context.Context, matchesToRemove []*primitive.ObjectID) error {
	filter := bson.M{"_id": bson.M{"$in": matchesToRemove}}

	result, err := r.matchColl.DeleteMany(ctx, filter)
	if err != nil {
		return fmt.Errorf("%w: error deleting matches: %s", customerrors.ErrDeleted, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("%w: no matches found with the provided ids", customerrors.ErrNotFound)
	}

	return nil
}
