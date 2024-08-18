package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateCompetitorMatchDAOReq struct {
	CompetitorID *primitive.ObjectID `bson:"competitor_id"`
	MatchID      *primitive.ObjectID `bson:"match_id"`
	Position     int                `bson:"position"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
	DeletedAt    *time.Time         `bson:"deleted_at,omitempty"`
}

func (u *CreateCompetitorMatchDAOReq) SetTimeStamp() {
	currentDate := time.Now().UTC()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = currentDate
	}
	u.UpdatedAt = currentDate
}
