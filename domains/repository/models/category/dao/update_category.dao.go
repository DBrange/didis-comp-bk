package dao

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateCategoryDAOReq struct {
	Name              *string                `bson:"name,omitempty"`
	Genre             *models.GENRE          `bson:"genre,omitempty"`
	TotalParticipants *int                   `bson:"total_participants,omitempty"`
	RangeMovement     *models.RANGE_MOVEMENT `bson:"range_movement,omitempty"`
	AverageScore      *float32               `bson:"average_score,omitempty"`
	Tournaments       *[]primitive.ObjectID  `bson:"tournaments,omitempty"`
	UpdatedAt         time.Time              `bson:"updated_at"`
}

func (u *UpdateCategoryDAOReq) RenewUpdate() {
	u.UpdatedAt = time.Now().UTC()
}
