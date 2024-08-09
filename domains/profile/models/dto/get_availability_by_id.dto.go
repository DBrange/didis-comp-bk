package dto

import "github.com/DBrange/didis-comp-bk/cmd/api/models"

type GetAvailabilityByIDDTORes struct {
	DailyAvailabilities []*GetDailyAvailabilityByIDDTORes `json:"daily_availabilities"`
}

type GetDailyAvailabilityByIDDTORes struct {
	Day       models.DAY                        `json:"day"`
	TimeSlots []*GetDailyTimeSlotByIDDTORes `json:"time_slots"`
}

type GetDailyTimeSlotByIDDTORes struct {
	TimeSlot string                     `json:"time_slot"`
	Status   models.AVAILABILITY_STATUS `json:"status"`
}
