package interfaces

import "context"

type UpdateQuantityGroupsInTournament interface {
	UpdateQuantityGroupsInTournament(ctx context.Context, tournamentID string, position int, add bool) error
}
