package dto

import "github.com/DBrange/didis-comp-bk/cmd/api/models"

type EndMatchDTOReq struct {
	TournamentID       string       `json:"tournament_id" validate:"required"`
	MatchID            string       `json:"match_id" validate:"required"`
	WinnerCompetitorID string       `json:"winner_competitor_id" validate:"required"`
	LosserCompetitorID string       `json:"losser_competitor_id" validate:"required"`
	Result             string       `json:"result" validate:"required"`
	Sport              models.SPORT `json:"sport" validate:"required,sport"`
	Round              models.ROUND `json:"round" validate:"round"`
	RoundID            string       `json:"round_id" validate:"required"`
	DoubleElimID       string       `json:"double_elimination_id" validate:""`
}
