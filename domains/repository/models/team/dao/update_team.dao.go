package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateTeamDAOReq struct {
	TotalMembers   *int                  `bson:"total_members,omitempty"`
	Image          *string               `bson:"image,omitempty"`
	AverageScore   *float32              `bson:"average_score,omitempty"`
	Name           *float32              `bson:"average_score,omitempty"`
	StatsID        *primitive.ObjectID   `bson:"competitor_stats_id,omitempty"`
	Admins         *[]primitive.ObjectID `bson:"admins,omitempty"`
	AvailabilityID *primitive.ObjectID   `bson:"availability_id,omitempty"`
	UpdatedAt      time.Time             `bson:"updated_at,omitempty"`
}

func (u *UpdateTeamDAOReq) RenewUpdate() {
	u.UpdatedAt = time.Now().UTC()
}
