package dto

type GetTournamentInfoToFinaliseItDTORes struct {
	CategoryID string  `bson:"category_id"`
	TotalPrize float64 `bson:"total_prize"`
	Points     int     `bson:"points"`
}
