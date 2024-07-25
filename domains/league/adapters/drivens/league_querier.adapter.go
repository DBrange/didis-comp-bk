package adapters

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/league/adapters/mappers"
	"github.com/DBrange/didis-comp-bk/domains/league/models/dto"
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

func (a *LeagueQueryerAdapter) OrganizeLeague(ctx context.Context, organizerID string, leagueInfoDTO *dto.OrganizeLeagueDTOReq) error {
	leagueInfoDAO := mappers.OrganizeLeagueMapper(leagueInfoDTO)

	return a.adapter.OrganizeLeague(ctx, organizerID, leagueInfoDAO)
}

func (a *LeagueQueryerAdapter) AddTournamentInLeague(ctx context.Context, leagueID string, tournamentID string) error {
	return a.adapter.AddTournamentInLeague(ctx, leagueID, tournamentID)
}
