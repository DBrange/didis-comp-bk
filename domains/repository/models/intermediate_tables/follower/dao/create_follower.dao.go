package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateFollowerDAOReq struct {
	From         primitive.ObjectID `bson:"from"`
	ToUser       *primitive.ObjectID `bson:"to_user"`
	ToCompetitor *primitive.ObjectID `bson:"to_competitor"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
	DeletedAt    *time.Time         `bson:"deleted_at,omitempty"`
}

func (u *CreateFollowerDAOReq) SetTimeStamp() {
	currentDate := time.Now().UTC()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = currentDate
	}
	u.UpdatedAt = currentDate
}
