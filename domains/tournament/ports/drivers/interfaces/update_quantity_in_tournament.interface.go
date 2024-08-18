package interfaces

import "context"

type UpdateQuantityPotsInTournament interface {
	UpdateQuantityPotsInTournament(ctx context.Context, tournamentID string, position int, add bool) error
}
