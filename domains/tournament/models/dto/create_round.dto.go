package dto

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type CreateRoundDTOReq struct {
	TournamentID string       `json:"tournament_id"`
	Name         models.ROUND `json:"round"`
	TotalPrize   float64      `json:"total_prize"`
}
