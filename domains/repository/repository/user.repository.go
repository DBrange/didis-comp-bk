package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/assets"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateUser(ctx context.Context, user *dao.CreateUserDAO) error {

	user.SetTimeStamp()

	_, err := r.user_coll.InsertOne(ctx, user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return fmt.Errorf("%w: error duplicate key for user: %s", customerrors.ErrUserDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return fmt.Errorf("%w: error user schema type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}
		return fmt.Errorf("error when inserting user: %w", err)
	}

	return nil
}

func (r *Repository) GetUserByID(ctx context.Context, id string) (*dao.GetUserByIDDAO, error) {
	var user dao.GetUserByIDDAO

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("%w: invalid user id format: %s", customerrors.ErrUserInvalidID, err.Error())
	}

	filter := bson.M{"_id": oid}

	err = r.user_coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for the user: %s", customerrors.ErrUserNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the user: %w", err)
	}

	return &user, nil
}

func (r *Repository) UpdateUser(ctx context.Context, userID string, newUserInfo *dao.UpdateUserDAOReq) error {
	userOID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("invalid id format: %w", err)
	}

	filter := bson.M{"_id": userOID}

	update, err := assets.StructToBsonMap(newUserInfo)
	if err != nil {
		return err
	}
	currentDate := time.Now().UTC()
	update["updated_at"] = currentDate

	result, err := r.user_coll.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": update},
	)
	if err != nil {
		return fmt.Errorf("%w: error updating user: %s", customerrors.ErrUserUpdated, err.Error())
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no user document found with the given filter", customerrors.ErrUserUpdated)
	}

	return nil
}

func (r *Repository) DeleteUser(ctx context.Context, userID string) (*dao.UserRelationsToDeleteDAO, error) {

	projections := bson.M{
		"location_id": 1,
		"payments_id": 1,
		"schedule_id": 1,
	}

	userRelationsToDelete, err := setDeletedAtAndReturnIDs(r.user_coll, ctx, userID, projections, &dao.UserRelationsToDeleteDAO{})
	if err != nil {
		return nil, err
	}

	return userRelationsToDelete, nil
}
