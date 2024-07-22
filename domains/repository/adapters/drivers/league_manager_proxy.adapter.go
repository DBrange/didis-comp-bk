package adapters

import (
	"context"

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

func (a *LeagueManagerProxyAdapter) OrganizeLeague(ctx context.Context, leagueInfoDAO any) error {
	return a.repository.OrganizeLeague(ctx, leagueInfoDAO)
}
