package dto

import "github.com/DBrange/didis-comp-bk/cmd/api/models"

type GetUserForLoginDTO struct {
	ID        string                       `json:"id"`
	Password  string                       `json:"password"`
	FirstName string                       `json:"first_name"`
	LastName  string                       `json:"last_name"`
	Username  string                       `json:"username"`
	Image     string                       `json:"image"`
	Roles     []string                     `json:"roles"`
	Sports    []models.SPORT               `json:"sports"`
	Organizer *GetUserForLoginOrganizerDTO `json:"organizer"`
}

type GetUserForLoginOrganizerDTO struct {
	ID               *string        `json:"id"`
	TournamentSports []models.SPORT `json:"tournament_sports"`
	CategorySports   []models.SPORT `json:"category_sports"`
}
