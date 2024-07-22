package repository

import (
	"context"
	"fmt"

	api_assets "github.com/DBrange/didis-comp-bk/cmd/api/assets"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/competitor_stats/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateCompetitorStats(ctx context.Context, competitorStatsInfoDAO *dao.CreateCompetitorStatsDAOReq) (string, error) {
	competitorStatsInfoDAO.SetTimeStamp()

	result, err := r.competitorStatsColl.InsertOne(ctx, competitorStatsInfoDAO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", fmt.Errorf("%w: error duplicate key for competitorStats: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error competitorStats scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting competitorStats: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetCompetitorStatsByID(ctx context.Context, competitorStatsID string) (*dao.GetCompetitorStatsByIDDAORes, error) {
	var competitorStats dao.GetCompetitorStatsByIDDAORes

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

func (r *Repository) UpdateCompetitorStats(ctx context.Context, competitorStatsID string, competitorStatsInfoDAO *dao.UpdateCompetitorStatsDAOReq) error {
	competitorStatsOID, err := r.ConvertToObjectID(competitorStatsID)
	if err != nil {
		return err
	}

	competitorStatsInfoDAO.RenewUpdate()

	filter := bson.M{"_id": *competitorStatsOID}
	update, err := api_assets.StructToBsonMap(competitorStatsInfoDAO)
	if err != nil {
		return err
	}

	result, err := r.competitorStatsColl.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": update},
	)
	if err != nil {
		return fmt.Errorf("error updating competitorStats: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no competitorStats found with id: %s", customerrors.ErrNotFound, competitorStatsID)
	}

	return nil
}

func (r *Repository) DeleteCompetitorStats(ctx context.Context, competitorStatsID string) error {
	err := r.setDeletedAt(ctx, r.competitorStatsColl, competitorStatsID, "competitorStats")
	if err != nil {
		return err
	}

	return nil
}
