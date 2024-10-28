package dto

type GetCompetitorsInTournamentDTORes struct {
	Competitors []*GetCompetitorsInTournamentCompetitorDTORes `json:"competitors"`
	Total       int                                           `json:"total"`
}
type GetCompetitorsInTournamentCompetitorDTORes struct {
	CompetitorID    string                                  `json:"id"`
	CurrentPosition *int                                    `json:"current_position"`
	Users           []*GetCompetitorsInTournamentUserDTORes `json:"users"`
	GuestUsers      []*GetCompetitorsInTournamentUserDTORes `json:"guest_users"`
}

type GetCompetitorsInTournamentUserDTORes struct {
	ID        string  `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Image     string  `json:"image"`
	Username  *string `json:"username"`
}
