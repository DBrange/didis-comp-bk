package dto

import "github.com/DBrange/didis-comp-bk/cmd/api/models"

type CreateTournamentDTOReq struct {
	Name             string                     `json:"name"`
	Points           *int                       `json:"points"`
	TotalPrize       float64                    `json:"total_prize"`
	TotalCompetitors int                        `json:"total_competitors"`
	MaxCapacity      models.TOURNAMENT_CAPACITY `json:"max_capacity"`
	Genre            models.GENRE               `json:"genre"`
	Sport            models.SPORT               `json:"sport"`
	Surface          models.TENNIS_SURFACE      `json:"surface"`
	CompetitorType   models.COMPETITOR_TYPE     `json:"competitor_type"`
}
