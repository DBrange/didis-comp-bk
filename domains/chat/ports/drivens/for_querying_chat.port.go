package ports

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/chat/models/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ForQueryingChat interface {
	WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) error) error
	ConvertToObjectID(ID string) (*primitive.ObjectID, error)
	CreateChat(ctx context.Context, matchID string, chatType models.CHAT) (string, error)
	CreateParticipantChats(ctx context.Context, chatID string, participantIDs []string, chatType models.CHAT) error
	GetMatchChatByID(ctx context.Context, chatID string) (*dto.GetMatchChatByIDDTORes, error)
	CreateChatMessage(ctx context.Context, chatMessageDTO *dto.CreateChatMessageDTOReq) (string, error)
	VerifyChatExists(ctx context.Context, chatID string) error
	VerifyUserExists(ctx context.Context, userID string) error
}
