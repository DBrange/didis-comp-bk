package dto

type CreateCompetitorMatchDTOReq struct {
	CompetitorID *string `json:"competitor_id"`
	Position     int     `json:"position"`
	MatchID      string  `json:"match_id"`
}
