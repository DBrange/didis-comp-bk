package dto

import "github.com/DBrange/didis-comp-bk/cmd/api/models"

type GetProfileDailyAvailabilityInfoByIDDTORes struct {
	Day       string                         `json:"day" validate:"day"`
	TimeSlots []*GetProfileDailyTimeSlotInfoByIDDTORes `json:"time_slots" validate:"dive"`
}

type GetProfileDailyTimeSlotInfoByIDDTORes struct {
	TimeSlot string                     `json:"time_slot" validate:"timeSlot"`
	Status   models.AVAILABILITY_STATUS `json:"status" validate:"availStatus"`
}
