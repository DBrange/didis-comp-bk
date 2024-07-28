package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SetDeletedAtAndReturnIDsGeneric interface {
	*dao.UserRelationsToDeleteDAOReq
}

func SetDeletedAtAndReturnIDs[T SetDeletedAtAndReturnIDsGeneric](ctx context.Context, mc *mongo.Collection, ID string, name string, projections bson.M, structToUpdate T) (T, error) {
	OID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, fmt.Errorf("%w: invalid availability id format: %s", customerrors.ErrInvalidID, err.Error())
	}

	currentDate := time.Now().UTC()

	update := bson.M{
		"deleted_at": currentDate,
	}
	update["updated_at"] = currentDate

	filter := bson.M{"_id": OID}

	opts := options.FindOneAndUpdate().
		SetReturnDocument(options.After).
		SetProjection(projections)

	err = mc.FindOneAndUpdate(
		ctx,
		filter,
		bson.M{"$set": update},
		opts,
	).Decode(structToUpdate)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching %s: %s", customerrors.ErrNotFound, name, err.Error())
		}
		return nil, fmt.Errorf("%w: error updating %s: %s", customerrors.ErrUpdated, name, err.Error())
	}

	return structToUpdate, nil
}

func (r *Repository) SetDeletedAt(ctx context.Context, mc *mongo.Collection, ID string, name string) error {
	OID, err := r.ConvertToObjectID(ID)
	if err != nil {
		return err
	}

	currentDate := time.Now().UTC()

	update := bson.M{
		"deleted_at": currentDate,
	}
	update["updated_at"] = currentDate

	filter := bson.M{"_id": *OID}

	result, err := mc.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": update},
	)
	if err != nil {
		return fmt.Errorf("%w: error updating %s: %s", customerrors.ErrUpdated, name, err.Error())
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no %s found with id: %s", customerrors.ErrNotFound, name, ID)
	}

	return nil
}

func (r *Repository) DeleteByID(ctx context.Context, mc *mongo.Collection, ID string, name string) error {
	OID, err := r.ConvertToObjectID(ID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": *OID}

	result, err := mc.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("%w: error deleting %s: %s", customerrors.ErrDeleted, name, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("%w: no %s found with id: %s", customerrors.ErrNotFound, name, ID)
	}

	return nil
}

// func (r *Repository) deleteManyByID(ctx context.Context, valuesForDelete []models.ValuesForDelete) error {
// 	for _, v := range valuesForDelete {
// 		if err := r.deleteByID(ctx, v.MC, v.ID, v.Name); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
