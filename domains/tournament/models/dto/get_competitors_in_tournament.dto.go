package dto

type GetCompetitorsInTournamentDTORes struct {
	CompetitorID    string                                  `json:"id"`
	CurrentPosition *int                                    `json:"current_position"`
	Users           []*GetCompetitorsInTournamentUserDTORes `json:"users"`
	GuestUsers      []*GetCompetitorsInTournamentUserDTORes `json:"guest_users"`
}

type GetCompetitorsInTournamentUserDTORes struct {
	ID        string `json:"_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Image     string `json:"image"`
}
