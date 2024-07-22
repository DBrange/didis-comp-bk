package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateCompetitorStatsDAOReq struct {
	TotalWins      *int                  `bson:"total_wins,omitempty"`
	TotalLosses    *int                  `bson:"total_losses,omitempty"`
	MoneyEarned    *float64              `bson:"money_earned,omitempty"`
	Matches        *[]primitive.ObjectID `bson:"matches,omitempty"`
	TournamentsWon *[]primitive.ObjectID `bson:"tournaments_won,omitempty"`
	UpdatedAt      time.Time             `bson:"updated_at,omitempty"`
}

func (u *UpdateCompetitorStatsDAOReq) RenewUpdate() {
	u.UpdatedAt = time.Now().UTC()
}
