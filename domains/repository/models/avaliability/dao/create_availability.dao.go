package dao

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateAvailability struct {
	ID                  primitive.ObjectID        `bson:"_id,omitempty"`
	DailyAvailabilities []*CreateDailyAvailability `bson:"daily_availabilities"`
	UserID              *primitive.ObjectID       `bson:"user_id"`
	CompetitorID        *primitive.ObjectID       `bson:"competitor_id"`
	CreatedAt           time.Time                 `bson:"created_at"`
	UpdatedAt           time.Time                 `bson:"updated_at"`
	DeletedAt           *time.Time                `bson:"deleted_at,omitempty"`
}

type CreateDailyAvailability struct {
	Day       models.DAY            `bson:"day"`
	TimeSlots []*CreateTimeSlot `bson:"time_slots"`
}

type CreateTimeSlot struct {
	TimeSlot string                     `bson:"time_slot"`
	Status   models.AVAILABILITY_STATUS `bson:"status"`
}
