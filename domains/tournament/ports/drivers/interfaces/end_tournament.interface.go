package interfaces

import (
	"context"
)

type EndTournament interface {
	EndTournament(ctx context.Context, tournamentID string, doubleElimID string) error
}
