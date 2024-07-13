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

type setDeletedAtAndReturnIDsGeneric interface {
	*dao.UserRelationsToDeleteDAO
}

func setDeletedAtAndReturnIDs[T setDeletedAtAndReturnIDsGeneric](mc *mongo.Collection, ctx context.Context, ID string, name string, projections bson.M, structToUpdate T) (T, error) {
	OID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, fmt.Errorf("%w: invalid %s id format: %s", customerrors.ErrInvalidID, name, err.Error())
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

func (r *Repository) setDeletedAt(mc *mongo.Collection, ctx context.Context, ID string, name string) error {
	OID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return fmt.Errorf("%w: invalid %s id format: %s", customerrors.ErrInvalidID, name, err.Error())
	}

	currentDate := time.Now().UTC()

	update := bson.M{
		"deleted_at": currentDate,
	}
	update["updated_at"] = currentDate

	filter := bson.M{"_id": OID}

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
	OID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return fmt.Errorf("%w: invalid %s id format: %s", customerrors.ErrInvalidID, name, err.Error())
	}

	filter := bson.M{"_id": OID}

	result, err := mc.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("%w: error deleting %s: %s", customerrors.ErrDeleted, name, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("%w: no %s found with id: %s", customerrors.ErrNotFound, name, ID)
	}

	return nil
}
