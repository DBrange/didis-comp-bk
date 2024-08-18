package dto

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type CreateMatchDTOReq struct {
	Sport        models.SPORT `json:"sport"`
	RoundID      string       `json:"round_id"`
	Result       string       `json:"result"`
	Winner       string       `json:"winner"`
	TournamentID string       `json:"tournament_id"`
	Position     int          `json:"position"`
}
