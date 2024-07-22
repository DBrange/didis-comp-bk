package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetOpinionByIDDAORes struct {
	ID           primitive.ObjectID `bson:"_id"`
	Commentary   string             `bson:"commentary"`
	Score        float32            `bson:"score"`
	Anonymous    bool               `bson:"anonymous"`
	TournamentID primitive.ObjectID `bson:"tournament_id"`
	LeagueID     primitive.ObjectID `bson:"league_id"`
	OrganizerID  primitive.ObjectID `bson:"organizer_id"`
	TeamID       primitive.ObjectID `bson:"team_id"`
	UserID       primitive.ObjectID `bson:"user_id"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
	DeletedAt    *time.Time         `bson:"deleted_at,omitempty"`
}
