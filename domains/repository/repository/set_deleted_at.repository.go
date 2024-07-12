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

func setDeletedAtAndReturnIDs[T setDeletedAtAndReturnIDsGeneric](mc *mongo.Collection, ctx context.Context, userID string, projections bson.M, structToUpdate T) (T, error) {
	userOID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("%w: invalid id format: %s", customerrors.ErrUserInvalidID, err.Error())
	}

	currentDate := time.Now().UTC()

	update := bson.M{
		"deleted_at": currentDate,
	}
	update["updated_at"] = currentDate

	filter := bson.M{"_id": userOID}

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
			return nil, fmt.Errorf("%w: error when searching: %s", customerrors.ErrUserNotFound, err.Error())
		}
		return nil, fmt.Errorf("%w: error updating: %s", customerrors.ErrUserUpdated, err.Error())
	}

	return structToUpdate, nil
}

func (r *Repository) setDeletedAt(mc *mongo.Collection, ctx context.Context, ID string) error {
	OID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return fmt.Errorf("%w: invalid id format: %s", customerrors.ErrUserInvalidID, err.Error())
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
		return fmt.Errorf("%w: error updating: %s", customerrors.ErrUserUpdated, err.Error())
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no document found with the given filter", customerrors.ErrUserUpdated)
	}

	return nil
}
