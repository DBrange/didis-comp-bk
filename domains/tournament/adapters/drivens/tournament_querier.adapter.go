package drivens

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	option_models "github.com/DBrange/didis-comp-bk/cmd/api/models/options/tournament"
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	ports "github.com/DBrange/didis-comp-bk/domains/repository/ports/drivers"
	"github.com/DBrange/didis-comp-bk/domains/tournament/adapters/mappers"
	tournament_dto "github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TournamentQueryerAdapter struct {
	adapter ports.ForManagingTournament
}

func NewTournamentQueryerAdapter(adapter ports.ForManagingTournament) *TournamentQueryerAdapter {
	return &TournamentQueryerAdapter{
		adapter: adapter,
	}
}

func (a *TournamentQueryerAdapter) CreateLocation(ctx context.Context, locationDTO *tournament_dto.CreateLocationDTOReq) (string, error) {
	locationDAO := mappers.CreateLocationDTOtoDAO(locationDTO)

	return a.adapter.CreateLocation(ctx, locationDAO)
}

func (a *TournamentQueryerAdapter) VerifyOrganizerExists(ctx context.Context, organizerID string) error {
	return a.adapter.VerifyOrganizerExists(ctx, organizerID)
}

func (a *TournamentQueryerAdapter) CreateTournament(
	ctx context.Context,
	tournamentDTO *tournament_dto.CreateTournamentDTOReq,
	locationID string,
	options *option_models.OrganizeTournamentOptions,
	categoryID *string,
	organizerID string,
) (string, error) {
	tournamentDAO := mappers.CreateTournamentDTOtoDAO(tournamentDTO)

	return a.adapter.CreateTournament(ctx, tournamentDAO, locationID, options, categoryID, organizerID)
}

func (a *TournamentQueryerAdapter) VerifyCategoryExists(ctx context.Context, categoryID string) error {
	return a.adapter.VerifyCategoryExists(ctx, categoryID)
}

func (a *TournamentQueryerAdapter) AddTournamentInCategory(ctx context.Context, categoryID string, tournamentID string) error {
	return a.adapter.AddTournamentInCategory(ctx, categoryID, tournamentID)
}

func (a *TournamentQueryerAdapter) WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) error) error {
	return a.adapter.WithTransaction(ctx, fn)
}

func (a *TournamentQueryerAdapter) CreateTournamentGroup(ctx context.Context, tournamentID string, position int) (string, error) {
	tournamentIOD, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return "", nil
	}

	return a.adapter.CreateTournamentGroup(ctx, tournamentIOD, position)
}

func (a *TournamentQueryerAdapter) CreatePot(ctx context.Context, tournamentID string, position int) (string, error) {
	tournamentIOD, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return "", nil
	}

	return a.adapter.CreatePot(ctx, tournamentIOD, position)
}

func (a *TournamentQueryerAdapter) CreateDoubleEliminationEmpty(ctx context.Context) (string, error) {
	return a.adapter.CreateDoubleEliminationEmpty(ctx)
}

func (a *TournamentQueryerAdapter) TournamentGroupColl() *mongo.Collection {
	return a.adapter.TournamentGroupColl()
}

func (a *TournamentQueryerAdapter) PotColl() *mongo.Collection {
	return a.adapter.PotColl()
}

func (a *TournamentQueryerAdapter) DoubleEliminationColl() *mongo.Collection {
	return a.adapter.DoubleEliminationColl()
}

func (a *TournamentQueryerAdapter) DeleteByID(ctx context.Context, mc *mongo.Collection, ID string, name string) error {
	return a.adapter.DeleteByID(ctx, mc, ID, name)
}

