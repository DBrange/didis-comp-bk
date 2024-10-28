package models

type UpdateAvailabilityDTOReq struct {
	DailyAvailabilities []*UpdateDailyAvailabilityDTOReq `json:"daily_availabilities" validate:"dive"`
}

type UpdateDailyAvailabilityDTOReq struct {
	Day       DAY                   `json:"day" validate:"day"`
	TimeSlots []*UpdateTimeSlotDTOReq `json:"time_slots" validate:"dive"`
}

type UpdateTimeSlotDTOReq struct {
	TimeSlot string                     `json:"time_slot" validate:"timeSlot"`
	Status   AVAILABILITY_STATUS `json:"status" validate:"availStatus"`
}

