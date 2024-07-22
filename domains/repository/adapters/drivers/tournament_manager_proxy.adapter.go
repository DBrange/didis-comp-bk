package adapters

import (
	"context"

	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	tournament_opts "github.com/DBrange/didis-comp-bk/domains/repository/models/tournament"
	tournament_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/tournament/dao"
	"github.com/DBrange/didis-comp-bk/domains/repository/repository"
)

type TournamentManagerProxyAdapter struct {
	repository *repository.Repository
}

func NewTournamentManagerProxyAdapter(repository *repository.Repository) *TournamentManagerProxyAdapter {
	return &TournamentManagerProxyAdapter{
		repository: repository,
	}
}

func (a *TournamentManagerProxyAdapter) OrganizeTournament(ctx context.Context, tournamentInfoDAO *tournament_dao.CreateTournamentDAOReq, locationInfoDAO *location_dao.CreateLocationDAOReq, options *tournament_opts.OrganizeTournamentOptions, leagueID *string, organizerID string) error {
	return a.repository.OrganizeTournament(ctx, tournamentInfoDAO, locationInfoDAO, options, leagueID, organizerID)
}
