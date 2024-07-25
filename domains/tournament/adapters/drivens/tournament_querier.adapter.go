package drivens

import (
	"context"

	ports "github.com/DBrange/didis-comp-bk/domains/repository/ports/drivers"
	"github.com/DBrange/didis-comp-bk/domains/tournament/adapters/mappers"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

type TournamentQueryerAdapter struct {
	adapter ports.ForManagingTournament
}

func NewTournamentQueryerAdapter(adapter ports.ForManagingTournament) *TournamentQueryerAdapter {
	return &TournamentQueryerAdapter{
		adapter: adapter,
	}
}

func (a *TournamentQueryerAdapter) OrganizeTournament(ctx context.Context, tournamentInfoDTO *dto.OrganizeTournamentDTOReq) error {
	tournamentInfoDAO, locationInfoDAO, organizeTournamentOptions, leagueID, organizerID := mappers.OrganizeTournamentMapper(tournamentInfoDTO)

	return a.adapter.OrganizeTournament(ctx, tournamentInfoDAO, locationInfoDAO, organizeTournamentOptions, leagueID, organizerID)
}
