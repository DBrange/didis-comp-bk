package dto

import "github.com/DBrange/didis-comp-bk/cmd/api/models"

type UpdateTournamentInfoDTOReq struct {
	Name                *string                `json:"name,omitempty"`
	Points              *int                   `json:"points,omitempty"`
	TotalPrize          *float64               `json:"total_prize,omitempty"`
	TotalCompetitors    *int                   `json:"total_competitors,omitempty"`
	AverageScore        *float32               `json:"average_score,omitempty"`
	Surface             *models.TENNIS_SURFACE `json:"surface,omitempty"`
}