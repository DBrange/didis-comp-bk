package dto

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type RegisterUserDTOReq struct {
	FirstName   string             `json:"first_name"`
	LastName    string             `json:"last_name"`
	Username    *string            `json:"username"`
	Birthdate   *time.Time         `json:"birthdate"`
	Password    *string            `json:"password"`
	Email       string             `json:"email"`
	Phone       *string            `json:"phone"`
	Image       *string            `json:"image"`
	Active      bool               `json:"active"`
	AccessLevel *int16             `json:"access_level"`
	Genre       []models.GENRE     `json:"genre"`
	Role        []models.ROLE      `json:"role"`
	Location    *CreateLocationDTOReq `json:"location"`
	ScheduleID  *string            `json:"schedule_id"`
	PaymentID   *string            `json:"payment_id"`
}

type CreateLocationDTOReq struct {
	State   *string `json:"state"`
	Country *string `json:"country"`
	City    *string `json:"city"`
	Lat     *string `json:"lat"`
	Long    *string `json:"long"`
}
