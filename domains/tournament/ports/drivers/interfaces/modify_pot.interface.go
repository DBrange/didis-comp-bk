package interfaces

import "context"

type ModifyPots interface {
	ModifyPots(ctx context.Context, tournamentID, potID, competitorID string, add bool) error
}
