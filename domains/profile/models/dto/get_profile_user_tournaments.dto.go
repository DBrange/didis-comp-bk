package dto

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type GetProfileUserTournamentsDTORes struct {
	Tournaments []*GetProfileUserTournamentDTORes `json:"tournaments"`
	Total       int                               `json:"total"`
}
type GetProfileUserTournamentDTORes struct {
	ID        string                                 `json:"id"`
	Name      string                                 `json:"name"`
	Location  *GetLocationByIDDTORes                 `json:"location"`
	Organizer *GetUserTournamentsOrganizerDTO        `json:"organizer"`
	Matches   []*GetProfileUserTournamentMatchDTORes `json:"matches"`
}

type GetProfileUserTournamentMatchDTORes struct {
	ID          string                               `json:"id"`
	Result      string                               `json:"result"`
	Winner      *string                              `json:"winner"`
	Date        *time.Time                           `json:"date"`
	Round       *GetProfileUserTournamentRoundDTORes `json:"round"`
	Competitors []*GetCompetitorsInTournamentDTORes  `json:"competitors"`
}

type GetProfileUserTournamentRoundDTORes struct {
	ID    string       `json:"id"`
	Round models.ROUND `json:"round"`
}

type GetUserTournamentsOrganizerDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type GetCompetitorsInTournamentDTORes struct {
	CompetitorID    *string                                `json:"id"`
	CurrentPosition *int                                   `json:"current_position"`
	Users           []*GetProfileInfoInCategoryUsersDTORes `json:"users"`
	GuestUsers      []*GetProfileInfoInCategoryUsersDTORes `json:"guest_users"`
}
