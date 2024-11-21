package dto

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type GetMatchDTORes struct {
	ID             string                                 `json:"id"`
	Date           *time.Time                             `json:"date"`
	Result         string                                 `json:"result"`
	Winner         string                                 `json:"winner"`
	Position       int                                    `json:"position"`
	PositionWinner *int                                   `json:"position_winner"`
	Sport          models.SPORT                           `json:"sport"`
	Competitors    []*GetRoundWithMatchesCompetitorDTORes `json:"competitors"`
	Tournament     *GetMatchTorurnamentDTORes             `json:"tournament"`
	Round          *GetMatchRoundDTORes                   `json:"round"`
}

type GetMatchTorurnamentDTORes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type GetMatchRoundDTORes struct {
	ID    string       `json:"id"`
	Round models.ROUND `json:"round"`
}
