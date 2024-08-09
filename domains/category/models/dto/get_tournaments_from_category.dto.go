package dto

type GetTournamentsFromCategoryDTORes struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Points      *int    `json:"points"`
	StartDate   *string `json:"start_date"`
	FinishtDate *string `json:"finish_date"`
}
