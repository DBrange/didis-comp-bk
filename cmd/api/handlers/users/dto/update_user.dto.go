package dto

import (
	"time"

	dto "github.com/DBrange/didis-comp-bk/cmd/api/assets/dto/location"
)

type UpdateUserDTOReq struct {
	// ID         string                 `json:"id,omitempty" validate:"omitempty"`
	FirstName  *string                `json:"first_name,omitempty" validate:"omitempty,min=2"`
	LastName   *string                `json:"last_name,omitempty" validate:"omitempty,min=2"`
	Username   *string                `json:"username,omitempty" validate:"omitempty,min=3"`
	Birthdate  *time.Time             `json:"birthdate,omitempty" validate:"omitempty"`
	Phone      *string                `json:"phone,omitempty" validate:"omitempty,e164"`
	Image      *string                `json:"image,omitempty" validate:"omitempty,url"`
	Location   *dto.UpdateLocationDTO `json:"location,omitempty" validate:"omitempty"`
	ScheduleID *string                `json:"schedule_id,omitempty" validate:"omitempty"`
}