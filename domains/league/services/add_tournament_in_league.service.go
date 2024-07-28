package services

import (
	"context"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (d *LeagueService) AddTournamentInLeague(ctx context.Context, leagueID string, tournamentID string) error {
	if err := d.leagueQueryer.AddTournamentInLeague(ctx, leagueID, tournamentID); err != nil {
		return customerrors.HandleErrMsg(err, "league", "error when add tournament in league")
	}

	return nil
}
