package dto

import "github.com/DBrange/didis-comp-bk/cmd/api/models"

type UpdateAvailabilityDTOReq struct {
	DailyAvailabilities []*UpdateDailyAvailabilityDTOReq `json:"daily_availabilities" validate:"dive"`
}

type UpdateDailyAvailabilityDTOReq struct {
	Day       string                  `json:"day" validate:"day"`
	TimeSlots []*UpdateTimeSlotDTOReq `json:"time_slots" validate:"dive"`
}

type UpdateTimeSlotDTOReq struct {
	TimeSlot string                     `json:"time_slot" validate:"timeSlot"`
	Status   models.AVAILABILITY_STATUS `json:"status" validate:"availStatus"`
}

