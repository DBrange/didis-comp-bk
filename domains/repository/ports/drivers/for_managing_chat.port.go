package ports

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	chat_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/chat/dao"
	chat_message_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/chat_message/dao"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ForManagingChat interface {
	WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) error) error
	ConvertToObjectID(ID string) (*primitive.ObjectID, error)
	CreateChat(ctx context.Context, matchOID *primitive.ObjectID, chatType models.CHAT) (string, error)
	CreateParticipantChats(ctx context.Context, chatOID *primitive.ObjectID, participantOIDs []*primitive.ObjectID, chatType models.CHAT) error
	GetMatchChatByID(ctx context.Context, chatID string) (*chat_dao.GetMatchChatByIDDAORes, error)
	CreateChatMessage(ctx context.Context, chatMessageDAO *chat_message_dao.CreateChatMessageDAOReq) (string, error)
	VerifyChatExists(ctx context.Context, chatOID *primitive.ObjectID) error
	VerifyUserExists(ctx context.Context, userOID *primitive.ObjectID) error
}
