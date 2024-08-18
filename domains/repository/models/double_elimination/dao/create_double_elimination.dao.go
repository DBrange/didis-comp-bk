package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateDoubleEliminationDAOReq struct {
	Matches   []*primitive.ObjectID `bson:"matches"`
	Rounds    []*primitive.ObjectID `bson:"rounds"`
	CreatedAt time.Time            `bson:"created_at"`
	UpdatedAt time.Time            `bson:"updated_at"`
	DeletedAt *time.Time           `bson:"deleted_at,omitempty"`
}

func (u *CreateDoubleEliminationDAOReq) SetTimeStamp() {
	currentDate := time.Now().UTC()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = currentDate
	}
	u.UpdatedAt = currentDate
}
