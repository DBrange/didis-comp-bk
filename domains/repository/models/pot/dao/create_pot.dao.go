package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreatePotDAOReq struct {
	TournamentID primitive.ObjectID   `bson:"tournament_id"`
	Competitors  []primitive.ObjectID `bson:"competitors"`
	Position     int                  `bson:"position"`
	CreatedAt    time.Time            `bson:"created_at"`
	UpdatedAt    time.Time            `bson:"updated_at"`
	DeletedAt    *time.Time           `bson:"deleted_at,omitempty"`
}

func (u *CreatePotDAOReq) SetTimeStamp() {
	currentDate := time.Now().UTC()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = currentDate
	}
	u.UpdatedAt = currentDate
}
