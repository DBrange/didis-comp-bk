package dto

import "github.com/DBrange/didis-comp-bk/cmd/api/models"

type GetRoundWithMatchesDTORes struct {
	ID         *string                           `json:"id"`
	Round      models.ROUND                      `json:"round"`
	TotalPrize float64                           `json:"total_prize"`
	Matches    []*GetRoundWithMatchesMatchDTORes `json:"matches"`
}

type GetRoundWithMatchesMatchDTORes struct {
	ID             *string                                `json:"id"`
	Result         string                                 `json:"result"`
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
