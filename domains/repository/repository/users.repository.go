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

func (r *Repository) GetUserForLogin(ctx context.Context, username string) (*user_dao.GetUserForLoginDAO, error) {
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

	var userData user_dao.GetUserForLoginDAO

	projection := bson.M{"password": 1, "_id": 1, "first_name": 1, "last_name": 1, "roles": 1, "username": 1, "image": 1}

	opts := options.FindOne().SetProjection(projection)

	err := r.userColl.FindOne(ctx, filter, opts).Decode(&userData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for the user: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the user: %w", err)

	}

	return &userData, nil
}

func (r *Repository) GetUserForRefreshToken(ctx context.Context, userID *primitive.ObjectID) (*user_dao.GetUserForRefreshTokenDAO, error) {

	filter := bson.M{
		"$and": []bson.M{
			{"_id": userID},
			models.OmitDeleted(),
		},
		"active": true,
	}

	var userData user_dao.GetUserForRefreshTokenDAO

	projection := bson.M{"_id": 1, "roles": 1}

	opts := options.FindOne().SetProjection(projection)

	err := r.userColl.FindOne(ctx, filter, opts).Decode(&userData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for the user: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the user: %w", err)

	}

	return &userData, nil
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
		"$and":   []bson.M{models.OmitDeleted()},
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

func (r *Repository) VerifyUserExists(ctx context.Context, userOID *primitive.ObjectID) error {
	var result struct{}

	filter := bson.M{"_id": userOID}

	opts := options.FindOne().SetProjection(bson.M{"_id": 1})

	err := r.userColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("%w: error when searching for user: %s", customerrors.ErrNotFound, err.Error())
		}
		return fmt.Errorf("error when searching for the user: %w", err)
	}

	return nil
}

func (r *Repository) GetUserPrimaryData(ctx context.Context, userOID *primitive.ObjectID) (*user_dao.GetUserPrimaryDataDAORes, error) {
	// Filter options username
	filter := bson.M{
		"$and": []bson.M{
			{"_id": userOID	},
			models.OmitDeleted(),
		},
		"active": true,
	}

	var userData user_dao.GetUserPrimaryDataDAORes

	projection := bson.M{ "_id": 1, "first_name": 1, "last_name": 1, "username": 1, "image": 1, "location_id": 1}

	opts := options.FindOne().SetProjection(projection)

	err := r.userColl.FindOne(ctx, filter, opts).Decode(&userData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for the user: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the user: %w", err)

	}

	return &userData, nil
}

// func (r *Repository) GetUserByName(ctx context.Context, name string, limit, page int) ([]*user_dao.GetUserByNameDAORes, error) {
// 	filterRegex := "^" + name
// 	categoryOID := "dsas"
// 	pipeline := mongo.Pipeline{
// 		bson.D{{"$match", bson.M{"category_id": categoryOID}}},
// 	}

// 	projection := bson.M{
// 		"_id":        1,
// 		"first_name": 1,
// 		"last_name":  1,
// 	}
// 	skip := (page - 1) * limit

// 	opts := options.Find().
// 		SetProjection(projection).
// 		SetLimit(int64(limit)).
// 		SetSkip(int64(skip))

// 	cursor, err := r.userColl.Find(ctx, filter, opts)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil, fmt.Errorf("%w: error when searching for user: %s", customerrors.ErrNotFound, err.Error())
// 		}
// 		return nil, fmt.Errorf("error when searching for the user: %w", err)
// 	}

// 	defer cursor.Close(ctx)

// 	var users []*user_dao.GetUserByNameDAORes

// 	if err = cursor.All(ctx, &users); err != nil {
// 		return nil, fmt.Errorf("error when decoding users: %w", err)
// 	}

// 	if len(users) == 0 {
// 		return nil, fmt.Errorf("%w: no user registrations found for name: %s on page %d", customerrors.ErrNotFound, name, page)
// 	}

// 	return users, nil
// }
