package dto

type UpdateCompetitorMatchDTOReq struct {
	MatchID      string `json:"match_id"`
	CompetitorID *string `json:"competitor_id"`
	Position     int    `json:"position"`
}
