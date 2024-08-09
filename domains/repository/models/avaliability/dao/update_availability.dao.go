package dao

import "github.com/DBrange/didis-comp-bk/cmd/api/models"

type UpdateAvailabilityDAOReq struct {
	DailyAvailabilities []*UpdateDailyAvailabilityDAOReq `bson:"daily_availabilities" validate:"dive"`
}

type UpdateDailyAvailabilityDAOReq struct {
	Day       models.DAY              `bson:"day" validate:"day"`
	TimeSlots []*UpdateTimeSlotDAOReq `bson:"time_slots" validate:"dive"`
}

type UpdateTimeSlotDAOReq struct {
	TimeSlot string                     `bson:"time_slot" validate:"timeSlot"`
	Status   models.AVAILABILITY_STATUS `bson:"status" validate:"availStatus"`
}
