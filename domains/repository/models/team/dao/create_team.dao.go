package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateTeamDAOReq struct {
	TotalMembers   int                  `bson:"total_members"`
	Image          string               `bson:"image"`
	AverageScore   float32              `bson:"average_score"`
	Name           float32              `bson:"average_score"`
	Admins         []primitive.ObjectID `bson:"admins"`
	CreatedAt      time.Time            `bson:"created_at"`
	UpdatedAt      time.Time            `bson:"updated_at"`
	DeletedAt      *time.Time           `bson:"deleted_at,omitempty"`
}

func (u *CreateTeamDAOReq) SetTimeStamp() {
	currentDate := time.Now().UTC()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = currentDate
	}
	u.UpdatedAt = currentDate
}
