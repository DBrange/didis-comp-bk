package repository

import (
	"context"
	"fmt"

	api_assets "github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/chat_message/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateChatMessage(ctx context.Context, chatMessageInfoDAO *dao.CreateChatMessageDAOReq) (string, error) {
	chatMessageInfoDAO.SetTimeStamp()

	result, err := r.chatMessageColl.InsertOne(ctx, chatMessageInfoDAO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", fmt.Errorf("%w: error duplicate key for chatMessage: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error chatMessage scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting chatMessage: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetChatMessageByID(ctx context.Context, chatMessageID string) (*dao.GetChatMessageByIDDAORes, error) {
	var chatMessage dao.GetChatMessageByIDDAORes

	chatMessageOID, err := r.ConvertToObjectID(chatMessageID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *chatMessageOID}

	err = r.chatMessageColl.FindOne(ctx, filter).Decode(&chatMessage)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for chatMessage: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the chatMessage: %w", err)
	}

	return &chatMessage, nil
}

func (r *Repository) UpdateChatMessage(ctx context.Context, chatMessageID string, chatMessageInfoDAO *dao.UpdateChatMessageDAOReq) error {
	chatMessageOID, err := r.ConvertToObjectID(chatMessageID)
	if err != nil {
		return err
	}

	chatMessageInfoDAO.RenewUpdate()

	filter := bson.M{"_id": *chatMessageOID}
	update, err := api_assets.StructToBsonMap(chatMessageInfoDAO)
	if err != nil {
		return err
	}

	result, err := r.chatMessageColl.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": update},
	)
	if err != nil {
		return fmt.Errorf("error updating chatMessage: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no chatMessage found with id: %s", customerrors.ErrNotFound, chatMessageID)
	}

	return nil
}

func (r *Repository) DeleteChatMessage(ctx context.Context, chatMessageID string) error {
	err := r.setDeletedAt(ctx, r.chatMessageColl, chatMessageID, "chatMessage")
	if err != nil {
		return err
	}

	return nil
}
