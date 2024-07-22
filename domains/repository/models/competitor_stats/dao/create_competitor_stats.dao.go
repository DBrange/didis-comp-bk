package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateCompetitorStatsDAOReq struct {
	TotalWins      int                  `bson:"total_wins"`
	TotalLosses    int                  `bson:"total_losses"`
	MoneyEarned    float64              `bson:"money_earned"`
	Matches        []primitive.ObjectID `bson:"matches"`
	TournamentsWon []primitive.ObjectID `bson:"tournaments_won"`
	CreatedAt      time.Time            `bson:"created_at"`
	UpdatedAt      time.Time            `bson:"updated_at"`
	DeletedAt      *time.Time           `bson:"deleted_at,omitempty"`
}

func (u *CreateCompetitorStatsDAOReq) SetTimeStamp() {
	currentDate := time.Now().UTC()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = currentDate
	}
	u.UpdatedAt = currentDate
}
