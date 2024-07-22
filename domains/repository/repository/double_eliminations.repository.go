package repository

import (
	"context"
	"fmt"

	api_assets "github.com/DBrange/didis-comp-bk/cmd/api/assets"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/double_elimination/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateDoubleElimination(ctx context.Context, doubleEliminationInfoDAO *dao.CreateDoubleEliminationDAOReq) (string, error) {
	doubleEliminationInfoDAO.SetTimeStamp()

	result, err := r.doubleEliminationColl.InsertOne(ctx, doubleEliminationInfoDAO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", fmt.Errorf("%w: error duplicate key for doubleElimination: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error doubleElimination scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting doubleElimination: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetDoubleEliminationByID(ctx context.Context, doubleEliminationID string) (*dao.GetDoubleEliminationByIDDAORes, error) {
	var doubleElimination dao.GetDoubleEliminationByIDDAORes

	doubleEliminationOID, err := r.ConvertToObjectID(doubleEliminationID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *doubleEliminationOID}

	err = r.doubleEliminationColl.FindOne(ctx, filter).Decode(&doubleElimination)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for doubleElimination: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the doubleElimination: %w", err)
	}

	return &doubleElimination, nil
}

func (r *Repository) UpdateDoubleElimination(ctx context.Context, doubleEliminationID string, doubleEliminationInfoDAO *dao.UpdateDoubleEliminationDAOReq) error {
	doubleEliminationOID, err := r.ConvertToObjectID(doubleEliminationID)
	if err != nil {
		return err
	}

	doubleEliminationInfoDAO.RenewUpdate()

	filter := bson.M{"_id": *doubleEliminationOID}
	update, err := api_assets.StructToBsonMap(doubleEliminationInfoDAO)
	if err != nil {
		return err
	}

	result, err := r.doubleEliminationColl.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": update},
	)
	if err != nil {
		return fmt.Errorf("error updating doubleElimination: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no doubleElimination found with id: %s", customerrors.ErrNotFound, doubleEliminationID)
	}

	return nil
}

func (r *Repository) DeleteDoubleElimination(ctx context.Context, doubleEliminationID string) error {
	err := r.setDeletedAt(ctx, r.doubleEliminationColl, doubleEliminationID, "doubleElimination")
	if err != nil {
		return err
	}

	return nil
}
