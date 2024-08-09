package repository

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/double_elimination/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) DoubleEliminationColl() *mongo.Collection {
	return r.doubleEliminationColl
}

func (r *Repository) CreateDoubleElimination(ctx context.Context, doubleEliminationDAO *dao.CreateDoubleEliminationDAOReq) (string, error) {
	doubleEliminationDAO.SetTimeStamp()

	result, err := r.doubleEliminationColl.InsertOne(ctx, doubleEliminationDAO)
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
func (r *Repository) CreateDoubleEliminationEmpty(ctx context.Context) (string, error) {
	var doubleEliminationEmptyDAO dao.CreateDoubleEliminationDAOReq

	doubleEliminationEmptyDAO.Matches = []primitive.ObjectID{}
	doubleEliminationEmptyDAO.Rounds = []primitive.ObjectID{}

	doubleEliminationEmptyDAO.SetTimeStamp()

	result, err := r.doubleEliminationColl.InsertOne(ctx, doubleEliminationEmptyDAO)
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

func (r *Repository) UpdateDoubleEliminationBsonStruct(ctx context.Context, doubleEliminationDAO *dao.UpdateDoubleEliminationDAOReq, update *bson.M, add bool) *bson.M {
	operation := "$pull"
	arrayModifier := "$in"

	if add {
		operation = "$push"
		arrayModifier = "$each"
	}

	setUpdates := bson.M{}

	if doubleEliminationDAO.Matches != nil {
		setUpdates["matches"] = bson.M{arrayModifier: doubleEliminationDAO.Matches}
	}

	if doubleEliminationDAO.Rounds != nil {
		setUpdates["rounds"] = bson.M{arrayModifier: doubleEliminationDAO.Rounds}
	}

	if len(setUpdates) > 0 {
		(*update)[operation] = setUpdates
	}

	return update

}

func (r *Repository) UpdateDoubleElimination(ctx context.Context, doubleEliminationOID *primitive.ObjectID, doubleEliminationInfoDAO *dao.UpdateDoubleEliminationDAOReq, add bool) error {
	doubleEliminationInfoDAO.RenewUpdate()

	filter := bson.M{"_id": doubleEliminationOID}

	update := bson.M{}

	update = *r.UpdateDoubleEliminationBsonStruct(ctx, doubleEliminationInfoDAO, &update, add)
	if update == nil {
		return fmt.Errorf("error updating doubleElimination, nothing to update: %w", nil)
	}

	result, err := r.doubleEliminationColl.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return fmt.Errorf("error updating doubleElimination: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no doubleElimination found with id: %s", customerrors.ErrNotFound, doubleEliminationOID.Hex())
	}

	return nil
}

func (r *Repository) DeleteDoubleElimination(ctx context.Context, doubleEliminationID string) error {
	err := r.SetDeletedAt(ctx, r.doubleEliminationColl, doubleEliminationID, "doubleElimination")
	if err != nil {
		return err
	}

	return nil
}
