package interfaces

import (
	"context"
)

type GetTournamentCompetitorIDs interface {
	GetTournamentCompetitorIDs(ctx context.Context, tournamentID string) ([]string, error)
}
