package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateOrganizerDAOReq struct {
	Categories              []primitive.ObjectID `bson:"categories"`
	Tournaments              []primitive.ObjectID `bson:"tournaments"`
	AverageScore            float32              `bson:"average_score"`
	AverageTournamentePrize float32              `bson:"average_tournament_score"`
	TotalCategories         int                  `bson:"total_categories"`
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
