package dao

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateCategoryDAOReq struct {
	Name              string                 `bson:"name"`
	Genre             models.GENRE           `bson:"genre"`
	TotalParticipants int                    `bson:"total_participants"`
	RangeMovement     models.RANGE_MOVEMENT  `bson:"range_movement"`
	AverageScore      float32                `bson:"average_score"`
	Sport             models.SPORT           `bson:"sport"`
	CompetitorType    models.COMPETITOR_TYPE `bson:"competitor_type"`
	OrganizerID       primitive.ObjectID     `bson:"organizer_id"`
	Tournaments       []primitive.ObjectID   `bson:"tournaments"`
	CreatedAt         time.Time              `bson:"created_at"`
	UpdatedAt         time.Time              `bson:"updated_at"`
	DeletedAt         *time.Time             `bson:"deleted_at,omitempty"`
}

func (u *CreateCategoryDAOReq) SetTimeStamp() {
	currentDate := time.Now().UTC()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = currentDate
	}
	u.UpdatedAt = currentDate
}