func (a *TournamentQueryerAdapter) UpdateTournamentRelations(
	ctx context.Context,
	tournamentID string,
	tournamentDTO *tournament_dto.UpdateTournamentOptionsDTOReq,
	add bool,
) error {
	tournamentDAO, err := mappers.UpdateTournamentOptionsDTOtoDAO(tournamentDTO, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.adapter.UpdateTournamentRelations(ctx, tournamentOID, tournamentDAO, add)
}

func (a *TournamentQueryerAdapter) ConvertToObjectID(ID string) (*primitive.ObjectID, error) {
	return a.adapter.ConvertToObjectID(ID)
}

func (a *TournamentQueryerAdapter) CreateTournamentRegistration(ctx context.Context, tournamentRegistrationDTO *tournament_dto.CreateTournamentRegistrationDTOReq) error {
	tournamentRegistrationDAO, err := mappers.CreateTournamentRegistrationDTOtoDAO(tournamentRegistrationDTO, a.ConvertToObjectID)
	if err != nil {
		return nil
	}

	return a.adapter.CreateTournamentRegistration(ctx, tournamentRegistrationDAO)
}

func (a *TournamentQueryerAdapter) CreateGuestUser(ctx context.Context, guestUserDTO *tournament_dto.CreateGuestUserDTOReq) (string, error) {
	guestUserDAO := mappers.CreateGuestUserDTOtoDAO(guestUserDTO)

	return a.adapter.CreateGuestUser(ctx, guestUserDAO)
}

func (a *TournamentQueryerAdapter) CreateCompetitor(ctx context.Context, sport models.SPORT, competitorType models.COMPETITOR_TYPE, ID string) (string, error) {
	OID, err := a.ConvertToObjectID(ID)
	if err != nil {
		return "", err
	}
	return a.adapter.CreateCompetitor(ctx, sport, competitorType, OID)
}

func (a *TournamentQueryerAdapter) CreateGuestCompetitor(ctx context.Context, guestCompetitorDTO *tournament_dto.CreateGuestCompetitorDTOReq) (string, error) {
	guestCompetitorDAO, err := mappers.CreateGuestCompetitorDTOtoDAO(guestCompetitorDTO, a.ConvertToObjectID)
	if err != nil {
		return "", err
	}

	return a.adapter.CreateGuestCompetitor(ctx, guestCompetitorDAO)
}

func (a *TournamentQueryerAdapter) CreateMatch(ctx context.Context, matchDTO *tournament_dto.CreateMatchDTOReq) (string, error) {
	matchDAO, err := mappers.CreateMatchDTOtoDAO(matchDTO, a.ConvertToObjectID)
	if err != nil {
		return "", err
	}

	return a.adapter.CreateMatch(ctx, matchDAO)
}

func (a *TournamentQueryerAdapter) CreateRound(ctx context.Context, roundDTO *tournament_dto.CreateRoundDTOReq) (string, error) {
	roundDAO, err := mappers.CreateRoundDTOtoDAO(roundDTO, a.ConvertToObjectID)
	if err != nil {
		return "", err
	}

	return a.adapter.CreateRound(ctx, roundDAO)
}

func (a *TournamentQueryerAdapter) CreateDoubleElimination(ctx context.Context, doubleEliminationDTO *tournament_dto.CreateDoubleEliminationDTOReq) (string, error) {
	doubleEliminationDAO, err := mappers.CreateDoubleEliminationDTOtoDAO(doubleEliminationDTO, a.ConvertToObjectID)
	if err != nil {
		return "", err
	}

	return a.adapter.CreateDoubleElimination(ctx, doubleEliminationDAO)
}

func (a *TournamentQueryerAdapter) CreateSingle(ctx context.Context, singleDTO *tournament_dto.CreateSingleDTOReq) (string, error) {
	singleDAO := mappers.CreateSingleDTOtoDAO(singleDTO)

	return a.adapter.CreateSingle(ctx, singleDAO)
}

func (a *TournamentQueryerAdapter) CreateDouble(ctx context.Context, doubleDTO *tournament_dto.CreateDoubleDTOReq) (string, error) {
	doubleDAO := mappers.CreateDoubleDTOtoDAO(doubleDTO)

	return a.adapter.CreateDouble(ctx, doubleDAO)
}

func (a *TournamentQueryerAdapter) CreateTeam(ctx context.Context, teamDTO *tournament_dto.CreateTeamDTOReq) (string, error) {
	teamDAO, err := mappers.CreateTeamDTOtoDAO(teamDTO, a.ConvertToObjectID)
	if err != nil {
		return "", nil
	}

	return a.adapter.CreateTeam(ctx, teamDAO)
}

func (a *TournamentQueryerAdapter) GetCompetitorsInTournament(
	ctx context.Context,
	tournamentID, categoryID,
	lastID string, limit int,
) ([]*tournament_dto.GetCompetitorsInTournamentDTORes, error) {
	competitorsDAO, err := a.adapter.ListCompetitorsInTournament(ctx, tournamentID, categoryID, lastID, limit)
	if err != nil {
		return nil, err
	}

	competitorsDTO := mappers.GetCompetitorsInTournamentDAOtoDTO(competitorsDAO)

	return competitorsDTO, nil
}

func (a *TournamentQueryerAdapter) VerifyCompetitorExists(ctx context.Context, competitorID string) error {
	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}
	return a.adapter.VerifyCompetitorExists(ctx, competitorOID)
}

