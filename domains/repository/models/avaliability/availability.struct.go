package models

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Availability struct {
	ID                  primitive.ObjectID   `bson:"_id,omitempty"`
	DailyAvailabilities []*DailyAvailability `bson:"daily_availabilities"`
	UserID              *primitive.ObjectID   `bson:"user_id"`
	CompetitorID        *primitive.ObjectID   `bson:"competitor_id"`
	CreatedAt           time.Time            `bson:"created_at"`
	UpdatedAt           time.Time            `bson:"updated_at"`
	DeletedAt           *time.Time           `bson:"deleted_at,omitempty"`
}

type DailyAvailability struct {
	Day       string      `bson:"day"`
	TimeSlots []*TimeSlot `bson:"time_slots"`
}

type TimeSlot struct {
	TimeSlot string                     `bson:"time_slot"`
	Status   models.AVAILABILITY_STATUS `bson:"status"`
}
