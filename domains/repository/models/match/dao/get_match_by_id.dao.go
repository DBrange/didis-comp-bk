package dao

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetMatchByIDDAORes struct {
	ID           *primitive.ObjectID `bson:"_id"`
	Sport        *models.SPORT       `bson:"sport"`
	RoundID      *primitive.ObjectID `bson:"round_id"`
	Result       *primitive.ObjectID `bson:"result"`
	Winner       *primitive.ObjectID `bson:"winner"`
	TournamentID *primitive.ObjectID `bson:"tournament_id"`
	Position     int                 `bson:"position"`
	Date         *time.Time          `bson:"date"`
	// Votes        map[string]string  `bson:"votes"`
	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}
