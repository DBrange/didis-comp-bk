package dto

import "time"

type GetTournamentsFromCategoryDTORes struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Points      *int    `json:"points"`
	StartDate   *time.Time `json:"start_date"`
	FinishtDate *time.Time `json:"finish_date"`
}
