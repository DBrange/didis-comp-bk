package models

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID                primitive.ObjectID
	Name              string
	Genre             models.GENRE
	TotalParticipants int
	RangeMovement     models.RANGE_MOVEMENT
	AverageScore      float32
	Sport             models.SPORT
	OrganizerID       primitive.ObjectID
	Tournaments       []primitive.ObjectID
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
}
