package dto

import (
	"time"

	location_dto "github.com/DBrange/didis-comp-bk/domains/location/models/dto"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type GetUserByIDDTORes struct {
	ID          string                              `json:"id" validate:"required"`
	FirstName   string                              `json:"first_name" validate:"required"`
	LastName    string                              `json:"last_name" validate:"required"`
	Username    *string                             `json:"username"`
	Birthdate   *time.Time                          `json:"birthdate"`
	Password    *string                             `json:"password"`
	Email       string                              `json:"email" validate:"required"`
	Phone       *string                             `json:"phone"`
	Image       *string                             `json:"image"`
	Active      bool                                `json:"active" validate:"required"`
	AccessLevel *int16                              `json:"access_level"`
	Genre       []models.GENRE                      `json:"genre" validate:"dive,genre"`
	Role        []models.ROLE                       `json:"role"`
	Location    *location_dto.GetLocationByIDDTORes `json:"location"`
	ScheduleID  *string                             `json:"schedule_id"`
	PaymentID   *string                             `json:"payment_id"`
}
