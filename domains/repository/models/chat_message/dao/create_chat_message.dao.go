package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateChatMessageDAOReq struct {
	ChatID    primitive.ObjectID `bson:"chat_id"`
	SenderID  primitive.ObjectID `bson:"sender_id"` // userID
	Content   string             `bson:"content"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	DeletedAt *time.Time         `bson:"deleted_at,omitempty"`
}

func (u *CreateChatMessageDAOReq) SetTimeStamp() {
	currentDate := time.Now().UTC()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = currentDate
	}
	u.UpdatedAt = currentDate
}
