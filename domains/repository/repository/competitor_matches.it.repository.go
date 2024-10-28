package repository

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/competitor_match/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateCompetitorMatch(ctx context.Context, competitorMatchDAO *dao.CreateCompetitorMatchDAOReq) error {
	competitorMatchDAO.SetTimeStamp()

	_, err := r.competitorMatchColl.InsertOne(ctx, competitorMatchDAO)
	if err != nil {
		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return fmt.Errorf("%w: error competitorMatch scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return fmt.Errorf("error when inserting competitorMatch: %w", err)
	}

	return nil
}

func (r *Repository) GetCompetitorMatchByID(ctx context.Context, competitorMatchID string) (*dao.GetCompetitorMatchByIDDAORes, error) {
	var competitorMatch dao.GetCompetitorMatchByIDDAORes

	competitorMatchOID, err := r.ConvertToObjectID(competitorMatchID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *competitorMatchOID}

	err = r.competitorMatchColl.FindOne(ctx, filter).Decode(&competitorMatch)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for competitorMatch: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the competitorMatch: %w", err)
	}

	return &competitorMatch, nil
}

func (r *Repository) UpdateCompetitorMatch(ctx context.Context, matchOID *primitive.ObjectID, competitorMatchDAO *dao.UpdateCompetitorMatchDAOReq) error {
	competitorMatchDAO.RenewUpdate()

	filter := bson.M{"match_id": matchOID, "position": competitorMatchDAO.Position}

	update := bson.M{"$set": competitorMatchDAO}

	result, err := r.competitorMatchColl.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return fmt.Errorf("error updating competitorMatch: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no competitorMatch found with id: %s", customerrors.ErrNotFound, matchOID.Hex())
	}

	return nil
}

func (r *Repository) DeleteCompetitorMatch(ctx context.Context, competitorMatchID string) error {
	err := r.SetDeletedAt(ctx, r.competitorMatchColl, competitorMatchID, "competitorMatch")
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateMultipleCompetitorMatches(ctx context.Context, competitorMatchDAOs []*dao.UpdateCompetitorMatchDAOReq) error {
	if len(competitorMatchDAOs) == 0 {
		return nil // No hay nada que actualizar
	}
	
	// Crear una operación de escritura para cada actualización
	var operations []mongo.WriteModel
	for _, dao := range competitorMatchDAOs {
		dao.RenewUpdate()
		filter := bson.M{"match_id": dao.MatchID, "position": dao.Position}
		update := bson.M{"$set": dao}
		operation := mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update)
		operations = append(operations, operation)
	}

	// Ejecutar todas las operaciones de escritura en una sola llamada a la base de datos
	result, err := r.competitorMatchColl.BulkWrite(ctx, operations)
	if err != nil {
		return fmt.Errorf("error updating competitorMatches: %w", err)
	}

	if result.MatchedCount != result.ModifiedCount {
		return fmt.Errorf("mismatch in update counts: matched %d, modified %d", result.MatchedCount, result.ModifiedCount)
	}
	fmt.Println("llegue aca al menos")
	return nil
}

func (r *Repository) SetCompetitorInNextMatch(ctx context.Context, matchOID, competitorOID *primitive.ObjectID, position int) error {
	filter := bson.M{"match_id": matchOID, "position": position}

	update := bson.M{"$set": bson.M{"competitor_id": competitorOID}}

	result, err := r.competitorMatchColl.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return fmt.Errorf("error updating match: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no match found with id: %s", customerrors.ErrNotFound, matchOID.Hex())
	}

	return nil
}

func (r *Repository) VerifyCompetitorsMatch(ctx context.Context, matchOID, competitorOID *primitive.ObjectID) error {
	filterLosser := bson.M{"match_id": matchOID, "competitor_id": competitorOID}

	var competitor struct{}

	err := r.competitorMatchColl.FindOne(ctx, filterLosser).Decode(&competitor)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("%w: no match found for losser competitor", customerrors.ErrNotFound)
		}
		return fmt.Errorf("error when searching for losser competitor: %w", err)
	}

	return nil
}

func (r *Repository) DeleteMultipleCompetitorMatches(ctx context.Context, matchesToRemove []*primitive.ObjectID) error {
	filter := bson.M{"match_id": bson.M{"$in": matchesToRemove}}
	
	result, err := r.competitorMatchColl.DeleteMany(ctx, filter)
	if err != nil {
		return fmt.Errorf("%w: error deleting matches: %s", customerrors.ErrDeleted, err.Error())
	}
	
	if result.DeletedCount == 0 {
		return fmt.Errorf("%w: no matches found with the provided ids", customerrors.ErrNotFound)
	}

	return nil
}


func (r *Repository) GetCompetitorIDsFromMatches(ctx context.Context, matches []*primitive.ObjectID) ([]*primitive.ObjectID, error) {
	// Si no se pasan matches, retorna un slice vacío
	if len(matches) == 0 {
		return []*primitive.ObjectID{}, nil
	}

	// Crear un filtro para buscar todos los competitorIDs de los matches proporcionados
	filter := bson.M{"match_id": bson.M{"$in": matches}}

	// Realizar la consulta a la colección
	cursor, err := r.competitorMatchColl.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error retrieving competitor IDs from matches: %w", err)
	}
	defer cursor.Close(ctx)

	var competitorIDs []*primitive.ObjectID

	// Iterar sobre el cursor para decodificar los resultados
	for cursor.Next(ctx) {
		var match struct {
			CompetitorID primitive.ObjectID `bson:"competitor_id"`
		}

		if err := cursor.Decode(&match); err != nil {
			return nil, fmt.Errorf("error decoding competitor ID: %w", err)
		}

		// Agregar el CompetitorID al slice
		competitorIDs = append(competitorIDs, &match.CompetitorID)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error iterating cursor: %w", err)
	}

	return competitorIDs, nil
}

func (r *Repository) GetCompetitorIDByMatchAndPosition(ctx context.Context, matchID *primitive.ObjectID, position int) (*primitive.ObjectID, error) {
	// Validar que el matchID no sea nil
	if matchID == nil {
		return nil, fmt.Errorf("matchID cannot be nil")
	}

	// Crear un filtro para buscar el competitorID por matchID y position
	filter := bson.M{
		"match_id": matchID,
		"position": position,
	}

	// Realizar la consulta a la colección
	var match struct {
		CompetitorID primitive.ObjectID `bson:"competitor_id"`
	}

	err := r.competitorMatchColl.FindOne(ctx, filter).Decode(&match)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no competitor found for match_id: %s and position: %d", matchID.Hex(), position)
		}
		return nil, fmt.Errorf("error retrieving competitor ID: %w", err)
	}

	return &match.CompetitorID, nil
}

