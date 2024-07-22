package ports

import (
	"context"

	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	tournament_opts "github.com/DBrange/didis-comp-bk/domains/repository/models/tournament"
	tournament_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/tournament/dao"
)

type ForManagingTournament interface {
	OrganizeTournament(ctx context.Context, tournamentInfoDAO *tournament_dao.CreateTournamentDAOReq, locationInfoDAO *location_dao.CreateLocationDAOReq, options *tournament_opts.OrganizeTournamentOptions, leagueID *string, organizerID string) error
}
