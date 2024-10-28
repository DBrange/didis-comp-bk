package repository

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	api_assets "github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/chat/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateChat(ctx, matchID, competitorIDs, models.CHAT_GROUP)
func (r *Repository) CreateChat(ctx context.Context, matchOID *primitive.ObjectID, chatType models.CHAT) (string, error) {
	chatDAO := &dao.CreateChatDAOReq{
		ChatType:           chatType,
		AvailabilityStatus: models.CHAT_AVAILABILITY_STATUS_INDECISION,
		MatchID:            matchOID,
	}

	chatDAO.SetTimeStamp()

	result, err := r.chatColl.InsertOne(ctx, chatDAO)
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

func (r *Repository) GetChatByID(ctx context.Context, chatOID *primitive.ObjectID) (*dao.GetChatByIDDAORes, error) {
	var chat dao.GetChatByIDDAORes

	filter := bson.M{"_id": *chatOID}

	err := r.chatColl.FindOne(ctx, filter).Decode(&chat)
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
	err := r.SetDeletedAt(ctx, r.chatColl, chatID, "chat")
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetMatchChatByID(ctx context.Context, chatOID *primitive.ObjectID) (*dao.GetMatchChatByIDDAORes, error) {
	pipeline := []bson.D{
		// Etapa 1: Comenzar con el chat específico
		{{
			Key: "$match", Value: bson.M{
				"_id": chatOID, // Asume que tienes el ID del chat
			},
		}},
		// Etapa 2: Buscar participant_chats relacionados
		{{
			Key: "$lookup", Value: bson.M{
				"from":         "participant_chats",
				"localField":   "_id",
				"foreignField": "chat_id",
				"as":           "participants",
			},
		}},
		// Etapa 3: Desagregar los participantes
		{{
			Key: "$unwind", Value: bson.M{
				"path": "$participants",
			},
		}},
		// Etapa 4: Buscar usuarios relacionados
		{{
			Key: "$lookup", Value: bson.M{
				"from":         "users",
				"localField":   "participants.user_id",
				"foreignField": "_id",
				"as":           "user",
			},
		}},
		// Etapa 5: Buscar competitor_users relacionados
		{{
			Key: "$lookup", Value: bson.M{
				"from":         "competitor_users",
				"localField":   "participants.competitor_id",
				"foreignField": "competitor_id",
				"as":           "competitor_user",
			},
		}},
		// Etapa 6: Buscar usuarios de competidores
		{{
			Key: "$lookup", Value: bson.M{
				"from":         "users",
				"localField":   "competitor_user.user_id",
				"foreignField": "_id",
				"as":           "competitor_user_details",
			},
		}},
		// Etapa 7: Agrupar y estructurar el resultado
		{{
			Key: "$group", Value: bson.M{
				"_id":                 "$_id",
				"availability_status": bson.M{"$first": "$availability_status"},
				"match_id":            bson.M{"$first": "$match_id"},
				"users": bson.M{
					"$addToSet": bson.M{
						"$cond": bson.M{
							"if": bson.M{"$eq": []interface{}{bson.M{"$size": "$user"}, 1}},
							"then": bson.M{
								"_id":        bson.M{"$arrayElemAt": []interface{}{"$user._id", 0}},
								"first_name": bson.M{"$arrayElemAt": []interface{}{"$user.first_name", 0}},
								"last_name":  bson.M{"$arrayElemAt": []interface{}{"$user.last_name", 0}},
								"image":      bson.M{"$arrayElemAt": []interface{}{"$user.image", 0}},
							},
							"else": nil,
						},
					},
				},
				"competitors": bson.M{
					"$addToSet": bson.M{
						"$cond": bson.M{
							"if": bson.M{"$eq": []interface{}{bson.M{"$size": "$competitor_user"}, 1}},
							"then": bson.M{
								"_id":                 "$participants.competitor_id",
								"availability_status": "$participants.availability_status",
								"users":               "$competitor_user_details",
							},
							"else": nil,
						},
					},
				},
			},
		}},
		// Etapa 8: Limpiar los resultados nulos
		{{
			Key: "$project", Value: bson.M{
				"_id":                 1,
				"availability_status": 1,
				"match_id":            1,
				"users": bson.M{
					"$filter": bson.M{
						"input": "$users",
						"as":    "user",
						"cond":  bson.M{"$ne": []interface{}{"$$user", nil}},
					},
				},
				"competitors": bson.M{
					"$filter": bson.M{
						"input": "$competitors",
						"as":    "competitor",
						"cond":  bson.M{"$ne": []interface{}{"$$competitor", nil}},
					},
				},
			},
		}},
	}

	cursor, err := r.chatColl.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result dao.GetMatchChatByIDDAORes
	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("error when decoding chat: %w", err)
		}
	} else {
		return nil, fmt.Errorf("%w: no chat found with id: %s", customerrors.ErrNotFound, chatOID.Hex())
	}

	return &result, nil
}

func (r *Repository) VerifyChatExists(ctx context.Context, chatOID *primitive.ObjectID) error {
	var result struct{}

	filter := bson.M{"_id": chatOID}

	opts := options.FindOne().SetProjection(bson.M{"_id": 1})

	err := r.chatColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("%w: error when searching for chat: %s", customerrors.ErrNotFound, err.Error())
		}
		return fmt.Errorf("error when searching for the chat: %w", err)
	}

	return nil
}
