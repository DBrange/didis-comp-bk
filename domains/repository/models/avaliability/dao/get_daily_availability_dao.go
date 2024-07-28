package dao

import "github.com/DBrange/didis-comp-bk/cmd/api/models"

type GetAvailabilityByIDDAORes struct {
	DailyAvailabilities []*GetDailyAvailabilityByIDDAORes `bson:"daily_availabilities"`
}

type GetDailyAvailabilityByIDDAORes struct {
	Day       string                        `bson:"day"`
	TimeSlots []*GetDailyTimeSlotByIDDAORes `bson:"time_slots"`
}

type GetDailyTimeSlotByIDDAORes struct {
	TimeSlot string                     `bson:"time_slot"`
	Status   models.AVAILABILITY_STATUS `bson:"status"`
}
