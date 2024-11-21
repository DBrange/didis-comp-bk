package dto

import "github.com/DBrange/didis-comp-bk/cmd/api/models"

type GetDayTimeSlotDTORes struct {
	ID                  string                                   `json:"_id"`
	DailyAvailabilities []*models.GetDailyAvailabilityByIDDTORes `json:"daily_availabilities"`
}
