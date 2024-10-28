package dto

type GetOrganizerDataDTORes struct {
	ID                     string `json:"id"`
	AverageScore           int    `json:"average_score"`
	AverageTournamentScore int    `json:"average_tournament_score"`
	TotalCategories        int    `json:"total_categories"`
	TotalTournaments       int    `json:"total_tournaments"`
}
