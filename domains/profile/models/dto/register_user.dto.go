package dto

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type RegisterUserDTOReq struct {
	FirstName string                `json:"first_name" validate:"required,min=2"`
	LastName  string                `json:"last_name" validate:"required,min=2"`
	Username  *string               `json:"username"`
	Birthdate *time.Time            `json:"birthdate"`
	Password  *string               `json:"password"`
	Email     string                `json:"email" validate:"required,email"`
	Phone     *string               `json:"phone"`
	Image     *string               `json:"image"`
	Active    bool                  `json:"active" `
	Genre     models.GENRE          `json:"genre" validate:"genre"`
	Location  *CreateLocationDTOReq `json:"location,omitempty"`
	Organizer bool                  `json:"organizer"`
}

type CreateLocationDTOReq struct {
	State   *string `json:"state"`
	Country *string `json:"country"`
	City    *string `json:"city"`
	Lat     *string `json:"lat"`
	Long    *string `json:"long"`
}