func (a *TournamentQueryerAdapter) VerifyTournamentsExists(ctx context.Context, tournamentID string) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.adapter.VerifyTournamentsExists(ctx, tournamentOID)
}

func (a *TournamentQueryerAdapter) CreateCompetitorStats(ctx context.Context, competitorID string) error {
	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}

	return a.adapter.CreateCompetitorStats(ctx, competitorOID)
}

func (a *TournamentQueryerAdapter) UpdateCompetitorMatch(ctx context.Context, matchID string, competitorMatchDTO *tournament_dto.UpdateCompetitorMatchDTOReq) error {
	competitorMatchDAO, matchOID, err := mappers.UpdateCompetitorMatchDTOtoDAO(competitorMatchDTO, matchID, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.adapter.UpdateCompetitorMatch(ctx, matchOID, competitorMatchDAO)
}

func (a *TournamentQueryerAdapter) VerifyMatchExists(ctx context.Context, matchID string) error {
	matchOID, err := a.ConvertToObjectID(matchID)
	if err != nil {
		return err
	}

	return a.adapter.VerifyMatchExists(ctx, matchOID)
}

func (a *TournamentQueryerAdapter) CreateCompetitorMatch(ctx context.Context, competitorMatchDTO *tournament_dto.CreateCompetitorMatchDTOReq) error {
	competitorMatchDAO, err := mappers.CreateCompetitorMatchDTOtoDAO(competitorMatchDTO, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.adapter.CreateCompetitorMatch(ctx, competitorMatchDAO)
}

func (a *TournamentQueryerAdapter) UpdateRoundTotalPrize(ctx context.Context, roundID string, totalPrize float64) error {
	roundOID, err := a.ConvertToObjectID(roundID)
	if err != nil {
		return err
	}

	return a.adapter.UpdateRoundTotalPrize(ctx, roundOID, totalPrize)
}

func (a *TournamentQueryerAdapter) VerifyRoundExists(ctx context.Context, roundID string) error {
	roundOID, err := a.ConvertToObjectID(roundID)
	if err != nil {
		return err
	}

	return a.adapter.VerifyRoundExists(ctx, roundOID)
}

func (a *TournamentQueryerAdapter) GetRoundWithMatches(ctx context.Context, roundID, categoryID string) (*tournament_dto.GetRoundWithMatchesDTORes, error) {
	roundDAO, err := a.adapter.GetRoundWithMatches(ctx, roundID, categoryID)
	if err != nil {
		return nil, err
	}

	roundDTO := mappers.MapRoundWithMatchesDAOToDTO(roundDAO)

	return roundDTO, nil
}

func (a *TournamentQueryerAdapter) GetPositionsBracketMatch(ctx context.Context, roundID string, organizeType models.ORGANIZE_TYPE, organizeBracket models.ORGANIZE_BRACKET) ([]*tournament_dto.GetPositionsBracketMatchDTORes, error) {
	positionsDAO, err := a.adapter.GetPositionsBracketMatch(ctx, roundID, organizeType, organizeBracket)
	if err != nil {
		return nil, err
	}

	positiondDTO := mappers.GetPositionsBracketMatchDAOtoDTO(positionsDAO)

	return positiondDTO, nil
}

func (a *TournamentQueryerAdapter) UpdateMultipleCompetitorMatches(ctx context.Context, competitorsDTOs []*tournament_dto.UpdateCompetitorMatchDTOReq) error {
	competitorsDAOs, err := mappers.UpdateMultipleCompetitorMatchesDTOtoDAO(competitorsDTOs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.adapter.UpdateMultipleCompetitorMatches(ctx, competitorsDAOs)
}

func (a *TournamentQueryerAdapter) VerifyMatchesExist(ctx context.Context, matchOIDs []string) error {
	competitorsDAOs, err := utils.ConvertToObjectIDs(&matchOIDs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.adapter.VerifyMatchesExist(ctx, *competitorsDAOs)
}

func (a *TournamentQueryerAdapter) VerifyMultipleCompetitorsExists(ctx context.Context, competitorOIDs []string) error {
	competitorsDAOs, err := utils.ConvertToObjectIDs(&competitorOIDs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.adapter.VerifyMultipleCompetitorsExists(ctx, *competitorsDAOs)
}

func (a *TournamentQueryerAdapter) VerifyMatchesInRoundExits(ctx context.Context, roundID string) (bool, error) {
	roundOID, err := a.ConvertToObjectID(roundID)
	if err != nil {
		return false, err
	}

	return a.adapter.VerifyMatchesInRoundExits(ctx, roundOID)
}

func (a *TournamentQueryerAdapter) SetWinnerInMatch(ctx context.Context, matchID, competitorID, result string) error {
	matchOID, err := a.ConvertToObjectID(matchID)
	if err != nil {
		return err
	}

	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}

	return a.adapter.SetWinnerInMatch(ctx, matchOID, competitorOID, result)
}

func (a *TournamentQueryerAdapter) FindMatchID(ctx context.Context, position int, roundID string) (string, error) {
	roundOID, err := a.ConvertToObjectID(roundID)
	if err != nil {
		return "", err
	}

	return a.adapter.FindMatchID(ctx, position, roundOID)
}

func (a *TournamentQueryerAdapter) AddMatchInTournament(ctx context.Context, tournamentID, matchID string) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	matchOID, err := a.ConvertToObjectID(matchID)
	if err != nil {
		return err
	}

	return a.adapter.AddMatchInTournament(ctx, tournamentOID, matchOID)
}

func (a *TournamentQueryerAdapter) AddMatchInCompetitorStats(ctx context.Context, competitorID, matchID string) error {
	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}

	matchOID, err := a.ConvertToObjectID(matchID)
	if err != nil {
		return err
	}

	return a.adapter.AddMatchInCompetitorStats(ctx, competitorOID, matchOID)
}

func (a *TournamentQueryerAdapter) UpdateCompetitorStats(ctx context.Context, competitorID string, winner bool) error {
	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}

	return a.adapter.UpdateCompetitorStats(ctx, competitorOID, winner)
}

func (a *TournamentQueryerAdapter) IncrementTotalCompetitorsInTournament(ctx context.Context, tournamentID string) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.adapter.IncrementTotalCompetitorsInTournament(ctx, tournamentOID)
}

