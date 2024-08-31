package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetDoubleEliminationByIDDAORes struct {
	ID         primitive.ObjectID    `bson:"_id"`
	Matches    []*primitive.ObjectID `bson:"matches"`
	Rounds     []*primitive.ObjectID `bson:"rounds"`
	TotalPrize float64               `bson:"total_prize"`
	Points     int                   `bson:"points"`
	CreatedAt  time.Time             `bson:"created_at"`
	UpdatedAt  time.Time             `bson:"updated_at"`
	DeletedAt  *time.Time            `bson:"deleted_at,omitempty"`
}
