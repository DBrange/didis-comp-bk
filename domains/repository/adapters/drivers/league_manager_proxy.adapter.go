package adapters

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/league/dao"
	"github.com/DBrange/didis-comp-bk/domains/repository/repository"
)

type LeagueManagerProxyAdapter struct {
	repository *repository.Repository
}

func NewLeagueManagerProxyAdapter(repository *repository.Repository) *LeagueManagerProxyAdapter {
	return &LeagueManagerProxyAdapter{
		repository: repository,
	}
}

func (a *LeagueManagerProxyAdapter) OrganizeLeague(ctx context.Context, organizerID string, leagueInfoDAO *dao.CreateLeagueDAOReq) error {
	return a.repository.OrganizeLeague(ctx, organizerID, leagueInfoDAO)
}

func (a *LeagueManagerProxyAdapter) AddTournamentInLeague(ctx context.Context, leagueID string, tournamentID string) error {
	return a.repository.AddTournamentInLeague(ctx, leagueID, tournamentID)
}
