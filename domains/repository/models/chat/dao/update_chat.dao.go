package dao

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateChatDAOReq struct {
	Name         *string               `bson:"name,omitempty"`
	Status       *models.CHAT_STATUS   `bson:"status,omitempty"`
	Participants *[]primitive.ObjectID `bson:"participants,omitempty"`
	UpdatedAt    time.Time            `bson:"updated_at,omitempty"`
}

func (u *UpdateChatDAOReq) RenewUpdate() {
	u.UpdatedAt = time.Now().UTC()
}