func (a *TournamentQueryerAdapter) VerifyTournamentsCapacity(ctx context.Context, tournamentID string) (bool, error) {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return false, err
	}

	return a.adapter.VerifyTournamentsCapacity(ctx, tournamentOID)
}

func (a *TournamentQueryerAdapter) UpdateRoundPoints(ctx context.Context, roundID string, points int) error {
	roundOID, err := a.ConvertToObjectID(roundID)
	if err != nil {
		return err
	}

	return a.adapter.UpdateRoundPoints(ctx, roundOID, points)
}

func (a *TournamentQueryerAdapter) UpdateTournamentInfo(ctx context.Context, tournamentID string, tournamentDTO *tournament_dto.UpdateTournamentInfoDTOReq) error {
	tournamentDAO, tournamentOID, err := mappers.UpdateTournamentInfoDTOtoDAO(tournamentDTO, tournamentID, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.adapter.UpdateTournamentInfo(ctx, tournamentOID, tournamentDAO)
}

func (a *TournamentQueryerAdapter) AddTournamentWonInCompetitorStats(ctx context.Context, competitorID, tournamentID string) error {
	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}

	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.adapter.AddTournamentWonInCompetitorStats(ctx, competitorOID, tournamentOID)
}

func (a *TournamentQueryerAdapter) GetRoundsWithCompetitors(ctx context.Context, tournamentID string) ([]*tournament_dto.GetRoundWithCompetitorsDTORes, error) {
	roundDAO, err := a.adapter.GetRoundsWithCompetitors(ctx, tournamentID)
	if err != nil {
		return nil, err
	}

	roundDTO := mappers.GetRoundsWithCompetitorsDAOtoDTO(roundDAO)

	return roundDTO, nil
}

