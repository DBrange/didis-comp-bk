package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetTournamentGroupByIDDAORes struct {
	ID           primitive.ObjectID   `bson:"_id"`
	TournamentID primitive.ObjectID   `bson:"tournament_id"`
	Competitors  []primitive.ObjectID `bson:"competitors"`
	Matches      []primitive.ObjectID `bson:"matches"`
	CreatedAt    time.Time            `bson:"created_at"`
	UpdatedAt    time.Time            `bson:"updated_at"`
	DeletedAt    *time.Time           `bson:"deleted_at,omitempty"`
}