package dao

import "github.com/DBrange/didis-comp-bk/cmd/api/models"

type GetAvailabilityInfoByIDDAORes struct {
	DailyAvailabilities []*GetDailyAvailabilityInfoByIDDAORes `bson:"daily_availabilities"`
}

type GetDailyAvailabilityInfoByIDDAORes struct {
	Day       string                            `bson:"day"`
	TimeSlots []*GetDailyTimeSlotInfoByIDDAORes `bson:"time_slots"`
}

type GetDailyTimeSlotInfoByIDDAORes struct {
	TimeSlot string                     `bson:"time_slot"`
	Status   models.AVAILABILITY_STATUS `bson:"status"`
}