func (a *TournamentQueryerAdapter) GetTournamentCompetitorIDs(ctx context.Context, tournamentID string) ([]string, error) {
	return a.adapter.GetTournamentCompetitorIDs(ctx, tournamentID)
}

func (a *TournamentQueryerAdapter) GetCompetitorsOutCategory(ctx context.Context, categoryID string, competitorIDs []string) ([]string, error) {
	return a.adapter.GetCompetitorsOutCategory(ctx, categoryID, competitorIDs)
}

func (a *TournamentQueryerAdapter) CreateCategoryRegistration(ctx context.Context, categoryRegistrationDTO *tournament_dto.CreateCategoryRegistrationDTOReq) error {
	categoryRegistrationDAO, err := mappers.CreateCategoryRegistrationDTOtoDAO(categoryRegistrationDTO, a.ConvertToObjectID)
	if err != nil {
		return nil
	}

	return a.adapter.CreateCategoryRegistration(ctx, categoryRegistrationDAO)
}

func (a *TournamentQueryerAdapter) IncrementTotalParticipants(ctx context.Context, categoryID string) error {
	categoryOID, err := a.ConvertToObjectID(categoryID)
	if err != nil {
		return err
	}

	return a.adapter.IncrementTotalParticipants(ctx, categoryOID)
}

func (a *TournamentQueryerAdapter) GetCategoryRegistrationSortedByPoints(ctx context.Context, categoryID string) ([]*tournament_dto.GetCategoryRegistrationSortedByPointsDTORes, error) {
	categoryRegistrationSortedDAO, err := a.adapter.GetCategoryRegistrationSortedByPoints(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	categoryRegistrationSortedDTO := mappers.GetCategoryRegistrationSortedByPointsDAOtoDTO(categoryRegistrationSortedDAO)

	return categoryRegistrationSortedDTO, nil
}

func (a *TournamentQueryerAdapter) UpdateCategoryRegistrationCurrentPosition(ctx context.Context, categoryID string, categoryRegistrationDTO []*tournament_dto.GetCategoryRegistrationSortedByPointsDTORes) error {
	categoryRegistrationDAO, categoryOID, err := mappers.UpdateCategoryRegistrationCurrentPositionDTOtoDAO(categoryRegistrationDTO, categoryID, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.adapter.UpdateCategoryRegistrationCurrentPosition(ctx, categoryOID, categoryRegistrationDAO)
}

func (a *TournamentQueryerAdapter) VerifyCompetitorsMatch(ctx context.Context, matchID, competitorID string) error {
	matchOID, err := a.ConvertToObjectID(matchID)
	if err != nil {
		return err
	}

	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}

	return a.adapter.VerifyCompetitorsMatch(ctx, matchOID, competitorOID)
}

func (a *TournamentQueryerAdapter) UpdateTournamentFinishDate(ctx context.Context, tournamentID string) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.adapter.UpdateTournamentFinishDate(ctx, tournamentOID)
}

func (a *TournamentQueryerAdapter) VerifyMatchPosition(ctx context.Context, matchID string, position int) error {
	matchOID, err := a.ConvertToObjectID(matchID)
	if err != nil {
		return err
	}

	return a.adapter.VerifyMatchPosition(ctx, matchOID, position)
}

func (a *TournamentQueryerAdapter) GetRoundQuantityMatches(ctx context.Context, roundID string) (int, error) {
	return a.adapter.GetRoundQuantityMatches(ctx, roundID)
}

func (a *TournamentQueryerAdapter) GetMatchPosition(ctx context.Context, matchID string) (int, error) {
	return a.adapter.GetMatchPosition(ctx, matchID)
}

func (a *TournamentQueryerAdapter) GetRoundID(ctx context.Context, tournamentID string, round models.ROUND) (string, error) {
	return a.adapter.GetRoundID(ctx, tournamentID, round)
}

