package drivers

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/tournament/services"
)

type TournamentProxyAdapter struct {
	tournamentService *services.TournamentService
}

func NewTournamentProxyAdapter(tournamentService *services.TournamentService) *TournamentProxyAdapter {
	return &TournamentProxyAdapter{
		tournamentService: tournamentService,
	}
}

func (a *TournamentProxyAdapter) OrganizeTournament(ctx context.Context, tournamentInfoDTO *dto.OrganizeTournamentDTOReq) error {
	return a.tournamentService.OrganizeTournament(ctx, tournamentInfoDTO)
}
