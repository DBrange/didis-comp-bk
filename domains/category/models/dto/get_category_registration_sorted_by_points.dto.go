package dto

type GetCategoryRegistrationSortedByPointsDTORes struct {
	CompetitorID    string `json:"competitor_id"`
	CurrentPosition *int   `json:"current_position"`
}
