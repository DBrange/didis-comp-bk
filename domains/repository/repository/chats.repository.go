package repository

import (
	"context"
	"fmt"

	api_assets "github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/chat/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateChat(ctx context.Context, chatInfoDAO *dao.CreateChatDAOReq) (string, error) {
	chatInfoDAO.SetTimeStamp()

	result, err := r.chatColl.InsertOne(ctx, chatInfoDAO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", fmt.Errorf("%w: error duplicate key for chat: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error chat scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting chat: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetChatByID(ctx context.Context, chatID string) (*dao.GetChatByIDDAORes, error) {
	var chat dao.GetChatByIDDAORes

	chatOID, err := r.ConvertToObjectID(chatID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *chatOID}

	err = r.chatColl.FindOne(ctx, filter).Decode(&chat)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for chat: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the chat: %w", err)
	}

	return &chat, nil
}

func (r *Repository) UpdateChat(ctx context.Context, chatID string, chatInfoDAO *dao.UpdateChatDAOReq) error {
	chatOID, err := r.ConvertToObjectID(chatID)
	if err != nil {
		return err
	}

	chatInfoDAO.RenewUpdate()

	filter := bson.M{"_id": *chatOID}
	update, err := api_assets.StructToBsonMap(chatInfoDAO)
	if err != nil {
		return err
	}

	result, err := r.chatColl.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": update},
	)
	if err != nil {
		return fmt.Errorf("error updating chat: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no chat found with id: %s", customerrors.ErrNotFound, chatID)
	}

	return nil
}

func (r *Repository) DeleteChat(ctx context.Context, chatID string) error {
	err := r.setDeletedAt(ctx, r.chatColl, chatID, "chat")
	if err != nil {
		return err
	}

	return nil
}
