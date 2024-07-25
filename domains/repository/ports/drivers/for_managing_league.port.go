package ports

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/league/dao"
)

type ForManagingLeague interface {
	OrganizeLeague(ctx context.Context, organizerID string, leagueInfoDAO *dao.CreateLeagueDAOReq) error
	AddTournamentInLeague(ctx context.Context, leagueID string, tournamentID string) error
}
