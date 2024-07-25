package repository

import (
	"context"
	"fmt"
	"sync"

	api_assets "github.com/DBrange/didis-comp-bk/cmd/api/utils"
	user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateUser(ctx context.Context, userDAO *user_dao.CreateUserDAOReq) (string, error) {
	userDAO.SetTimeStamp()

	result, err := r.userColl.InsertOne(ctx, userDAO)
	if err != nil {
		fmt.Printf("Error inserting user: %v\n", err)
		if mongo.IsDuplicateKeyError(err) {
			return "", fmt.Errorf("%w: error duplicate key for user: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error user scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}
		return "", fmt.Errorf("error when inserting user: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetUserByID(ctx context.Context, userID string) (*user_dao.GetUserByIDDAO, error) {
	var user user_dao.GetUserByIDDAO

	userOID, err := r.ConvertToObjectID(userID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *userOID}

	err = r.userColl.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for the user: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the user: %w", err)
	}

	return &user, nil
}

func (r *Repository) UpdateUser(ctx context.Context, userID string, userInfoDAO *user_dao.UpdateUserDAOReq) error {
	userOID, err := r.ConvertToObjectID(userID)
	if err != nil {
		return fmt.Errorf("invalid id format: %w", err)
	}

	userInfoDAO.RenewUpdate()

	filter := bson.M{"_id": userOID}

	update, err := api_assets.StructToBsonMap(userInfoDAO)
	fmt.Printf("%+v", &update)
	if err != nil {
		return err
	}
	// currentDate := time.Now().UTC()
	// update["updated_at"] = currentDate

	result, err := r.userColl.UpdateOne(
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

func (r *Repository) updateUserConcurrently(sessCtx mongo.SessionContext, userID string, userInfoDAO *user_dao.UpdateUserDAOReq, wg *sync.WaitGroup, errCh chan<- error) {
	defer wg.Done()
	if err := r.UpdateUser(sessCtx, userID, userInfoDAO); err != nil {
		errCh <- err
	}
}

func (r *Repository) DeleteUser(ctx context.Context, userID string) (*user_dao.UserRelationsToDeleteDAO, error) {

	projections := bson.M{
		"location_id": 1,
	}

	userRelationsToDelete, err := setDeletedAtAndReturnIDs(ctx, r.userColl, userID, "user", projections, &user_dao.UserRelationsToDeleteDAO{})
	if err != nil {
		return nil, err
	}

	return userRelationsToDelete, nil
}

func (r *Repository) UpdateUserPassword(ctx context.Context, userID, newPassword, oldPassword string) error {
	userOID, err := r.ConvertToObjectID(userID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": userOID}

	var user struct {
		Password string `bson:"password"`
	}

	err = r.userColl.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("%w: no user found with id: %s", customerrors.ErrNotFound, userID)
		}
		return fmt.Errorf("%w: error finding user: %s", customerrors.ErrUpdated, err.Error())
	}

	// Paso 2: Comparar la contraseña antigua con la almacenada
	if user.Password != oldPassword {
		return fmt.Errorf("%w: old password does not match", customerrors.ErrInsertionFailed)
	}

	update := bson.M{"password": newPassword}

	result, err := r.userColl.UpdateOne(
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
