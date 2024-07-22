package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateOrganizerDAOReq struct {
	Leagues                 []primitive.ObjectID `bson:"leagues"`
	AverageScore            float32              `bson:"average_score"`
	AverageTournamentePrize float32              `bson:"average_tournament_score"`
	TotalLeagues            int                  `bson:"total_leagues"`
	TotalTournaments        int                  `bson:"total_tournaments"`
	UserID                  primitive.ObjectID   `bson:"user_id"`
	CreatedAt               time.Time            `bson:"created_at"`
	UpdatedAt               time.Time            `bson:"updated_at"`
	DeletedAt               *time.Time           `bson:"deleted_at,omitempty"`
}

func (u *CreateOrganizerDAOReq) SetTimeStamp() {
	currentDate := time.Now().UTC()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = currentDate
	}
	u.UpdatedAt = currentDate
}

func (u *CreateOrganizerDAOReq) RenewUpdate() {
	u.UpdatedAt = time.Now().UTC()
}
