package dto

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
)

type GetTournamentPrimaryInfoDTORes struct {
	ID               string                                  `json:"id"`
	Name             string                                  `json:"name"`
	FinishDate       *time.Time                              `json:"finish_date"`
	StartDate        *time.Time                              `json:"start_date"`
	Points           int                                     `json:"points"`
	Image            *string                                 `json:"image"`
	TotalPrize       float64                                 `json:"total_prize"`
	TotalCompetitors int                                     `json:"total_competitors"`
	MaxCapacity      models.TOURNAMENT_CAPACITY              `json:"max_capacity"`
	AverageScore     float32                                 `json:"average_score"`
	Genre            models.GENRE                            `json:"genre"`
	Sport            models.SPORT                            `json:"sport"`
	CompetitorType   models.SPORT                            `json:"competitor_type"`
	Surface          models.TENNIS_SURFACE                   `json:"surface"`
	Rounds           []*GetTournamentPrimaryInforRoundDTORes `json:"rounds"`
	Location         *dto.GetLocationByIDDTORes              `json:"location"`
	Organizer        *GetUserTournamentOrganizerDTORes       `json:"organizer"`
	Category         *GetTournamentPrimaryInfoCategoryDTORes `json:"category"`
}

type GetTournamentPrimaryInforRoundDTORes struct {
	ID    string       `json:"id"`
	Round models.ROUND `json:"round"`
}

type GetTournamentPrimaryInfoCategoryDTORes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
