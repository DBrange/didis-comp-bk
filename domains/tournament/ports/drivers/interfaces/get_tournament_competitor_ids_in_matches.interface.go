package interfaces

import (
	"context"
)

type GetTournamentCompetitorIDsInMatches interface {
	GetTournamentCompetitorIDsInMatches(ctx context.Context, tournamentID string) ([]string, error)
}
