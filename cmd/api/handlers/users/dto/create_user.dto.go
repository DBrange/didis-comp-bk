package dto

import (
	"time"

	dto "github.com/DBrange/didis-comp-bk/cmd/api/assets/dto/location"
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type RegisterUserDTOReq struct {
	FirstName   string                 `json:"first_name" validate:"required"`
	LastName    string                 `json:"last_name" validate:"required"`
	Username    *string                `json:"username"`
	Birthdate   *time.Time             `json:"birthdate"`
	Password    *string                `json:"password"`
	Email       string                 `json:"email" validate:"required"`
	Phone       *string                `json:"phone"`
	Image       *string                `json:"image"`
	Active      bool                   `json:"active"`
	AccessLevel *int16                 `json:"access_level"`
	Genre       []models.GENRE         `json:"genre" validate:"dive,genre"`
	Role        []models.ROLE          `json:"role"`
	Location    *dto.CreateLocationDTO `json:"location"`
	ScheduleID  *string                `json:"schedule_id"`
	PaymentID   *string                `json:"payment_id"`
}
