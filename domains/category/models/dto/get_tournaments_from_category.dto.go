package dto

import "time"

type GetTournamentsFromCategoryDTORes struct {
	Total       int                                           `json:"total"`
	Tournaments []*GetTournamentsFromCategoryTournamentDTORes `json:"tournaments"`
}
type GetTournamentsFromCategoryTournamentDTORes struct {
	ID           string                 `json:"id"`
	Name         string                 `json:"name"`
	Points       *int                   `json:"points"`
	Location     *GetLocationByIDDTORes `json:"location"`
	TotalPrize   float64                `json:"total_prize"`
	AverageScore int                    `json:"average_score"`
	StartDate    *time.Time             `json:"start_date"`
	FinishtDate  *time.Time             `json:"finish_date"`
}