func (a *TournamentQueryerAdapter) AddPointsInMultipleCategoryRegistration(ctx context.Context, categoryID string, competitorIDs []string, points int) error {
	categoryOID, err := a.ConvertToObjectID(categoryID)
	if err != nil {
		return err
	}

	competitorOIDs, err := utils.ConvertToObjectIDs(&competitorIDs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.adapter.AddPointsInMultipleCategoryRegistration(ctx, categoryOID, *competitorOIDs, points)
}

func (a *TournamentQueryerAdapter) AddPrizeInMultipleCompetitorStats(ctx context.Context, competitorIDs []string, prize float64) error {
	competitorOIDs, err := utils.ConvertToObjectIDs(&competitorIDs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.adapter.AddPrizeInMultipleCompetitorStats(ctx, *competitorOIDs, prize)
}

func (a *TournamentQueryerAdapter) GetTournamentInfoToFinaliseIt(ctx context.Context, tournamentID string) (*tournament_dto.GetTournamentInfoToFinaliseItDTORes, error) {
	tournamentInfoDAO, err := a.adapter.GetTournamentInfoToFinaliseIt(ctx, tournamentID)
	if err != nil {
		return nil, err
	}

	tournamentInfoDTO := mappers.GetTournamentInfoToFinaliseItDAOtoDTO(tournamentInfoDAO)

	return tournamentInfoDTO, nil
}

func (a *TournamentQueryerAdapter) VerifyTournamentsAlreadyFinished(ctx context.Context, tournamentID string) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.adapter.VerifyTournamentsAlreadyFinished(ctx, tournamentOID)
}

func (a *TournamentQueryerAdapter) VerifyMatchAndRoundCoincidence(ctx context.Context, matchID, roundID string, round models.ROUND) error {
	matchOID, err := a.ConvertToObjectID(matchID)
	if err != nil {
		return err
	}

	roundOID, err := a.ConvertToObjectID(roundID)
	if err != nil {
		return err
	}

	return a.adapter.VerifyMatchAndRoundCoincidence(ctx, matchOID, roundOID, round)
}

func (a *TournamentQueryerAdapter) VerifyMultipleCompetitorsExistsInTournament(ctx context.Context, tournamentID string, competitorIDs []string) error {
	competitorOIDs, err := utils.ConvertToObjectIDs(&competitorIDs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.adapter.VerifyMultipleCompetitorsExistsInTournament(ctx, tournamentOID, *competitorOIDs)
}

func (a *TournamentQueryerAdapter) VerifyCompetitorExistsInTournament(ctx context.Context, tournamentID string, competitorID string) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}

	return a.adapter.VerifyCompetitorExistsInTournament(ctx, tournamentOID, competitorOID)
}

func (a *TournamentQueryerAdapter) AddCompetitorInGroup(ctx context.Context, groupID, competitorID string) error {
	groupOID, err := a.ConvertToObjectID(groupID)
	if err != nil {
		return err
	}

	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}

	return a.adapter.AddCompetitorInGroup(ctx, groupOID, competitorOID)
}

func (a *TournamentQueryerAdapter) AddCompetitorsToTournamentGroups(ctx context.Context, tournamentID string, groupDTOs []*tournament_dto.AddCompetitorsToTournamentGroupsDTOReq) error {
	groupDAOs, tournamentOID, err := mappers.AddCompetitorsToTournamentGroupsDTOtoDAO(groupDTOs, tournamentID, a.ConvertToObjectID)
	if err != nil {
		return nil
	}

	return a.adapter.AddCompetitorsToTournamentGroups(ctx, tournamentOID, groupDAOs)
}

func (a *TournamentQueryerAdapter) AddMatchInTournamentGroup(ctx context.Context, groupID, tournamentID, matchID string) error {
	groupOID, tournamentOID, matchOID, err := mappers.AddMatchInTournamentGroupDTOtoDAO(groupID, tournamentID, matchID, a.ConvertToObjectID)
	if err != nil {
		return nil
	}

	return a.adapter.AddMatchInTournamentGroup(ctx, groupOID, tournamentOID, matchOID)
}

