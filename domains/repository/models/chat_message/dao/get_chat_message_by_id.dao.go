package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetChatMessageByIDDAORes struct {
	ID        *primitive.ObjectID `bson:"_id"`
	Content   string             `bson:"content"`
	SenderID  *primitive.ObjectID `bson:"sender_id"` // userID
	ChatID    *primitive.ObjectID `bson:"chat_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	DeletedAt *time.Time         `bson:"deleted_at,omitempty"`
}
