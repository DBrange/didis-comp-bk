package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

type GetTournamentPrimaryInfo interface {
	GetTournamentPrimaryInfo(ctx context.Context, tournamentID string) (*dto.GetTournamentPrimaryInfoDTORes, error)
}
