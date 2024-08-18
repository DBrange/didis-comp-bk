package dto

type UpdateCompetitorStatsDAOReq struct {
	TotalWins      *int     `bson:"total_wins,omitempty"`
	TotalLosses    *int     `bson:"total_losses,omitempty"`
	MoneyEarned    *float64 `bson:"money_earned,omitempty"`
	Matches        string   `bson:"matches,omitempty"`
	TournamentsWon string   `bson:"tournaments_won,omitempty"`
}
