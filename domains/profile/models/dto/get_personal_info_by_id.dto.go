package dto

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type GetPersonalInfoByIDDTORes struct {
	ID        string                          `json:"id"`
	FirstName string                          `json:"first_name"`
	LastName  string                          `json:"last_name"`
	Username  *string                         `json:"username"`
	Birthdate *string                         `json:"birthdate"`
	Password  *string                         `json:"password"`
	Email     string                          `json:"email"`
	Phone     *string                         `json:"phone"`
	Image     *string                         `json:"image"`
	Genre     models.GENRE                    `json:"genre"`
	Location  *GetPersonalInfoLocationByIDRes `json:"location"`
}

type GetPersonalInfoLocationByIDRes struct {
	ID      string  `json:"id"`
	State   *string `json:"state"`
	Country *string `json:"country"`
	City    *string `json:"city"`
	Lat     *string `json:"lat"`
	Long    *string `json:"long"`
}
