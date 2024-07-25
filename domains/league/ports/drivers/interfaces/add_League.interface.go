package interfaces

import "context"

type AddTournamentInLeague interface {
	AddTournamentInLeague(ctx context.Context, leagueID string, tournamentID string) error
}
