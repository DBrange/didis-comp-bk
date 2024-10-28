package dto

import "github.com/DBrange/didis-comp-bk/cmd/api/models"

type GetUserForLoginDTO struct {
	ID          string         `json:"id"`
	Password    string         `json:"password"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	Username    string         `json:"username"`
	Image       string         `json:"image"`
	Roles       []string       `json:"roles"`
	Sports      []models.SPORT `json:"sports"`
	OrganizerID *string        `json:"organizer_id"`
}
