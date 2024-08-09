package dto

import "github.com/DBrange/didis-comp-bk/cmd/api/models"

type CreateGuestUserDTOReq struct {
	FirstName string       `json:"first_name" validate:"required"`
	LastName  string       `json:"last_name" validate:"required"`
	Email     string       `json:"email" validate:"required,email"`
	Image     *string      `json:"image"`
	Genre     models.GENRE `json:"genre" validate:"required,genre"`
}
