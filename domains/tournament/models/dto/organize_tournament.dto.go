package dto

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type OrganizeTournamentDTOReq struct {
	Name              string                            `json:"name" validate:"required,min=2"`
	Points            *int                              `json:"points"`
	TotalPrize        float64                           `json:"total_prize"`
	TotalCompetitors  int                               `json:"total_competitors" validate:"required"`
	MaxCapacity       int                               `json:"max_capacity" validate:"required"`
	Genre             models.GENRE                      `json:"genre" validate:"required,genre"`
	Sport             models.SPORT                      `json:"sport" validate:"required,sport"`
	Location          *OrganizeTournamentLocationDTOReq `json:"location"`
	Surface           *models.TENNIS_SURFACE            `json:"surface"`
	DoubleElimination bool                              `json:"double_elimination" validate:"bool"`
	Pots              bool                              `json:"pots" validate:"bool"`
	Groups            bool                              `json:"groups" validate:"bool"`
	LeagueID          *string                           `json:"league_id"`
	OrganizerID       string                            `json:"organizer_id" validate:"required"`
}

type OrganizeTournamentLocationDTOReq struct {
	State   *string `json:"state"`
	Country *string `json:"country"`
	City    *string `json:"city"`
	Lat     *string `json:"lat"`
	Long    *string `json:"long"`
}
