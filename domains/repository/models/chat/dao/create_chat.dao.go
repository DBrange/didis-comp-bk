package dao

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateChatDAOReq struct {
	ChatType models.CHAT `bson:"chat_type"`
	// Status       models.CHAT_AVAILABILITY_STATUS `bson:"status"`
	AvailabilityStatus models.CHAT_AVAILABILITY_STATUS `bson:"availability_status"`
	MatchID            *primitive.ObjectID             `bson:"match_id"`
	CreatedAt          time.Time                       `bson:"created_at"`
	UpdatedAt          time.Time                       `bson:"updated_at"`
	DeletedAt          *time.Time                      `bson:"deleted_at,omitempty"`
}

func (u *CreateChatDAOReq) SetTimeStamp() {
	currentDate := time.Now().UTC()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = currentDate
	}
	u.UpdatedAt = currentDate
}
