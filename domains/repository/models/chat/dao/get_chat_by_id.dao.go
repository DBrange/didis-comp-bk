package dao

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetChatByIDDAORes struct {
	ID       *primitive.ObjectID `bson:"_id"`
	ChatType models.CHAT         `bson:"chat_type"`
	// Status             models.CHAT_AVAILABILITY_STATUS `bson:"status"`
	AvailabilityStatus *models.CHAT_AVAILABILITY_STATUS `bson:"availability_status"`
	MatchID            *primitive.ObjectID              `bson:"match_id"`
	CreatedAt          time.Time                        `bson:"created_at"`
	UpdatedAt          time.Time                        `bson:"updated_at"`
	DeletedAt          *time.Time                       `bson:"deleted_at,omitempty"`
}
