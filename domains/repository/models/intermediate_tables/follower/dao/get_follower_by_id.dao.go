package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetFollowerByIDDAORes struct {
	ID           primitive.ObjectID  `bson:"_id"`
	From         primitive.ObjectID  `bson:"from"`
	ToUser       *primitive.ObjectID `bson:"to_user"`
	ToCompetitor *primitive.ObjectID `bson:"to_competitor"`
	CreatedAt    time.Time           `bson:"created_at"`
	UpdatedAt    time.Time           `bson:"updated_at"`
	DeletedAt    *time.Time          `bson:"deleted_at,omitempty"`
}