func (a *TournamentQueryerAdapter) VerifyMultipleGroupsInTournament(ctx context.Context, tournamentID string, groupIDs []string) error {
	groupOIDs, err := utils.ConvertToObjectIDs(&groupIDs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.adapter.VerifyMultipleGroupsInTournament(ctx, tournamentOID, *groupOIDs)
}

func (a *TournamentQueryerAdapter) VerifyRoundInTournament(ctx context.Context, roundID, tournamentID string) error {
	roundOID, err := a.ConvertToObjectID(roundID)
	if err != nil {
		return err
	}

	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.adapter.VerifyRoundInTournament(ctx, roundOID, tournamentOID)
}

func (a *TournamentQueryerAdapter) AddMultipleMatchesInTournamentGroup(ctx context.Context, groupID, tournamentID string, matchIDs []string) error {
	groupOID, tournamentOID, matchOIDs, err := mappers.AddMultipleMatchesInTournamentGroupDTOtoDAO(groupID, tournamentID, matchIDs, a.ConvertToObjectID)
	if err != nil {
		return nil
	}

	return a.adapter.AddMultipleMatchesInTournamentGroup(ctx, groupOID, tournamentOID, matchOIDs)
}

func (a *TournamentQueryerAdapter) AddMultipleMatchesInTournament(ctx context.Context, tournamentID string, matchIDs []string) error {
	tournamentOID, matchOIDs, err := mappers.AddMultipleMatchesInTournamentDTOtoDAO(tournamentID, matchIDs, a.ConvertToObjectID)
	if err != nil {
		return nil
	}

	return a.adapter.AddMultipleMatchesInTournament(ctx, tournamentOID, matchOIDs)
}

func (a *TournamentQueryerAdapter) VerifyTournamentGroupInTournament(ctx context.Context, tournamentID string, groupIDs []string) error {
	groupOIDs, err := utils.ConvertToObjectIDs(&groupIDs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.adapter.VerifyTournamentGroupInTournament(ctx, tournamentOID, *groupOIDs)
}

func (a *TournamentQueryerAdapter) GetTournamentGroupMatches(ctx context.Context, groupID string) ([]string, []string, error) {
	return a.adapter.GetTournamentGroupMatches(ctx, groupID)
}

func (a *TournamentQueryerAdapter) RemoveMultipleTournamentMatches(ctx context.Context, tournamentID string, matchesToRemoveIDs []string) error {
	matchesToRemoveOIDs, err := utils.ConvertToObjectIDs(&matchesToRemoveIDs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.adapter.RemoveMultipleTournamentMatches(ctx, tournamentOID, *matchesToRemoveOIDs)
}

func (a *TournamentQueryerAdapter) RemoveMultipleCompetitorStatsMatches(ctx context.Context, competitorIDs, matchesToRemoveIDs []string) error {
	competitorOIDs, err := utils.ConvertToObjectIDs(&competitorIDs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	matchesToRemoveOIDs, err := utils.ConvertToObjectIDs(&matchesToRemoveIDs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.adapter.RemoveMultipleCompetitorStatsMatches(ctx, *competitorOIDs, *matchesToRemoveOIDs)
}

func (a *TournamentQueryerAdapter) DeleteMultipleMatches(ctx context.Context, matchesToRemoveIDs []string) error {
	matchesToRemoveOIDs, err := utils.ConvertToObjectIDs(&matchesToRemoveIDs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.adapter.DeleteMultipleMatches(ctx, *matchesToRemoveOIDs)
}

func (a *TournamentQueryerAdapter) DeleteMultipleCompetitorMatches(ctx context.Context, matchesToRemoveIDs []string) error {
	matchesToRemoveOIDs, err := utils.ConvertToObjectIDs(&matchesToRemoveIDs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.adapter.DeleteMultipleCompetitorMatches(ctx, *matchesToRemoveOIDs)
}

func (a *TournamentQueryerAdapter) VerifyTournamentPot(ctx context.Context, tournamentID, potID string) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	potOID, err := a.ConvertToObjectID(potID)
	if err != nil {
		return err
	}

	return a.adapter.VerifyTournamentPot(ctx, tournamentOID, potOID)
}

func (a *TournamentQueryerAdapter) AddCompetitorInPot(ctx context.Context, potID, competitorID string) error {
	potOID, err := a.ConvertToObjectID(potID)
	if err != nil {
		return err
	}

	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}

	return a.adapter.AddCompetitorInPot(ctx, potOID, competitorOID)
}

func (a *TournamentQueryerAdapter) RemoveCompetitorOfPot(ctx context.Context, potID, competitorID string) error {
	potOID, err := a.ConvertToObjectID(potID)
	if err != nil {
		return err
	}

	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}

	return a.adapter.RemoveCompetitorOfPot(ctx, potOID, competitorOID)
}

func (a *TournamentQueryerAdapter) SetCompetitorsInPots(ctx context.Context, tournamentID string, potDTOs []*tournament_dto.SetPotCompetitorDTOReq) error {
	potDAOs, tournamentOID, err := mappers.SetCompetitorsInPotsDTOtoDAO(potDTOs, tournamentID, a.ConvertToObjectID)
	if err != nil {
		return nil
	}

	return a.adapter.SetCompetitorsInPots(ctx, tournamentOID, potDAOs)
}

func (a *TournamentQueryerAdapter) VerifyMultipleTournamentPot(ctx context.Context, tournamentID string, potIDs []string) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	potOIDs, err := utils.ConvertToObjectIDs(&potIDs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.adapter.VerifyMultipleTournamentPot(ctx, tournamentOID, *potOIDs)
}

func (a *TournamentQueryerAdapter) AddPotInTournament(ctx context.Context, tournamentID, potID string) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	potOID, err := a.ConvertToObjectID(potID)
	if err != nil {
		return err
	}

	return a.adapter.AddPotInTournament(ctx, tournamentOID, potOID)
}

func (a *TournamentQueryerAdapter) RemovePotToTournament(ctx context.Context, tournamentID string, position int) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.adapter.RemovePotToTournament(ctx, tournamentOID, position)
}

func (a *TournamentQueryerAdapter) GetTournamentPotPositions(ctx context.Context, tournamentID string) ([]*tournament_dto.PotOrGroupPositionDTORes, error) {
	potPositionsDAO, err := a.adapter.GetTournamentPotPositions(ctx, tournamentID)
	if err != nil {
		return nil, err
	}

	potPositionsDTO := mappers.GetTournamentPotPositionsDAOtoDTO(potPositionsDAO)

	return potPositionsDTO, nil
}

func (a *TournamentQueryerAdapter) UpdatePotPositions(ctx context.Context, potID string, position int) error {
	potOID, err := a.ConvertToObjectID(potID)
	if err != nil {
		return err
	}

	return a.adapter.UpdatePotPositions(ctx, potOID, position)
}

func (a *TournamentQueryerAdapter) DeletePotByPosition(ctx context.Context, position int, tournamentID string) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.adapter.DeletePotByPosition(ctx, position, tournamentOID)
}

func (a *TournamentQueryerAdapter) VerifyNumberPotsInTournament(ctx context.Context, tournamentID string, position int) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.adapter.VerifyNumberPotsInTournament(ctx, tournamentOID, position)
}

