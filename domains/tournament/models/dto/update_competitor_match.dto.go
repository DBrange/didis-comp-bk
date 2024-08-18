package dto

type UpdateCompetitorMatchDTOReq struct {
	MatchID *string `json:"match_id,omitempty"`
	CompetitorID *string `json:"competitor_id"`
	Position     *int    `json:"position,omitempty"`
}
