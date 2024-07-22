package dto

import "github.com/DBrange/didis-comp-bk/cmd/api/models"

type ModifyProfileAvailabilityDTOReq struct {
	DailyAvailabilities []*ModifyProfileDailyAvailabilityDTOReq `json:"daily_availabilities" validate:"dive"`
}

type ModifyProfileDailyAvailabilityDTOReq struct {
	Day       string                         `json:"day" validate:"day"`
	TimeSlots []*ModifyProfileTimeSlotDTOReq `json:"time_slots" validate:"dive"`
}

type ModifyProfileTimeSlotDTOReq struct {
	TimeSlot string                     `json:"time_slot" validate:"timeSlot"`
	Status   models.AVAILABILITY_STATUS `json:"status" validate:"availStatus"`
}