func (a *TournamentQueryerAdapter) VerifyNumberGroupsInTournament(ctx context.Context, tournamentID string, position int) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.adapter.VerifyNumberGroupsInTournament(ctx, tournamentOID, position)
}

func (a *TournamentQueryerAdapter) AddGroupInTournament(ctx context.Context, tournamentID, groupID string) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	groupOID, err := a.ConvertToObjectID(groupID)
	if err != nil {
		return err
	}

	return a.adapter.AddGroupInTournament(ctx, tournamentOID, groupOID)
}

func (a *TournamentQueryerAdapter) UpdateGroupPositions(ctx context.Context, groupID string, position int) error {
	groupOID, err := a.ConvertToObjectID(groupID)
	if err != nil {
		return err
	}

	return a.adapter.UpdateGroupPositions(ctx, groupOID, position)
}

func (a *TournamentQueryerAdapter) RemoveGroupToTournament(ctx context.Context, tournamentID string, position int) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.adapter.RemoveGroupToTournament(ctx, tournamentOID, position)
}

func (a *TournamentQueryerAdapter) DeleteGroupByPosition(ctx context.Context, position int, tournamentID string) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.adapter.DeleteGroupByPosition(ctx, position, tournamentOID)
}

func (a *TournamentQueryerAdapter) GetTournamentGroupPositions(ctx context.Context, tournamentID string) ([]*tournament_dto.PotOrGroupPositionDTORes, error) {
	groupPositionsDAO, err := a.adapter.GetTournamentGroupPositions(ctx, tournamentID)
	if err != nil {
		return nil, err
	}

	potPositionsDTO := mappers.GetTournamentPotPositionsDAOtoDTO(groupPositionsDAO)

	return potPositionsDTO, nil
}

func (a *TournamentQueryerAdapter) GetTournamentGroupMatchesByPosition(ctx context.Context, position int, tournamentID string) ([]string, []string, error) {
	return a.adapter.GetTournamentGroupMatchesByPosition(ctx,position, tournamentID)
}

