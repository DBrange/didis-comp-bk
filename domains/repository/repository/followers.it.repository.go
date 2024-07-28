package repository

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/follower/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateFollower(ctx context.Context, followerInfoDAO *dao.CreateFollowerDAOReq) (string, error) {
	followerInfoDAO.SetTimeStamp()

	result, err := r.followerColl.InsertOne(ctx, followerInfoDAO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", fmt.Errorf("%w: error duplicate key for follower: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error follower scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting follower: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetFollowerByID(ctx context.Context, followerID string) (*dao.GetFollowerByIDDAORes, error) {
	var follower dao.GetFollowerByIDDAORes

	followerOID, err := r.ConvertToObjectID(followerID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *followerOID}

	err = r.followerColl.FindOne(ctx, filter).Decode(&follower)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for follower: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the follower: %w", err)
	}

	return &follower, nil
}

// func (r *Repository) UpdateFollower(ctx context.Context, followerID string, followerInfoDAO *dao.UpdateFollowerDAOReq) error {
// 	followerOID, err := r.ConvertToObjectID(followerID)
// 	if err != nil {
// 		return err
// 	}

// 	followerInfoDAO.RenewUpdate()

// 	filter := bson.M{"_id": *followerOID}
// 	update, err := api_assets.StructToBsonMap(followerInfoDAO)
// 	if err != nil {
// 		return err
// 	}

// 	result, err := r.followerColl.UpdateOne(
// 		ctx,
// 		filter,
// 		bson.M{"$set": update},
// 	)
// 	if err != nil {
// 		return fmt.Errorf("error updating follower: %w", err)
// 	}

// 	if result.MatchedCount == 0 {
// 		return fmt.Errorf("%w: no follower found with id: %s", customerrors.ErrNotFound, followerID)
// 	}

// 	return nil
// }

func (r *Repository) DeleteFollower(ctx context.Context, followerID string) error {
	err := r.SetDeletedAt(ctx, r.followerColl, followerID, "follower")
	if err != nil {
		return err
	}

	return nil
}
