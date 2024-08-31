package dao

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetParticipantChatByIDDAORes struct {
	UserID             *primitive.ObjectID             `bson:"user_id"`
	CompetitorID       *primitive.ObjectID             `bson:"competitor_id"`
	ChatID             *primitive.ObjectID             `bson:"chat_id"`
	AvailabilityStatus models.CHAT_AVAILABILITY_STATUS `bson:"availability_status"`
	common.GetBaseDAO  `bson:",inline"`
}
