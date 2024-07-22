package dao

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateCompetitorDAOReq struct {
	SingleID  *primitive.ObjectID `bson:"single_id"`
	DoubleID  *primitive.ObjectID `bson:"double_id"`
	TeamID    *primitive.ObjectID `bson:"team_id"`
	Sport     models.SPORT        `bson:"sport"`
	CreatedAt time.Time           `bson:"created_at"`
	UpdatedAt time.Time           `bson:"updated_at"`
	DeletedAt *time.Time          `bson:"deleted_at,omitempty"`
}

func (u *CreateCompetitorDAOReq) SetTimeStamp() {
	currentDate := time.Now().UTC()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = currentDate
	}
	u.UpdatedAt = currentDate
}
