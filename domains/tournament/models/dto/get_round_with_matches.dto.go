package dto

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type GetRoundWithMatchesDTORes struct {
	ID         *string                           `json:"id"`
	Round      models.ROUND                      `json:"round"`
	TotalPrize float64                           `json:"total_prize"`
	Points     int                               `json:"points"`
	Matches    []*GetRoundWithMatchesMatchDTORes `json:"matches"`
}

type GetRoundWithMatchesMatchDTORes struct {
	ID             *string                                `json:"id"`
	Date           *time.Time                             `json:"date"`
	Result         string                                 `json:"result"`
	Position       int                                    `json:"position"`
	PositionWinner *int                                   `json:"position_winner"`
	Competitors    []*GetRoundWithMatchesCompetitorDTORes `json:"competitors"`
}

type GetRoundWithMatchesCompetitorDTORes struct {
	ID              *string                          `json:"id"`
	CurrentPosition *int                             `json:"current_position"`
	Position        int                              `json:"position"`
	Users           []*GetRoundWithMatchesUserDTORes `json:"users"`
	GuestUsers      []*GetRoundWithMatchesUserDTORes `json:"guest_users"`
}

type GetRoundWithMatchesUserDTORes struct {
	ID        *string `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Image     string  `json:"image"`
}
