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

func (r *Repository) CreateCompetitorMatch(ctx context.Context, competitorMatchInfoDAO *dao.CreateCompetitorMatchDAOReq) (string, error) {
	competitorMatchInfoDAO.SetTimeStamp()

	result, err := r.competitorMatchColl.InsertOne(ctx, competitorMatchInfoDAO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", fmt.Errorf("%w: error duplicate key for competitorMatch: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error competitorMatch scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting competitorMatch: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
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

// func (r *Repository) UpdateCompetitorMatch(ctx context.Context, competitorMatchID string, competitorMatchInfoDAO *dao.UpdateCompetitorMatchDAOReq) error {
// 	competitorMatchOID, err := r.ConvertToObjectID(competitorMatchID)
// 	if err != nil {
// 		return err
// 	}

// 	competitorMatchInfoDAO.RenewUpdate()

// 	filter := bson.M{"_id": *competitorMatchOID}
// 	update, err := api_assets.StructToBsonMap(competitorMatchInfoDAO)
// 	if err != nil {
// 		return err
// 	}

// 	result, err := r.competitorMatchColl.UpdateOne(
// 		ctx,
// 		filter,
// 		bson.M{"$set": update},
// 	)
// 	if err != nil {
// 		return fmt.Errorf("error updating competitorMatch: %w", err)
// 	}

// 	if result.MatchedCount == 0 {
// 		return fmt.Errorf("%w: no competitorMatch found with id: %s", customerrors.ErrNotFound, competitorMatchID)
// 	}

// 	return nil
// }

func (r *Repository) DeleteCompetitorMatch(ctx context.Context, competitorMatchID string) error {
	err := r.setDeletedAt(ctx, r.competitorMatchColl, competitorMatchID, "competitorMatch")
	if err != nil {
		return err
	}

	return nil
}
