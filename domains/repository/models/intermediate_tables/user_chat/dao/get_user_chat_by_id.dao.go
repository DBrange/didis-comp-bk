package dao

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetUserChatByIDDAORes struct {
	UserID primitive.ObjectID `bson:"user_id"`
	ChatID primitive.ObjectID `bson:"chat_id"`
	common.GetBaseDAO `bson:",inline"`
}
