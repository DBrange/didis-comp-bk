package repository

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/double/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateDouble(ctx context.Context, doubleInfoDAO *dao.CreateDoubleDAOReq) (*primitive.ObjectID, error) {
	doubleInfoDAO.SetTimeStamp()

	result, err := r.doubleColl.InsertOne(ctx, doubleInfoDAO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, fmt.Errorf("%w: error duplicate key for double: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return nil, fmt.Errorf("%w: error double scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return nil, fmt.Errorf("error when inserting double: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID)

	return &id, nil
}

func (r *Repository) GetDoubleByID(ctx context.Context, doubleID string) (*dao.GetDoubleByIDDAORes, error) {
	var double dao.GetDoubleByIDDAORes

	doubleOID, err := r.ConvertToObjectID(doubleID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *doubleOID}

	err = r.doubleColl.FindOne(ctx, filter).Decode(&double)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for double: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the double: %w", err)
	}

	return &double, nil
}

// func (r *Repository) UpdateDouble(ctx context.Context, doubleID string, doubleInfoDAO *dao.UpdateDoubleDAOReq) error {
// 	doubleOID, err := r.ConvertToObjectID(doubleID)
// 	if err != nil {
// 		return err
// 	}

// 	doubleInfoDAO.RenewUpdate()

// 	filter := bson.M{"_id": *doubleOID}
// 	update, err := api_assets.StructToBsonMap(doubleInfoDAO)
// 	if err != nil {
// 		return err
// 	}

// 	result, err := r.doubleColl.UpdateOne(
// 		ctx,
// 		filter,
// 		bson.M{"$set": update},
// 	)
// 	if err != nil {
// 		return fmt.Errorf("error updating double: %w", err)
// 	}

// 	if result.MatchedCount == 0 {
// 		return fmt.Errorf("%w: no double found with id: %s", customerrors.ErrNotFound, doubleID)
// 	}

// 	return nil
// }

func (r *Repository) DeleteDouble(ctx context.Context, doubleID string) error {
	err := r.setDeletedAt(ctx, r.doubleColl, doubleID, "double")
	if err != nil {
		return err
	}

	return nil
}
