package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

type GetTournamentFilters interface {
	GetTournamentFilters(ctx context.Context, tournamentID string) (*dto.GetTournamentFiltersDTORes, error)
}
