package dao

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetAvailabilityByIDDAORes struct {
	ID                  *primitive.ObjectID               `bson:"_id"`
	DailyAvailabilities []*GetDailyAvailabilityByIDDAORes `bson:"daily_availabilities"`
}

type GetDailyAvailabilityByIDDAORes struct {
	Day       models.DAY                    `bson:"day"`
	TimeSlots []*GetDailyTimeSlotByIDDAORes `bson:"time_slots"`
}

type GetDailyTimeSlotByIDDAORes struct {
	TimeSlot string                     `bson:"time_slot"`
	Status   models.AVAILABILITY_STATUS `bson:"status"`
}
