package drivers

import (
	"context"
	"time"

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
	lastID string,
	limit int,
) (*dto.GetCompetitorsInTournamentDTORes, error) {
	return a.tournamentService.ListCompetitorsInTournament(ctx, tournamentID, lastID, limit)
}

func (a *TournamentProxyAdapter) ModifyBracketMatch(ctx context.Context, tournamentID string, competitors []*dto.UpdateCompetitorMatchDTOReq) error {
	return a.tournamentService.ModifyBracketMatch(ctx, tournamentID, competitors)
}

func (a *TournamentProxyAdapter) ModifyRoundTotalPrize(ctx context.Context, roundID string, totalPrize float64) error {
	return a.tournamentService.ModifyRoundTotalPrize(ctx, roundID, totalPrize)
}

func (a *TournamentProxyAdapter) GetRoundWithMatches(ctx context.Context, roundID, categoryID string) (*dto.GetRoundWithMatchesDTORes, error) {
	return a.tournamentService.GetRoundWithMatches(ctx, roundID, categoryID)
}

func (a *TournamentProxyAdapter) OrganizeBracket(ctx context.Context, tournamentID string, competitorMatchDTOs []*dto.UpdateCompetitorMatchDTOReq, availableCourts, averageHours int) error {
	return a.tournamentService.OrganizeBracket(ctx, tournamentID, competitorMatchDTOs, availableCourts, averageHours)
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

func (a *TournamentProxyAdapter) OrganizeTournamentGroups(ctx context.Context, tournamentID, roundID string, sport models.SPORT, orderType, top int, availableCourts, averageHours int) error {
	return a.tournamentService.OrganizeTournamentGroups(ctx, tournamentID, roundID, sport, orderType, top, availableCourts, averageHours)
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

func (a *TournamentProxyAdapter) GetUserTournaments(
	ctx context.Context,
	userID string,
	sport models.SPORT,
	limit int,
	lastID string,
) (*dto.GetUserTournamentsDTORes, error) {
	return a.tournamentService.GetUserTournaments(ctx, userID, sport, limit, lastID)
}

func (a *TournamentProxyAdapter) GetTournamentPrimaryInfo(ctx context.Context, tournamentID string) (*dto.GetTournamentPrimaryInfoDTORes, error) {
	return a.tournamentService.GetTournamentPrimaryInfo(ctx, tournamentID)
}

func (a *TournamentProxyAdapter) ListCompetitorsByNameInTournament(
	ctx context.Context,
	tournamentID string,
	name string,
	limit int,
) ([]*dto.GetCompetitorsInTournamentCompetitorDTORes, error) {
	return a.tournamentService.ListCompetitorsByNameInTournament(ctx, tournamentID, name, limit)
}

func (a *TournamentProxyAdapter) SearchCompetitorForTournament(ctx context.Context, userID string, name string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) ([]*dto.GetCompetitorFollowedDTORes, error) {
	return a.tournamentService.SearchCompetitorForTournament(ctx, userID, name, sport, competitorType)
}

func (a *TournamentProxyAdapter) RegisterDoubleCompetitorInTournament(ctx context.Context, tournamentID string, userIDs []string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) error {
	return a.tournamentService.RegisterDoubleCompetitorInTournament(ctx, tournamentID, userIDs, sport, competitorType)
}

func (a *TournamentProxyAdapter) GetTournamentFilters(ctx context.Context, tournamentID string) (*dto.GetTournamentFiltersDTORes, error) {
	return a.tournamentService.GetTournamentFilters(ctx, tournamentID)
}

func (a *TournamentProxyAdapter) GetTournamentsInOrganizer(
	ctx context.Context,
	organizerID string,
	sport models.SPORT,
	limit int,
	lastID string,
) (*dto.GetUserTournamentsDTORes, error) {
	return a.tournamentService.GetTournamentsInOrganizer(ctx, organizerID, sport, limit, lastID)
}

func (a *TournamentProxyAdapter) GetTournamentCompetitorIDs(ctx context.Context, tournamentID string) ([]string, error) {
	return a.tournamentService.GetTournamentCompetitorIDs(ctx, tournamentID)
}

func (a *TournamentProxyAdapter) RemoveCompetitorFromTournament(ctx context.Context, tournamentID, competitorID string) error {
	return a.tournamentService.RemoveCompetitorFromTournament(ctx, tournamentID, competitorID)
}

func (a *TournamentProxyAdapter) GetRoundGroups(ctx context.Context, roundID, categoryID string) (*dto.GetRoundGroupsDTORes, error) {
	return a.tournamentService.GetRoundGroups(ctx, roundID, categoryID)
}

func (a *TournamentProxyAdapter) GetTournamentAvailability(ctx context.Context, tournamentID string, day string) (*models.GetDailyAvailabilityByIDDTORes, string, error) {
	return a.tournamentService.GetTournamentAvailability(ctx, tournamentID, day)
}

func (a *TournamentProxyAdapter) GetTournamentSportsInOrganizer(ctx context.Context, organizerID string) ([]models.SPORT, error) {
	return a.tournamentService.GetTournamentSportsInOrganizer(ctx, organizerID)
}

func (a *TournamentProxyAdapter) ModifyTournamentAvailability(ctx context.Context, availabilityID string, availabilityInfoDTO *models.UpdateDailyAvailabilityDTOReq) error {
	return a.tournamentService.ModifyTournamentAvailability(ctx, availabilityID, availabilityInfoDTO)
}

func (a *TournamentProxyAdapter) GetMatchByID(ctx context.Context, matchID string) (*dto.GetMatchDTORes, error) {
	return a.tournamentService.GetMatchByID(ctx, matchID)
}

func (a *TournamentProxyAdapter) UpdateMatchDate(ctx context.Context, matchID string, date *time.Time)  error {
	return a.tournamentService.UpdateMatchDate(ctx, matchID, date)
}


func (a *TournamentProxyAdapter) GetTournamentCompetitorIDsInMatches(ctx context.Context, tournamentID string) ([]string, error) {
	return a.tournamentService.GetTournamentCompetitorIDsInMatches(ctx, tournamentID)
}