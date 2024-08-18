package repository

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/single/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateSingle(ctx context.Context, singleInfoDAO *dao.CreateSingleDAOReq) (string, error) {
	singleInfoDAO.SetTimeStamp()

	result, err := r.singleColl.InsertOne(ctx, singleInfoDAO)
	if err != nil {
		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error single scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting single: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetSingleByID(ctx context.Context, singleID string) (*dao.GetSingleByIDDAORes, error) {
	var single dao.GetSingleByIDDAORes

	singleOID, err := r.ConvertToObjectID(singleID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *singleOID}

	err = r.singleColl.FindOne(ctx, filter).Decode(&single)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for single: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the single: %w", err)
	}

	return &single, nil
}

// func (r *Repository) UpdateSingle(ctx context.Context, singleID string, singleInfoDAO *dao.UpdateSingleDAOReq) error {
// 	singleOID, err := r.ConvertToObjectID(singleID)
// 	if err != nil {
// 		return err
// 	}

// 	singleInfoDAO.RenewUpdate()

// 	filter := bson.M{"_id": *singleOID}
// 	update, err := api_assets.StructToBsonMap(singleInfoDAO)
// 	if err != nil {
// 		return err
// 	}

// 	result, err := r.singleColl.UpdateOne(
// 		ctx,
// 		filter,
// 		bson.M{"$set": update},
// 	)
// 	if err != nil {
// 		return fmt.Errorf("error updating single: %w", err)
// 	}

// 	if result.MatchedCount == 0 {
// 		return fmt.Errorf("%w: no single found with id: %s", customerrors.ErrNotFound, singleID)
// 	}

// 	return nil
// }

func (r *Repository) DeleteSingle(ctx context.Context, singleID string) error {
	err := r.SetDeletedAt(ctx, r.singleColl, singleID, "single")
	if err != nil {
		return err
	}

	return nil
}
