package dto

type CreateTournamentRegistrationDTOReq struct {
	TournamentID string `json:"tournament_id"`
	CompetitorID string `json:"competitor_id"`
}
