package drivers

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	option_models "github.com/DBrange/didis-comp-bk/cmd/api/models/options/tournament"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	tournament_service "github.com/DBrange/didis-comp-bk/domains/tournament/services"
)

type TournamentProxyAdapter struct {
	tournamentService *tournament_service.TournamentService
}

func NewTournamentProxyAdapter(tournamentService *tournament_service.TournamentService) *TournamentProxyAdapter {
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

func (a *TournamentProxyAdapter) AddGuestUserInTournament(ctx context.Context, tournamentID string, guestUsersDTO []*dto.CreateGuestUserDTOReq, sport models.SPORT, competitorType models.COMPETITOR_TYPE) error {
	return a.tournamentService.AddGuestUserInTournament(ctx, tournamentID, guestUsersDTO, sport, competitorType)
}

func (a *TournamentProxyAdapter) ListCompetitorsInTournament(
	ctx context.Context,
	tournamentID,
	categoryID,
	lastID string,
	limit int,
) ([]*dto.GetCompetitorsInTournamentDTORes, error) {
	return a.tournamentService.ListCompetitorsInTournament(ctx, tournamentID, categoryID, lastID, limit)
}

func (a *TournamentProxyAdapter) ModifyBracketMatch(ctx context.Context, tournamentID, userID string, competitors []*dto.UpdateCompetitorMatchDTOReq) error {
	return a.tournamentService.ModifyBracketMatch(ctx, tournamentID, userID, competitors)
}

func (a *TournamentProxyAdapter) ModifyRoundTotalPrize(ctx context.Context, roundID string, totalPrize float64) error {
	return a.tournamentService.ModifyRoundTotalPrize(ctx, roundID, totalPrize)
}

func (a *TournamentProxyAdapter) GetRoundWithMatches(ctx context.Context, roundID, categoryID string) (*dto.GetRoundWithMatchesDTORes, error) {
	return a.tournamentService.GetRoundWithMatches(ctx, roundID, categoryID)
}

func (a *TournamentProxyAdapter) OrganizeBracket(ctx context.Context, tournamentID string, competitorsDTOs []*dto.UpdateCompetitorMatchDTOReq) error {
	return a.tournamentService.OrganizeBracket(ctx, tournamentID, competitorsDTOs)
}

func (a *TournamentProxyAdapter) EndMatch(ctx context.Context, match *dto.EndMatchDTOReq) error {
	return a.tournamentService.EndMatch(ctx, match)
}

func (a *TournamentProxyAdapter) ModifyRoundPoints(ctx context.Context, roundID string, points int) error {
	return a.tournamentService.ModifyRoundPoints(ctx, roundID, points)
}

func (a *TournamentProxyAdapter) EndTournament(ctx context.Context, tournamentID string, doubleElimID string) error {
	return a.tournamentService.EndTournament(ctx, tournamentID, doubleElimID)
}

func (a *TournamentProxyAdapter) AddCompetitorInTournamentGroup(ctx context.Context, groupID, tournamentID string, competitorID string) error {
	return a.tournamentService.AddCompetitorInTournamentGroup(ctx, groupID, tournamentID, competitorID)
}

func (a *TournamentProxyAdapter) OrganizeTournamentGroups(ctx context.Context, tournamentID, roundID string, competitorDTOs []*dto.AddCompetitorsToTournamentGroupsDTOReq, sport models.SPORT) error {
	return a.tournamentService.OrganizeTournamentGroups(ctx, tournamentID, roundID, competitorDTOs, sport)
}

func (a *TournamentProxyAdapter) ModifyTournamentGroups(ctx context.Context, tournamentID, roundID string, competitorDTOs []*dto.AddCompetitorsToTournamentGroupsDTOReq, sport models.SPORT) error {
	return a.tournamentService.ModifyTournamentGroups(ctx, tournamentID, roundID, competitorDTOs, sport)
}

func (a *TournamentProxyAdapter) OrganizePots(ctx context.Context, tournamentID string, potDTOs []*dto.SetPotCompetitorDTOReq) error {
	return a.tournamentService.OrganizePots(ctx, tournamentID, potDTOs)
}

func (a *TournamentProxyAdapter) ModifyPots(ctx context.Context, tournamentID, potID, competitorID string, add bool) error {
	return a.tournamentService.ModifyPots(ctx, tournamentID, potID, competitorID, add)
}

func (a *TournamentProxyAdapter) UpdateQuantityPotsInTournament(ctx context.Context, tournamentID string, position int, add bool) error {
	return a.tournamentService.UpdateQuantityPotsInTournament(ctx, tournamentID, position, add)
}

func (a *TournamentProxyAdapter) UpdateQuantityGroupsInTournament(ctx context.Context, tournamentID string, position int, add bool) error {
	return a.tournamentService.UpdateQuantityGroupsInTournament(ctx, tournamentID, position, add)
}
