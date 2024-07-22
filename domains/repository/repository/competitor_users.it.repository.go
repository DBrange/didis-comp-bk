package repository

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/competitor_user/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateCompetitorUser(ctx context.Context, competitorUserInfoDAO *dao.CreateCompetitorUserDAOReq) (string, error) {
	competitorUserInfoDAO.SetTimeStamp()

	result, err := r.competitorUserColl.InsertOne(ctx, competitorUserInfoDAO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", fmt.Errorf("%w: error duplicate key for competitorUser: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error competitorUser scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting competitorUser: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetCompetitorUserByID(ctx context.Context, competitorUserID string) (*dao.GetCompetitorUserByIDDAORes, error) {
	var competitorUser dao.GetCompetitorUserByIDDAORes

	competitorUserOID, err := r.ConvertToObjectID(competitorUserID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *competitorUserOID}

	err = r.competitorUserColl.FindOne(ctx, filter).Decode(&competitorUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for competitorUser: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the competitorUser: %w", err)
	}

	return &competitorUser, nil
}

// func (r *Repository) UpdateCompetitorUser(ctx context.Context, competitorUserID string, competitorUserInfoDAO *dao.UpdateCompetitorUserDAOReq) error {
// 	competitorUserOID, err := r.ConvertToObjectID(competitorUserID)
// 	if err != nil {
// 		return err
// 	}

// 	competitorUserInfoDAO.RenewUpdate()

// 	filter := bson.M{"_id": *competitorUserOID}
// 	update, err := api_assets.StructToBsonMap(competitorUserInfoDAO)
// 	if err != nil {
// 		return err
// 	}

// 	result, err := r.competitorUserColl.UpdateOne(
// 		ctx,
// 		filter,
// 		bson.M{"$set": update},
// 	)
// 	if err != nil {
// 		return fmt.Errorf("error updating competitorUser: %w", err)
// 	}

// 	if result.MatchedCount == 0 {
// 		return fmt.Errorf("%w: no competitorUser found with id: %s", customerrors.ErrNotFound, competitorUserID)
// 	}

// 	return nil
// }

func (r *Repository) DeleteCompetitorUser(ctx context.Context, competitorUserID string) error {
	err := r.setDeletedAt(ctx, r.competitorUserColl, competitorUserID, "competitorUser")
	if err != nil {
		return err
	}

	return nil
}
