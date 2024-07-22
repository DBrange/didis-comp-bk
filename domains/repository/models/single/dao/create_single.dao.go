package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateSingleDAOReq struct {
	StatsID        *primitive.ObjectID `bson:"competitor_stats_id"`
	AvailabilityID *primitive.ObjectID `bson:"availability_id"`
	CreatedAt      time.Time           `bson:"created_at"`
	UpdatedAt      time.Time           `bson:"updated_at"`
	DeletedAt      *time.Time          `bson:"deleted_at,omitempty"`
}

func (u *CreateSingleDAOReq) SetTimeStamp() {
	currentDate := time.Now().UTC()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = currentDate
	}
	u.UpdatedAt = currentDate
}
