package repository

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/participant_chat/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateParticipantChats(ctx context.Context, chatOID *primitive.ObjectID, participantOIDs []*primitive.ObjectID, chatType models.CHAT) error {
	// Crear un slice para almacenar las solicitudes de inserción
	participantChats := make([]interface{}, len(participantOIDs))

	for i, participantOID := range participantOIDs {
		participantChatInfoDAO := &dao.CreateParticipantChatDAOReq{
			ChatID:             chatOID,
			AvailabilityStatus: models.CHAT_AVAILABILITY_STATUS_INDECISION,
		}

		if chatType == models.CHAT_GROUP {
			participantChatInfoDAO.CompetitorID = participantOID

		} else if chatType == models.CHAT_INDIVIDUAL {
			participantChatInfoDAO.UserID = participantOID
		}

		participantChatInfoDAO.SetTimeStamp()
		participantChats[i] = participantChatInfoDAO
	}

	// Insertar múltiples documentos con InsertMany
	_, err := r.participantChatColl.InsertMany(ctx, participantChats)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return fmt.Errorf("%w: error duplicate key for participantChat: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return fmt.Errorf("%w: error participantChat scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return fmt.Errorf("error when inserting participantChat: %w", err)
	}

	return nil
}

func (r *Repository) GetParticipantChatByID(ctx context.Context, participantChatID string) (*dao.GetParticipantChatByIDDAORes, error) {
	var participantChat dao.GetParticipantChatByIDDAORes

	participantChatOID, err := r.ConvertToObjectID(participantChatID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *participantChatOID}

	err = r.participantChatColl.FindOne(ctx, filter).Decode(&participantChat)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for participantChat: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the participantChat: %w", err)
	}

	return &participantChat, nil
}

// func (r *Repository) UpdateParticipantChat(ctx context.Context, participantChatID string, participantChatInfoDAO *dao.UpdateParticipantChatDAOReq) error {
// 	participantChatOID, err := r.ConvertToObjectID(participantChatID)
// 	if err != nil {
// 		return err
// 	}

// 	participantChatInfoDAO.RenewUpdate()

// 	filter := bson.M{"_id": *participantChatOID}
// 	update, err := api_assets.StructToBsonMap(participantChatInfoDAO)
// 	if err != nil {
// 		return err
// 	}

// 	result, err := r.participantChatColl.UpdateOne(
// 		ctx,
// 		filter,
// 		bson.M{"$set": update},
// 	)
// 	if err != nil {
// 		return fmt.Errorf("error updating participantChat: %w", err)
// 	}

// 	if result.MatchedCount == 0 {
// 		return fmt.Errorf("%w: no participantChat found with id: %s", customerrors.ErrNotFound, participantChatID)
// 	}

// 	return nil
// }

func (r *Repository) DeleteParticipantChat(ctx context.Context, participantChatID string) error {
	err := r.SetDeletedAt(ctx, r.participantChatColl, participantChatID, "participantChat")
	if err != nil {
		return err
	}

	return nil
}
