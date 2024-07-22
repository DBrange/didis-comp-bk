package adapters

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/league/services"
)

type LeagueProxyAdapter struct {
	leagueService *services.LeagueService
}

func NewLeagueProxyAdapter(leagueService *services.LeagueService) *LeagueProxyAdapter {
	return &LeagueProxyAdapter{
		leagueService: leagueService,
	}
}

func (a *LeagueProxyAdapter) OrganizeLeague(ctx context.Context, leagueInfoDTO any) error {
	return a.leagueService.OrganizeLeague(ctx, leagueInfoDTO)
}
