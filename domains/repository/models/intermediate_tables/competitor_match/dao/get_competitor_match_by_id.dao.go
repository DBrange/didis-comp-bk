package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetCompetitorMatchByIDDAORes struct {
	ID           primitive.ObjectID `bson:"_id"`
	CompetitorID primitive.ObjectID `bson:"competitor_id"`
	MatchID      primitive.ObjectID `bson:"match_id"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
	DeletedAt    *time.Time         `bson:"deleted_at,omitempty"`
}
