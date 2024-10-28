package dto

type GetRoundWithCompetitorsDTORes struct {
	ID            string   `json:"_id"`
	TotalPrize    float64  `json:"total_prize"`
	Points        int      `json:"points"`
	CompetitorIDs []string `json:"competitor_ids"`
}
