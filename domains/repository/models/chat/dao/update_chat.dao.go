package dao

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type UpdateChatDAOReq struct {
	// Status       *models.CHAT_AVAILABILITY_STATUS `bson:"status,omitempty"`
	AvailabilityStatus models.CHAT_AVAILABILITY_STATUS `bson:"availability_status,omitempty"`
	UpdatedAt          time.Time                       `bson:"updated_at,omitempty"`
}

func (u *UpdateChatDAOReq) RenewUpdate() {
	u.UpdatedAt = time.Now().UTC()
}
