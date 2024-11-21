package dto

import "github.com/DBrange/didis-comp-bk/cmd/api/models"

type GetTournamentFiltersDTORes struct {
	Surface          models.TENNIS_SURFACE  `json:"surface"`
	Sport            models.SPORT           `json:"sport"`
	CompetitorType   models.COMPETITOR_TYPE `json:"competitor_type"`
	MaxCapacity      int                    `json:"max_capacity"`
	TotalCompetitors int                    `json:"total_competitors"`
	CategoryID       *string                `json:"category_id"`
}
