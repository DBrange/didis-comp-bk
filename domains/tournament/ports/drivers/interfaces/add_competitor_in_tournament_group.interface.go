package interfaces

import "context"

type AddCompetitorInTournamentGroup interface {
	AddCompetitorInTournamentGroup(ctx context.Context, groupID, tournamentID string, competitorID string) error
}
