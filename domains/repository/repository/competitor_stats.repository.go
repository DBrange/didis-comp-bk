package repository

import (
	"context"
	"fmt"
	"time"

	competitor_stats_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/competitor_stats/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateCompetitorStats(ctx context.Context, competitorOID *primitive.ObjectID) error {
	competitorStatsDAO := &competitor_stats_dao.CreateCompetitorStatsDAOReq{}

	competitorStatsDAO.CompetitorID = *competitorOID
	competitorStatsDAO.Matches = []primitive.ObjectID{}
	competitorStatsDAO.TournamentsWon = []primitive.ObjectID{}

	competitorStatsDAO.SetTimeStamp()

	_, err := r.competitorStatsColl.InsertOne(ctx, competitorStatsDAO)
	if err != nil {
		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return fmt.Errorf("%w: error competitorStats scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return fmt.Errorf("error when inserting competitorStats: %w", err)
	}

	return nil
}

func (r *Repository) GetCompetitorStatsByID(ctx context.Context, competitorStatsID string) (*competitor_stats_dao.GetCompetitorStatsByIDDAORes, error) {
	var competitorStats competitor_stats_dao.GetCompetitorStatsByIDDAORes

	competitorStatsOID, err := r.ConvertToObjectID(competitorStatsID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *competitorStatsOID}

	err = r.competitorStatsColl.FindOne(ctx, filter).Decode(&competitorStats)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for competitorStats: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the competitorStats: %w", err)
	}

	return &competitorStats, nil
}

// func (r *Repository) UpdateCompetitorStats(ctx context.Context, competitorStatsID string, competitorStatsInfoDAO *competitor_stats_dao.UpdateCompetitorStatsDAOReq) error {
// 	competitorStatsOID, err := r.ConvertToObjectID(competitorStatsID)
// 	if err != nil {
// 		return err
// 	}

// 	competitorStatsInfoDAO.RenewUpdate()

// 	filter := bson.M{"_id": *competitorStatsOID}
// 	update, err := api_assets.StructToBsonMap(competitorStatsInfoDAO)
// 	if err != nil {
// 		return err
// 	}

// 	result, err := r.competitorStatsColl.UpdateOne(
// 		ctx,
// 		filter,
// 		bson.M{"$set": update},
// 	)
// 	if err != nil {
// 		return fmt.Errorf("error updating competitorStats: %w", err)
// 	}

// 	if result.MatchedCount == 0 {
// 		return fmt.Errorf("%w: no competitorStats found with id: %s", customerrors.ErrNotFound, competitorStatsID)
// 	}

// 	return nil
// }

func (r *Repository) AddMatchInCompetitorStats(ctx context.Context, competitorOID, matchOID *primitive.ObjectID) error {
	if competitorOID == nil{
		return nil
	}

	filter := bson.M{"competitor_id": competitorOID}
	update := bson.M{
		"$push": bson.M{"matches": matchOID},
		"$set":  bson.M{"updated_at": time.Now().UTC()},
	}

	result, err := r.competitorStatsColl.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return fmt.Errorf("error updating competitorStats: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no competitorStats found with competitor_id: %s", customerrors.ErrNotFound, competitorOID.Hex())
	}

	return nil
}

func (r *Repository) UpdateCompetitorStats(ctx context.Context, competitorOID *primitive.ObjectID, winner bool) error {
	incrementField := "total_wins"
	if !winner {
		incrementField = "total_losses"
	}

	filter := bson.M{"competitor_id": competitorOID}
	update := bson.M{
		"$inc": bson.M{incrementField: 1},
		"$set": bson.M{"updated_at": time.Now().UTC()},
	}

	result, err := r.competitorStatsColl.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return fmt.Errorf("error updating competitorStats: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no competitorStats found with competitor_id: %s", customerrors.ErrNotFound, competitorOID.Hex())
	}

	return nil
}

func (r *Repository) DeleteCompetitorStats(ctx context.Context, competitorStatsID string) error {
	err := r.SetDeletedAt(ctx, r.competitorStatsColl, competitorStatsID, "competitorStats")
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) AddTournamentWonInCompetitorStats(ctx context.Context, competitorOID, tournamentOID *primitive.ObjectID) error {
	filter := bson.M{"competitor_id": competitorOID}

	update := bson.M{"$push": bson.M{"tournaments_won": tournamentOID}}

	result, err := r.competitorStatsColl.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return fmt.Errorf("error updating competitorStats: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no competitorStats found with competitor_id: %s", customerrors.ErrNotFound, competitorOID.Hex())
	}

	return nil
}

func (r *Repository) AddPrizeInMultipleCompetitorStats(ctx context.Context, competitorOIDs []*primitive.ObjectID, prize float64) error {
	// Filtro para seleccionar los documentos que coincidan con los IDs de los competidores
	filter := bson.M{"competitor_id": bson.M{"$in": competitorOIDs}}

	// Operación de actualización para incrementar el campo total_prize
	update := bson.M{
		"$inc": bson.M{
			"money_earned": prize,
		},
	}

	// Ejecutar la actualización en la colección
	_, err := r.competitorStatsColl.UpdateMany(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("error updating total_prize: %v", err)
	}

	return nil
}

func (r *Repository) RemoveMultipleCompetitorStatsMatches(ctx context.Context, competitorOIDs, matchesToRemove []*primitive.ObjectID) error {
	filter := bson.M{"competitor_id": bson.M{"$in": competitorOIDs}}
		
		cursor, err := r.competitorStatsColl.Find(ctx, filter)
		if err != nil {
			return fmt.Errorf("error retrieving category registrations: %w", err)
		}
		defer cursor.Close(ctx)
		
		for cursor.Next(ctx) {
			var result struct {
				Matches      []*primitive.ObjectID `bson:"matches"`
				CompetitorID *primitive.ObjectID `bson:"competitor_id"`
			}
			
			if err := cursor.Decode(&result); err != nil {
				return fmt.Errorf("error decoding competitor: %w", err)
			}
			
		competitorFilter := bson.M{"competitor_id": result.CompetitorID}

		update := bson.M{
			"$pull": bson.M{
				"matches": bson.M{
					"$in": matchesToRemove,
				},
			},
		}

		_, err := r.competitorStatsColl.UpdateOne(ctx, competitorFilter, update)
		if err != nil {
			return fmt.Errorf("error updating competitor %v: %w", result.CompetitorID, err)
		}
	}

	if err := cursor.Err(); err != nil {
		return fmt.Errorf("error iterating cursor: %w", err)
	}

	return nil

}