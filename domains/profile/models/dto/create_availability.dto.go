package dto

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type CreateAvailabilityDTOReq struct {
	ID                  string                           `json:"_id,omitempty"`
	DailyAvailabilities []*CreateDailyAvailabilityDTOReq `json:"daily_availabilities"`
	UserID              *string                          `json:"user_id"`
	CompetitorID        *string                          `json:"competitor_id"`
	CreatedAt           time.Time                        `json:"created_at"`
	UpdatedAt           time.Time                        `json:"updated_at"`
	DeletedAt           *time.Time                       `json:"deleted_at,omitempty"`
}

type CreateDailyAvailabilityDTOReq struct {
	Day       models.DAY              `json:"day"`
	TimeSlots []*CreateTimeSlotDTOReq `json:"time_slots"`
}

type CreateTimeSlotDTOReq struct {
	TimeSlot string                     `json:"time_slot"`
	Status   models.AVAILABILITY_STATUS `json:"status"`
}
