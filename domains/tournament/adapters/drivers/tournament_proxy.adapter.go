package drivers

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	option_models "github.com/DBrange/didis-comp-bk/cmd/api/models/options/tournament"
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

func (a *TournamentProxyAdapter) OrganizeTournament(ctx context.Context, tournamentInfoDTO *dto.OrganizeTournamentDTOReq, options *option_models.OrganizeTournamentOptions) error {
	return a.tournamentService.OrganizeTournament(ctx, tournamentInfoDTO, options)
}

func (a *TournamentProxyAdapter) AddCompetitorInTournament(ctx context.Context, tournamentResgistrationDTO *dto.CreateTournamentRegistrationDTOReq) error {
	return a.tournamentService.AddCompetitorInTournament(ctx, tournamentResgistrationDTO)
}

func (a *TournamentProxyAdapter) AddGuestUserInTournament(ctx context.Context, tournamentID string, guestUserDTO *dto.CreateGuestUserDTOReq, sport models.SPORT, competitorType models.COMPETITOR_TYPE) error {
	return a.tournamentService.AddGuestUserInTournament(ctx, tournamentID, guestUserDTO, sport, competitorType)

}
