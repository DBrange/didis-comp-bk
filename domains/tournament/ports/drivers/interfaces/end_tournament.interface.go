package interfaces

import (
	"context"
)

type EndTournament interface {
	EndTournament(ctx context.Context, tournamentID, competitorID string) error
}
