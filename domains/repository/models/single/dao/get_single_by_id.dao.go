package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetSingleByIDDAORes struct {
	ID             *primitive.ObjectID `bson:"_id"`
	CreatedAt      time.Time           `bson:"created_at"`
	UpdatedAt      time.Time           `bson:"updated_at"`
	DeletedAt      *time.Time          `bson:"deleted_at,omitempty"`
}
