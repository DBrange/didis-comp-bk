package models


type GetAvailabilityByIDDTORes struct {
	DailyAvailabilities []*GetDailyAvailabilityByIDDTORes `json:"daily_availabilities"`
}

type GetDailyAvailabilityByIDDTORes struct {
	Day       DAY                        `json:"day"`
	TimeSlots []*GetDailyTimeSlotByIDDTORes `json:"time_slots"`
}

type GetDailyTimeSlotByIDDTORes struct {
	TimeSlot string                     `json:"time_slot"`
	Status   AVAILABILITY_STATUS `json:"status"`
}
