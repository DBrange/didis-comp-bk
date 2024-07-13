package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/assets"
	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateUserAndLocation(ctx context.Context, userInfoDAO *user_dao.CreateUserDAO, locationInfoDAO *location_dao.CreateLocationDAOReq) error {

	locationID, err := r.CreateLocation(ctx, locationInfoDAO)
	if err != nil {
		return err
	}

	userInfoDAO.LocationID = &locationID

	err = r.CreateUser(ctx, userInfoDAO)
	if err != nil {
		if errDel := r.DeleteByID(ctx, r.location_coll, locationID, "location"); errDel != nil {
			return errDel
		}

		return err
	}

	return nil
}

func (r *Repository) CreateUser(ctx context.Context, user *user_dao.CreateUserDAO) error {

	user.SetTimeStamp()

	_, err := r.user_coll.InsertOne(ctx, user)

	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return fmt.Errorf("%w: error duplicate key for user: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return fmt.Errorf("%w: error user scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}
		return fmt.Errorf("error when inserting user: %w", err)
	}
	return nil
}

func (r *Repository) GetUserByID(ctx context.Context, id string) (*user_dao.GetUserByIDDAO, error) {
	var user user_dao.GetUserByIDDAO

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("%w: invalid user id format: %s", customerrors.ErrInvalidID, err.Error())
	}

	filter := bson.M{"_id": oid}

	err = r.user_coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for the user: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the user: %w", err)
	}

	return &user, nil
}

func (r *Repository) UpdateUser(ctx context.Context, userID string, newUserInfo *user_dao.UpdateUserDAOReq) error {
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
		return fmt.Errorf("%w: error updating user: %s", customerrors.ErrUpdated, err.Error())
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no user found with id: %s", customerrors.ErrNotFound, userID)
	}

	return nil
}

func (r *Repository) DeleteUser(ctx context.Context, userID string) (*user_dao.UserRelationsToDeleteDAO, error) {

	projections := bson.M{
		"location_id": 1,
		"payments_id": 1,
		"schedule_id": 1,
	}

	userRelationsToDelete, err := setDeletedAtAndReturnIDs(r.user_coll, ctx, userID, "user", projections, &user_dao.UserRelationsToDeleteDAO{})
	if err != nil {
		return nil, err
	}

	return userRelationsToDelete, nil
}

// Ejemplo de como utilizar session
// func (r *Repository) RegisterUser(ctx context.Context, userInfoDAO *user_dao.CreateUserDAO, locationInfoDAO *location_dao.CreateLocationDAOReq) error {
// 	if r.client == nil {
// 		return fmt.Errorf("%w: MongoDB client is not initialized", customerrors.ErrStartSessionFailed)
// 	}
// 	session, err := r.client.StartSession()
// 	if err != nil {
// 		return fmt.Errorf("%w: error starting session: %s", customerrors.ErrStartSessionFailed, err.Error())
// 	}
// 	defer session.EndSession(ctx)

// 	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {

// 		locationID, err := r.CreateLocation2(sessCtx, locationInfoDAO)
// 		if err != nil {
// 			return nil, err
// 		}

// 		userInfoDAO.LocationID = &locationID

// 		err = r.CreateUser2(sessCtx, userInfoDAO)
// 		if err != nil {
// 			return nil, err
// 		}

// 		return nil, nil
// 	}

// 	_, err = session.WithTransaction(ctx, callback)
// 	if err != nil {
// 		fmt.Printf("%v", err)
// 		return fmt.Errorf("%w: transaction error: %s", customerrors.ErrTransaction, err.Error())
// 	}

// 	return nil
// }
