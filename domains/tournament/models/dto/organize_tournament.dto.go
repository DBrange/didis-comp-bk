package dto

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type OrganizeTournamentDTOReq struct {
	Name           string                             `json:"name" validate:"required,min=2"`
	Points         *int                               `json:"points"`
	TotalPrize     float64                            `json:"total_prize"`
	DoubleElim     *OrganizeTournamentDoubleElimDTOReq `json:"double_elimination,omitempty"`
	MaxCapacity    models.TOURNAMENT_CAPACITY         `json:"max_capacity" validate:"required,tournamentCapacity"`
	Genre          models.GENRE                       `json:"genre" validate:"required,genre"`
	Sport          models.SPORT                       `json:"sport" validate:"required,sport"`
	Location       *OrganizeTournamentLocationDTOReq   `json:"location,omitempty"`
	Surface        *models.TENNIS_SURFACE             `json:"surface" validate:"surface,omitempty"`
	CompetitorType models.COMPETITOR_TYPE             `json:"competitor_type" validate:"required,competitorType"`
	CategoryID     *string                            `json:"category_id"`
	OrganizerID    string                             `json:"organizer_id" validate:"required"`
}

type OrganizeTournamentLocationDTOReq struct {
	State   *string `json:"state"`
	Country *string `json:"country"`
	City    *string `json:"city"`
	Lat     *string `json:"lat"`
	Long    *string `json:"long"`
}
type OrganizeTournamentDoubleElimDTOReq struct {
	Points     *int    `json:"points"`
	TotalPrize float64 `json:"total_prize"`
}
