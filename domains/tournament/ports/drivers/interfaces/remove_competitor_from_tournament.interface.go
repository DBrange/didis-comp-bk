package interfaces

import (
	"context"
)

type RemoveCompetitorFromTournament interface {
	RemoveCompetitorFromTournament(ctx context.Context, tournamentID, competitorID string) error
}
