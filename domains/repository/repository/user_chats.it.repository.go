package repository

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/user_chat/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateUserChat(ctx context.Context, userChatInfoDAO *dao.CreateUserChatDAOReq) (string, error) {
	userChatInfoDAO.SetTimeStamp()

	result, err := r.userChatColl.InsertOne(ctx, userChatInfoDAO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", fmt.Errorf("%w: error duplicate key for userChat: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error userChat scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting userChat: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetUserChatByID(ctx context.Context, userChatID string) (*dao.GetUserChatByIDDAORes, error) {
	var userChat dao.GetUserChatByIDDAORes

	userChatOID, err := r.ConvertToObjectID(userChatID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *userChatOID}

	err = r.userChatColl.FindOne(ctx, filter).Decode(&userChat)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for userChat: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the userChat: %w", err)
	}

	return &userChat, nil
}

// func (r *Repository) UpdateUserChat(ctx context.Context, userChatID string, userChatInfoDAO *dao.UpdateUserChatDAOReq) error {
// 	userChatOID, err := r.ConvertToObjectID(userChatID)
// 	if err != nil {
// 		return err
// 	}

// 	userChatInfoDAO.RenewUpdate()

// 	filter := bson.M{"_id": *userChatOID}
// 	update, err := api_assets.StructToBsonMap(userChatInfoDAO)
// 	if err != nil {
// 		return err
// 	}

// 	result, err := r.userChatColl.UpdateOne(
// 		ctx,
// 		filter,
// 		bson.M{"$set": update},
// 	)
// 	if err != nil {
// 		return fmt.Errorf("error updating userChat: %w", err)
// 	}

// 	if result.MatchedCount == 0 {
// 		return fmt.Errorf("%w: no userChat found with id: %s", customerrors.ErrNotFound, userChatID)
// 	}

// 	return nil
// }

func (r *Repository) DeleteUserChat(ctx context.Context, userChatID string) error {
	err := r.setDeletedAt(ctx, r.userChatColl, userChatID, "userChat")
	if err != nil {
		return err
	}

	return nil
}
