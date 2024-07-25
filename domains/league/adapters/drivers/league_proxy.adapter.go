package adapters

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/league/models/dto"
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

func (a *LeagueProxyAdapter) OrganizeLeague(ctx context.Context, organizerID string, leagueInfoDTO *dto.OrganizeLeagueDTOReq) error {
	return a.leagueService.OrganizeLeague(ctx, organizerID, leagueInfoDTO)
}

func (a *LeagueProxyAdapter) AddTournamentInLeague(ctx context.Context, leagueID string, tournamentID string) error {
	return a.leagueService.AddTournamentInLeague(ctx,leagueID, tournamentID)
}
