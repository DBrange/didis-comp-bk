package interfaces

import "context"

type AddTournamentInCategory interface {
	AddTournamentInCategory(ctx context.Context, categoryID string, tournamentID string) error
}
