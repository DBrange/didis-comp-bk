package repository

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/domains/repository/models"
	user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) CreateUser(ctx context.Context, userDAO *user_dao.CreateUserDAOReq) (string, error) {
	userDAO.SetTimeStamp()

	// SACAR ESTO LUEGO
	userDAO.Active = true

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

func (r *Repository) GetUserByID(ctx context.Context, userID string) (*user_dao.GetUserByIDDAORes, error) {
	var user user_dao.GetUserByIDDAORes

	userOID, err := r.ConvertToObjectID(userID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": *userOID,
		"$or": []bson.M{
			{"deleted_at": bson.M{"$exists": false}},
			{"deleted_at": nil},
		},
		"active": true,
	}

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

	// update, err := api_assets.StructToBsonMap(userInfoDAO)
	// if err != nil {
	// 	return err
	// }
	// currentDate := time.Now().UTC()
	// update["updated_at"] = currentDate

	result, err := r.userColl.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": userInfoDAO},
	)
	if err != nil {
		return fmt.Errorf("%w: error updating user: %s", customerrors.ErrUpdated, err.Error())
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no user found with id: %s", customerrors.ErrNotFound, userID)
	}

	return nil
}

// func (r *Repository) updateUserConcurrently(sessCtx mongo.SessionContext, userID string, userInfoDAO *user_dao.UpdateUserDAOReq, wg *sync.WaitGroup, errCh chan<- error) {
// 	defer wg.Done()
// 	if err := r.UpdateUser(sessCtx, userID, userInfoDAO); err != nil {
// 		errCh <- err
// 	}
// }

func (r *Repository) DeleteUser(ctx context.Context, userID string) (*user_dao.UserRelationsToDeleteDAOReq, error) {

	projections := bson.M{
		"location_id": 1,
	}

	userRelationsToDelete, err := SetDeletedAtAndReturnIDs(ctx, r.userColl, userID, "user", projections, &user_dao.UserRelationsToDeleteDAOReq{})
	if err != nil {
		return nil, err
	}

	return userRelationsToDelete, nil
}

func (r *Repository) UpdateUserPassword(ctx context.Context, userID, newPassword string) error {
	userOID, err := r.ConvertToObjectID(userID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": userOID}

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

func (r *Repository) GetUserPasswordByID(ctx context.Context, userID string) (string, error) {
	userOID, err := r.ConvertToObjectID(userID)
	if err != nil {
		return "", err
	}
	filter := bson.M{
		"_id":    userOID,
		"$and":   []bson.M{models.OmitDeleted()},
		"active": true,
	}

	type password struct {
		ID       string `bson:"_id"`
		Password string `bson:"password"`
	}

	var getPassword password

	projection := bson.M{"password": 1, "_id": 1}

	opts := options.FindOne().SetProjection(projection)

	err = r.userColl.FindOne(ctx, filter, opts).Decode(&getPassword)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", fmt.Errorf("%w: error when searching for the user: %s", customerrors.ErrNotFound, err.Error())
		}
		return "", fmt.Errorf("error when searching for the user: %w", err)

	}

	return getPassword.Password, nil
}

func (r *Repository) GetUserPasswordForLogin(ctx context.Context, username string) (string, string, error) {
	// Filter options username
	filterUsername := bson.M{
		"$or": []bson.M{
			{"email": username},
			{"username": username},
		},
	}

	filter := bson.M{
		"$and": []bson.M{
			filterUsername,
			models.OmitDeleted(),
		},
		"active": true,
	}

	type password struct {
		ID       string `bson:"_id"`
		Password string `bson:"password"`
	}

	var getPassword password

	projection := bson.M{"password": 1, "_id": 1}

	opts := options.FindOne().SetProjection(projection)

	err := r.userColl.FindOne(ctx, filter, opts).Decode(&getPassword)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", "", fmt.Errorf("%w: error when searching for the user: %s", customerrors.ErrNotFound, err.Error())
		}
		return "", "", fmt.Errorf("error when searching for the user: %w", err)

	}

	return getPassword.Password, getPassword.ID, nil
}

func (r *Repository) GetUserRoles(ctx context.Context, userID string) ([]string, error) {
	type Roles struct {
		Roles []primitive.ObjectID `bson:"roles"`
	}

	var roles Roles

	userOID, err := r.ConvertToObjectID(userID)
	if err != nil {
		return []string{}, nil
	}

	filter := bson.M{
		"_id":    userOID,
		"$and":   models.OmitDeleted(),
		"active": true,
	}

	projection := bson.M{"roles": 1}

	opts := options.FindOne().SetProjection(projection)

	err = r.userColl.FindOne(ctx, filter, opts).Decode(&roles)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []string{}, fmt.Errorf("%w: error when searching for the user: %s", customerrors.ErrNotFound, err.Error())
		}
		return []string{}, fmt.Errorf("error when searching for the user: %w", err)
	}

	rolesStr := make([]string, len(roles.Roles))

	for i, role := range roles.Roles {
		rolesStr[i] = role.Hex()
	}

	return rolesStr, nil
}

func (r *Repository) GetUserForListOfTournament(ctx context.Context, userOID *primitive.ObjectID)(any, error){
return nil,nil
}


