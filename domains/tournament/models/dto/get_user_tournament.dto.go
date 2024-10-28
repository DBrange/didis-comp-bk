package dto

import (
	"time"
)

type GetUserTournamentsDTORes struct {
	Tournaments []*GetUserTournamentDTORes `json:"tournaments"`
	Total       int                        `json:"total"`
}

type GetUserTournamentDTORes struct {
	ID           string                            `json:"id"`
	Name         string                            `json:"name"`
	StartDate    *time.Time                        `json:"start_date"`
	FinishDate   *time.Time                        `json:"finish_date"`
	Points       int                               `json:"points"`
	Image        *string                           `json:"image"`
	AverageScore float32                           `json:"average_score"`
	TotalPrize   float64                           `json:"total_prize"`
	Location     *GetLocationByIDDTORes            `json:"location"`
	Organizer    *GetUserTournamentOrganizerDTORes `json:"organizer"`
}

type GetUserTournamentOrganizerDTORes struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
