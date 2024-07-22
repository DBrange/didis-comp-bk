package adapters

import (
	"context"

	ports "github.com/DBrange/didis-comp-bk/domains/repository/ports/drivers"
)

type LeagueQueryerAdapter struct {
	adapter ports.ForManagingLeague
}

func NewLeagueQueryerAdapter(adapter ports.ForManagingLeague) *LeagueQueryerAdapter {
	return &LeagueQueryerAdapter{
		adapter: adapter,
	}
}

func (a *LeagueQueryerAdapter) OrganizeLeague(ctx context.Context, leagueInfoDTO any) error {
	// leagueInfoDAO := mappers.OrganizeLeagueMapper(leagueInfoDTO)
	leagueInfoDAO := leagueInfoDTO

	return a.adapter.OrganizeLeague(ctx, leagueInfoDAO)
}
