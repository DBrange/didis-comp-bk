package dto

type GetProfileInfoInCategoryDTORes struct {
	ID              string                                 `json:"id"`
	Points          int                                    `json:"points"`
	CurrentPosition int                                    `json:"current_position"`
	Users           []*GetProfileInfoInCategoryUsersDTORes `json:"users"`
	GuestUsers      []*GetProfileInfoInCategoryUsersDTORes `json:"guest_users"`
	CompetitorStats *GetProfileInfoInCategoryStatsDTORes   `json:"competitor_stats"`
}

type GetProfileInfoInCategoryUsersDTORes struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Image     string `json:"image"`
}

type GetProfileInfoInCategoryStatsDTORes struct {
	ID             string   `json:"id"`
	TotalWins      int      `json:"total_wins"`
	TotalLosses    int      `json:"total_losses"`
	MoneyEarned    float64  `json:"money_earned"`
	TournamentsWon []string `json:"tournaments_won"`
}
