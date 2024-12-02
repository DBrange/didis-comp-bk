package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetCompetitorStatsByIDDAORes struct {
	ID             primitive.ObjectID   `bson:"_id"`
	TotalWins      int                  `bson:"total_wins"`
	TotalLosses    int                  `bson:"total_losses"`
	MoneyEarned    float64              `bson:"money_earned"`
	Matches        []primitive.ObjectID `bson:"matches"`
	TournamentsWon []primitive.ObjectID `bson:"tournaments_won"`
	CompetitorID   primitive.ObjectID   `bson:"competitor_id"`
	CreatedAt      time.Time            `bson:"created_at"`
	UpdatedAt      time.Time            `bson:"updated_at"`
	DeletedAt      *time.Time           `bson:"deleted_at,omitempty"`
}
