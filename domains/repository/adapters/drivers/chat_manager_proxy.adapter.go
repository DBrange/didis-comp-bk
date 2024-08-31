package adapters

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	chat_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/chat/dao"
	chat_message_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/chat_message/dao"
	"github.com/DBrange/didis-comp-bk/domains/repository/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatManagerProxyAdapter struct {
	repository *repository.Repository
}

func NewChatManagerProxyAdapter(repository *repository.Repository) *ChatManagerProxyAdapter {
	return &ChatManagerProxyAdapter{
		repository: repository,
	}
}

func (a *ChatManagerProxyAdapter) WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) error) error {
	return a.repository.WithTransaction(ctx, fn)
}

func (a *ChatManagerProxyAdapter) ConvertToObjectID(ID string) (*primitive.ObjectID, error) {
	return a.repository.ConvertToObjectID(ID)
}

func (a *ChatManagerProxyAdapter) CreateChat(ctx context.Context, matchOID *primitive.ObjectID, chatType models.CHAT) (string, error) {
	return a.repository.CreateChat(ctx, matchOID, chatType)
}

func (a *ChatManagerProxyAdapter) CreateParticipantChats(ctx context.Context, chatOID *primitive.ObjectID, participantOIDs []*primitive.ObjectID, chatType models.CHAT) error {
	return a.repository.CreateParticipantChats(ctx, chatOID, participantOIDs, chatType)
}

func (a *ChatManagerProxyAdapter) GetMatchChatByID(ctx context.Context, chatID string) (*chat_dao.GetMatchChatByIDDAORes, error) {
	chatOID, err := a.ConvertToObjectID(chatID)
	if err != nil {
		return nil, err
	}

	return a.repository.GetMatchChatByID(ctx, chatOID)
}

func (a *ChatManagerProxyAdapter) CreateChatMessage(ctx context.Context, chatMessageDAO *chat_message_dao.CreateChatMessageDAOReq) (string, error) {
	return a.repository.CreateChatMessage(ctx, chatMessageDAO)
}

func (a *ChatManagerProxyAdapter) VerifyChatExists(ctx context.Context, chatOID *primitive.ObjectID) error {
	return a.repository.VerifyChatExists(ctx, chatOID)
}

func (a *ChatManagerProxyAdapter) VerifyUserExists(ctx context.Context, userOID *primitive.ObjectID) error {
	return a.repository.VerifyChatExists(ctx, userOID)
}

func (a *ChatManagerProxyAdapter) UpdateTournamentStartDate(ctx context.Context, tournamentOID *primitive.ObjectID) error {
	return a.repository.UpdateTournamentStartDate(ctx, tournamentOID)
}
