package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateTournamentRegistrationDAOReq struct {
	CompetitorID primitive.ObjectID `bson:"competitor_id"`
	TournamentID primitive.ObjectID `bson:"tournament_id"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
	DeletedAt    *time.Time         `bson:"deleted_at,omitempty"`
}

func (u *CreateTournamentRegistrationDAOReq) SetTimeStamp() {
	currentDate := time.Now().UTC()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = currentDate
	}
	u.UpdatedAt = currentDate
}
