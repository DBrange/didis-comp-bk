package adapters

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/chat/adapters/mappers"
	"github.com/DBrange/didis-comp-bk/domains/chat/models/dto"
	ports "github.com/DBrange/didis-comp-bk/domains/repository/ports/drivers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatQuerierAdapter struct {
	adapter ports.ForManagingChat
}

func NewChatQuerierAdapter(adapter ports.ForManagingChat) *ChatQuerierAdapter {
	return &ChatQuerierAdapter{
		adapter: adapter,
	}
}
func (a *ChatQuerierAdapter) WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) error) error {
	return a.adapter.WithTransaction(ctx, fn)
}

func (a *ChatQuerierAdapter) ConvertToObjectID(ID string) (*primitive.ObjectID, error) {
	return a.adapter.ConvertToObjectID(ID)
}

func (a *ChatQuerierAdapter) CreateChat(ctx context.Context, matchID string, chatType models.CHAT) (string, error) {
	matchOID, err := a.ConvertToObjectID(matchID)
	if err != nil {
		return "", err
	}

	return a.adapter.CreateChat(ctx, matchOID, chatType)
}

func (a *ChatQuerierAdapter) CreateParticipantChats(ctx context.Context, chatID string, participantIDs []string, chatType models.CHAT) error {
	chatOID, err := a.ConvertToObjectID(chatID)
	if err != nil {
		return err
	}

	participantOIDs, err := utils.ConvertToObjectIDs(&participantIDs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.adapter.CreateParticipantChats(ctx, chatOID, *participantOIDs, chatType)
}

func (a *ChatQuerierAdapter) GetMatchChatByID(ctx context.Context, chatID string) (*dto.GetMatchChatByIDDTORes, error) {
	chatDAO, err := a.adapter.GetMatchChatByID(ctx, chatID)
	if err != nil {
		return nil, err
	}

	chatDTO := mappers.GetMatchChatByIDDAOtoDTO(chatDAO)

	return chatDTO, nil
}

func (a *ChatQuerierAdapter) CreateChatMessage(ctx context.Context, chatMessageDTO *dto.CreateChatMessageDTOReq) (string, error) {
	chatMessageDAO, err := mappers.CreateChatMessageDTOtoDAO(chatMessageDTO, a.ConvertToObjectID)
	if err != nil {
		return "", err
	}

	return a.adapter.CreateChatMessage(ctx, chatMessageDAO)
}

func (a *ChatQuerierAdapter) VerifyChatExists(ctx context.Context, chatID string) error {
	chatOID, err := a.ConvertToObjectID(chatID)
	if err != nil {
		return err
	}

	return a.adapter.VerifyChatExists(ctx, chatOID)
}

func (a *ChatQuerierAdapter) VerifyUserExists(ctx context.Context, userID string) error {
	userOID, err := a.ConvertToObjectID(userID)
	if err != nil {
		return err
	}

	return a.adapter.VerifyUserExists(ctx, userOID)
}
