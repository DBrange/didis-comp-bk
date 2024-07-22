package dao

import (
	"time"
)

type UpdateChatMessageDAOReq struct {
	Content   *string    `bson:"content,omitempty"`
	UpdatedAt time.Time `bson:"updated_at"`
}

func (u *UpdateChatMessageDAOReq) RenewUpdate() {
	u.UpdatedAt = time.Now().UTC()
}
