package drivens

import (
	"context"
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	option_models "github.com/DBrange/didis-comp-bk/cmd/api/models/options/tournament"
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"

	chat_ports "github.com/DBrange/didis-comp-bk/domains/chat/ports/drivers"
	repo_ports "github.com/DBrange/didis-comp-bk/domains/repository/ports/drivers"
	"github.com/DBrange/didis-comp-bk/domains/tournament/adapters/mappers"
	tournament_dto "github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TournamentQuerierAdapter struct {
	repo_adapter repo_ports.ForManagingTournament
	chat_adapter chat_ports.ForChat
}

func NewTournamentQuerierAdapter(repo_adapter repo_ports.ForManagingTournament, chat_adapter chat_ports.ForChat) *TournamentQuerierAdapter {
	return &TournamentQuerierAdapter{
		repo_adapter: repo_adapter,
		chat_adapter: chat_adapter,
	}
}

func (a *TournamentQuerierAdapter) CreateLocation(ctx context.Context, locationDTO *tournament_dto.CreateLocationDTOReq) (string, error) {
	locationDAO := mappers.CreateLocationDTOtoDAO(locationDTO)

	return a.repo_adapter.CreateLocation(ctx, locationDAO)
}

func (a *TournamentQuerierAdapter) VerifyOrganizerExists(ctx context.Context, organizerID string) error {
	return a.repo_adapter.VerifyOrganizerExists(ctx, organizerID)
}

func (a *TournamentQuerierAdapter) CreateTournament(
	ctx context.Context,
	tournamentDTO *tournament_dto.CreateTournamentDTOReq,
	locationID string,
	options *option_models.OrganizeTournamentOptions,
	categoryID *string,
	organizerID string,
) (string, error) {
	tournamentDAO := mappers.CreateTournamentDTOtoDAO(tournamentDTO)

	return a.repo_adapter.CreateTournament(ctx, tournamentDAO, locationID, options, categoryID, organizerID)
}

func (a *TournamentQuerierAdapter) VerifyCategoryExists(ctx context.Context, categoryID string) error {
	return a.repo_adapter.VerifyCategoryExists(ctx, categoryID)
}

func (a *TournamentQuerierAdapter) AddTournamentInCategory(ctx context.Context, categoryID string, tournamentID string) error {
	return a.repo_adapter.AddTournamentInCategory(ctx, categoryID, tournamentID)
}

func (a *TournamentQuerierAdapter) WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) error) error {
	return a.repo_adapter.WithTransaction(ctx, fn)
}

func (a *TournamentQuerierAdapter) CreateTournamentGroup(ctx context.Context, tournamentID string, position int) (string, error) {
	tournamentIOD, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return "", nil
	}

	return a.repo_adapter.CreateTournamentGroup(ctx, tournamentIOD, position)
}

func (a *TournamentQuerierAdapter) CreatePot(ctx context.Context, tournamentID string, position int) (string, error) {
	tournamentIOD, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return "", nil
	}

	return a.repo_adapter.CreatePot(ctx, tournamentIOD, position)
}

func (a *TournamentQuerierAdapter) CreateDoubleEliminationEmpty(ctx context.Context) (string, error) {
	return a.repo_adapter.CreateDoubleEliminationEmpty(ctx)
}

func (a *TournamentQuerierAdapter) TournamentGroupColl() *mongo.Collection {
	return a.repo_adapter.TournamentGroupColl()
}

func (a *TournamentQuerierAdapter) PotColl() *mongo.Collection {
	return a.repo_adapter.PotColl()
}

func (a *TournamentQuerierAdapter) DoubleEliminationColl() *mongo.Collection {
	return a.repo_adapter.DoubleEliminationColl()
}

func (a *TournamentQuerierAdapter) DeleteByID(ctx context.Context, mc *mongo.Collection, ID string, name string) error {
	return a.repo_adapter.DeleteByID(ctx, mc, ID, name)
}

func (a *TournamentQuerierAdapter) UpdateTournamentRelations(
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

	return a.repo_adapter.UpdateTournamentRelations(ctx, tournamentOID, tournamentDAO, add)
}

func (a *TournamentQuerierAdapter) ConvertToObjectID(ID string) (*primitive.ObjectID, error) {
	return a.repo_adapter.ConvertToObjectID(ID)
}

func (a *TournamentQuerierAdapter) CreateTournamentRegistration(ctx context.Context, tournamentRegistrationDTO *tournament_dto.CreateTournamentRegistrationDTOReq) error {
	tournamentRegistrationDAO, err := mappers.CreateTournamentRegistrationDTOtoDAO(tournamentRegistrationDTO, a.ConvertToObjectID)
	if err != nil {
		return nil
	}

	return a.repo_adapter.CreateTournamentRegistration(ctx, tournamentRegistrationDAO)
}

func (a *TournamentQuerierAdapter) CreateGuestUser(ctx context.Context, guestUserDTO *tournament_dto.CreateGuestUserDTOReq) (string, error) {
	guestUserDAO := mappers.CreateGuestUserDTOtoDAO(guestUserDTO)

	return a.repo_adapter.CreateGuestUser(ctx, guestUserDAO)
}

func (a *TournamentQuerierAdapter) CreateCompetitor(ctx context.Context, sport models.SPORT, competitorType models.COMPETITOR_TYPE, ID string) (string, error) {
	OID, err := a.ConvertToObjectID(ID)
	if err != nil {
		return "", err
	}
	return a.repo_adapter.CreateCompetitor(ctx, sport, competitorType, OID)
}

func (a *TournamentQuerierAdapter) CreateGuestCompetitor(ctx context.Context, guestCompetitorDTO *tournament_dto.CreateGuestCompetitorDTOReq) (string, error) {
	guestCompetitorDAO, err := mappers.CreateGuestCompetitorDTOtoDAO(guestCompetitorDTO, a.ConvertToObjectID)
	if err != nil {
		return "", err
	}

	return a.repo_adapter.CreateGuestCompetitor(ctx, guestCompetitorDAO)
}

func (a *TournamentQuerierAdapter) CreateMatch(ctx context.Context, matchDTO *tournament_dto.CreateMatchDTOReq) (string, error) {
	matchDAO, err := mappers.CreateMatchDTOtoDAO(matchDTO, a.ConvertToObjectID)
	if err != nil {
		return "", err
	}

	return a.repo_adapter.CreateMatch(ctx, matchDAO)
}

func (a *TournamentQuerierAdapter) CreateRound(ctx context.Context, roundDTO *tournament_dto.CreateRoundDTOReq) (string, error) {
	roundDAO, err := mappers.CreateRoundDTOtoDAO(roundDTO, a.ConvertToObjectID)
	if err != nil {
		return "", err
	}

	return a.repo_adapter.CreateRound(ctx, roundDAO)
}

func (a *TournamentQuerierAdapter) CreateDoubleElimination(ctx context.Context, doubleEliminationDTO *tournament_dto.CreateDoubleEliminationDTOReq) (string, error) {
	doubleEliminationDAO, err := mappers.CreateDoubleEliminationDTOtoDAO(doubleEliminationDTO, a.ConvertToObjectID)
	if err != nil {
		return "", err
	}

	return a.repo_adapter.CreateDoubleElimination(ctx, doubleEliminationDAO)
}

func (a *TournamentQuerierAdapter) CreateSingle(ctx context.Context, singleDTO *tournament_dto.CreateSingleDTOReq) (string, error) {
	singleDAO := mappers.CreateSingleDTOtoDAO(singleDTO)

	return a.repo_adapter.CreateSingle(ctx, singleDAO)
}

func (a *TournamentQuerierAdapter) CreateDouble(ctx context.Context, doubleDTO *tournament_dto.CreateDoubleDTOReq) (string, error) {
	doubleDAO := mappers.CreateDoubleDTOtoDAO(doubleDTO)

	return a.repo_adapter.CreateDouble(ctx, doubleDAO)
}

func (a *TournamentQuerierAdapter) CreateTeam(ctx context.Context, teamDTO *tournament_dto.CreateTeamDTOReq) (string, error) {
	teamDAO, err := mappers.CreateTeamDTOtoDAO(teamDTO, a.ConvertToObjectID)
	if err != nil {
		return "", nil
	}

	return a.repo_adapter.CreateTeam(ctx, teamDAO)
}

func (a *TournamentQuerierAdapter) GetCompetitorsInTournament(
	ctx context.Context,
	tournamentID, categoryID,
	lastID string, limit int,
) ([]*tournament_dto.GetCompetitorsInTournamentDTORes, error) {
	competitorsDAO, err := a.repo_adapter.ListCompetitorsInTournament(ctx, tournamentID, categoryID, lastID, limit)
	if err != nil {
		return nil, err
	}

	competitorsDTO := mappers.GetCompetitorsInTournamentDAOtoDTO(competitorsDAO)

	return competitorsDTO, nil
}

func (a *TournamentQuerierAdapter) VerifyCompetitorExists(ctx context.Context, competitorID string) error {
	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}
	return a.repo_adapter.VerifyCompetitorExists(ctx, competitorOID)
}

func (a *TournamentQuerierAdapter) VerifyTournamentExists(ctx context.Context, tournamentID string) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.repo_adapter.VerifyTournamentExists(ctx, tournamentOID)
}

func (a *TournamentQuerierAdapter) CreateCompetitorStats(ctx context.Context, competitorID string) error {
	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}

	return a.repo_adapter.CreateCompetitorStats(ctx, competitorOID)
}

func (a *TournamentQuerierAdapter) UpdateCompetitorMatch(ctx context.Context, matchID string, competitorMatchDTO *tournament_dto.UpdateCompetitorMatchDTOReq) error {
	competitorMatchDAO, matchOID, err := mappers.UpdateCompetitorMatchDTOtoDAO(competitorMatchDTO, matchID, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.repo_adapter.UpdateCompetitorMatch(ctx, matchOID, competitorMatchDAO)
}

func (a *TournamentQuerierAdapter) VerifyMatchExists(ctx context.Context, matchID string) error {
	matchOID, err := a.ConvertToObjectID(matchID)
	if err != nil {
		return err
	}

	return a.repo_adapter.VerifyMatchExists(ctx, matchOID)
}

func (a *TournamentQuerierAdapter) CreateCompetitorMatch(ctx context.Context, competitorMatchDTO *tournament_dto.CreateCompetitorMatchDTOReq) error {
	competitorMatchDAO, err := mappers.CreateCompetitorMatchDTOtoDAO(competitorMatchDTO, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.repo_adapter.CreateCompetitorMatch(ctx, competitorMatchDAO)
}

func (a *TournamentQuerierAdapter) UpdateRoundTotalPrize(ctx context.Context, roundID string, totalPrize float64) error {
	roundOID, err := a.ConvertToObjectID(roundID)
	if err != nil {
		return err
	}

	return a.repo_adapter.UpdateRoundTotalPrize(ctx, roundOID, totalPrize)
}

func (a *TournamentQuerierAdapter) VerifyRoundExists(ctx context.Context, roundID string) error {
	roundOID, err := a.ConvertToObjectID(roundID)
	if err != nil {
		return err
	}

	return a.repo_adapter.VerifyRoundExists(ctx, roundOID)
}

func (a *TournamentQuerierAdapter) GetRoundWithMatches(ctx context.Context, roundID, categoryID string) (*tournament_dto.GetRoundWithMatchesDTORes, error) {
	roundDAO, err := a.repo_adapter.GetRoundWithMatches(ctx, roundID, categoryID)
	if err != nil {
		return nil, err
	}

	roundDTO := mappers.MapRoundWithMatchesDAOToDTO(roundDAO)

	return roundDTO, nil
}

func (a *TournamentQuerierAdapter) GetPositionsBracketMatch(ctx context.Context, roundID string, organizeType models.ORGANIZE_TYPE, organizeBracket models.ORGANIZE_BRACKET) ([]*tournament_dto.GetPositionsBracketMatchDTORes, error) {
	positionsDAO, err := a.repo_adapter.GetPositionsBracketMatch(ctx, roundID, organizeType, organizeBracket)
	if err != nil {
		return nil, err
	}

	positiondDTO := mappers.GetPositionsBracketMatchDAOtoDTO(positionsDAO)

	return positiondDTO, nil
}

func (a *TournamentQuerierAdapter) UpdateMultipleCompetitorMatches(ctx context.Context, competitorsDTOs []*tournament_dto.UpdateCompetitorMatchDTOReq) error {
	competitorsDAOs, err := mappers.UpdateMultipleCompetitorMatchesDTOtoDAO(competitorsDTOs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.repo_adapter.UpdateMultipleCompetitorMatches(ctx, competitorsDAOs)
}

func (a *TournamentQuerierAdapter) VerifyMatchesExist(ctx context.Context, matchOIDs []string) error {
	competitorsDAOs, err := utils.ConvertToObjectIDs(&matchOIDs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.repo_adapter.VerifyMatchesExist(ctx, *competitorsDAOs)
}

func (a *TournamentQuerierAdapter) VerifyMultipleCompetitorsExists(ctx context.Context, competitorOIDs []string) error {
	competitorsDAOs, err := utils.ConvertToObjectIDs(&competitorOIDs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.repo_adapter.VerifyMultipleCompetitorsExists(ctx, *competitorsDAOs)
}

func (a *TournamentQuerierAdapter) VerifyMatchesInRoundExits(ctx context.Context, roundID string) (bool, error) {
	roundOID, err := a.ConvertToObjectID(roundID)
	if err != nil {
		return false, err
	}

	return a.repo_adapter.VerifyMatchesInRoundExits(ctx, roundOID)
}

func (a *TournamentQuerierAdapter) SetWinnerInMatch(ctx context.Context, matchID, competitorID, result string) error {
	matchOID, err := a.ConvertToObjectID(matchID)
	if err != nil {
		return err
	}

	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}

	return a.repo_adapter.SetWinnerInMatch(ctx, matchOID, competitorOID, result)
}

func (a *TournamentQuerierAdapter) FindMatchID(ctx context.Context, position int, roundID string) (string, error) {
	roundOID, err := a.ConvertToObjectID(roundID)
	if err != nil {
		return "", err
	}

	return a.repo_adapter.FindMatchID(ctx, position, roundOID)
}

func (a *TournamentQuerierAdapter) AddMatchInTournament(ctx context.Context, tournamentID, matchID string) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	matchOID, err := a.ConvertToObjectID(matchID)
	if err != nil {
		return err
	}

	return a.repo_adapter.AddMatchInTournament(ctx, tournamentOID, matchOID)
}

func (a *TournamentQuerierAdapter) AddMatchInCompetitorStats(ctx context.Context, competitorID, matchID string) error {
	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}

	matchOID, err := a.ConvertToObjectID(matchID)
	if err != nil {
		return err
	}

	return a.repo_adapter.AddMatchInCompetitorStats(ctx, competitorOID, matchOID)
}

func (a *TournamentQuerierAdapter) UpdateCompetitorStats(ctx context.Context, competitorID string, winner bool) error {
	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}

	return a.repo_adapter.UpdateCompetitorStats(ctx, competitorOID, winner)
}

func (a *TournamentQuerierAdapter) IncrementTotalCompetitorsInTournament(ctx context.Context, tournamentID string) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.repo_adapter.IncrementTotalCompetitorsInTournament(ctx, tournamentOID)
}

func (a *TournamentQuerierAdapter) VerifyTournamentsCapacity(ctx context.Context, tournamentID string) (bool, error) {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return false, err
	}

	return a.repo_adapter.VerifyTournamentsCapacity(ctx, tournamentOID)
}

func (a *TournamentQuerierAdapter) UpdateRoundPoints(ctx context.Context, roundID string, points int) error {
	roundOID, err := a.ConvertToObjectID(roundID)
	if err != nil {
		return err
	}

	return a.repo_adapter.UpdateRoundPoints(ctx, roundOID, points)
}

func (a *TournamentQuerierAdapter) UpdateTournamentInfo(ctx context.Context, tournamentID string, tournamentDTO *tournament_dto.UpdateTournamentInfoDTOReq) error {
	tournamentDAO, tournamentOID, err := mappers.UpdateTournamentInfoDTOtoDAO(tournamentDTO, tournamentID, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.repo_adapter.UpdateTournamentInfo(ctx, tournamentOID, tournamentDAO)
}

func (a *TournamentQuerierAdapter) AddTournamentWonInCompetitorStats(ctx context.Context, competitorID, tournamentID string) error {
	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}

	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.repo_adapter.AddTournamentWonInCompetitorStats(ctx, competitorOID, tournamentOID)
}

func (a *TournamentQuerierAdapter) GetRoundsWithCompetitors(ctx context.Context, tournamentID string) ([]*tournament_dto.GetRoundWithCompetitorsDTORes, error) {
	roundDAO, err := a.repo_adapter.GetRoundsWithCompetitors(ctx, tournamentID)
	if err != nil {
		return nil, err
	}

	roundDTO := mappers.GetRoundsWithCompetitorsDAOtoDTO(roundDAO)

	return roundDTO, nil
}

func (a *TournamentQuerierAdapter) GetTournamentCompetitorIDs(ctx context.Context, tournamentID string) ([]string, error) {
	return a.repo_adapter.GetTournamentCompetitorIDs(ctx, tournamentID)
}

func (a *TournamentQuerierAdapter) GetCompetitorsOutCategory(ctx context.Context, categoryID string, competitorIDs []string) ([]string, error) {
	return a.repo_adapter.GetCompetitorsOutCategory(ctx, categoryID, competitorIDs)
}

func (a *TournamentQuerierAdapter) CreateCategoryRegistration(ctx context.Context, categoryRegistrationDTO *tournament_dto.CreateCategoryRegistrationDTOReq) error {
	categoryRegistrationDAO, err := mappers.CreateCategoryRegistrationDTOtoDAO(categoryRegistrationDTO, a.ConvertToObjectID)
	if err != nil {
		return nil
	}

	return a.repo_adapter.CreateCategoryRegistration(ctx, categoryRegistrationDAO)
}

func (a *TournamentQuerierAdapter) IncrementTotalParticipants(ctx context.Context, categoryID string) error {
	categoryOID, err := a.ConvertToObjectID(categoryID)
	if err != nil {
		return err
	}

	return a.repo_adapter.IncrementTotalParticipants(ctx, categoryOID)
}

func (a *TournamentQuerierAdapter) GetCategoryRegistrationSortedByPoints(ctx context.Context, categoryID string) ([]*tournament_dto.GetCategoryRegistrationSortedByPointsDTORes, error) {
	categoryRegistrationSortedDAO, err := a.repo_adapter.GetCategoryRegistrationSortedByPoints(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	categoryRegistrationSortedDTO := mappers.GetCategoryRegistrationSortedByPointsDAOtoDTO(categoryRegistrationSortedDAO)

	return categoryRegistrationSortedDTO, nil
}

func (a *TournamentQuerierAdapter) UpdateCategoryRegistrationCurrentPosition(ctx context.Context, categoryID string, categoryRegistrationDTO []*tournament_dto.GetCategoryRegistrationSortedByPointsDTORes) error {
	categoryRegistrationDAO, categoryOID, err := mappers.UpdateCategoryRegistrationCurrentPositionDTOtoDAO(categoryRegistrationDTO, categoryID, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.repo_adapter.UpdateCategoryRegistrationCurrentPosition(ctx, categoryOID, categoryRegistrationDAO)
}

func (a *TournamentQuerierAdapter) VerifyCompetitorsMatch(ctx context.Context, matchID, competitorID string) error {
	matchOID, err := a.ConvertToObjectID(matchID)
	if err != nil {
		return err
	}

	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}

	return a.repo_adapter.VerifyCompetitorsMatch(ctx, matchOID, competitorOID)
}

func (a *TournamentQuerierAdapter) UpdateTournamentFinishDate(ctx context.Context, tournamentID string) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.repo_adapter.UpdateTournamentFinishDate(ctx, tournamentOID)
}

func (a *TournamentQuerierAdapter) VerifyMatchPosition(ctx context.Context, matchID string, position int) error {
	matchOID, err := a.ConvertToObjectID(matchID)
	if err != nil {
		return err
	}

	return a.repo_adapter.VerifyMatchPosition(ctx, matchOID, position)
}

func (a *TournamentQuerierAdapter) GetRoundQuantityMatches(ctx context.Context, roundID string) (int, error) {
	return a.repo_adapter.GetRoundQuantityMatches(ctx, roundID)
}

func (a *TournamentQuerierAdapter) GetMatchPosition(ctx context.Context, matchID string) (int, error) {
	return a.repo_adapter.GetMatchPosition(ctx, matchID)
}

func (a *TournamentQuerierAdapter) GetRoundID(ctx context.Context, tournamentID string, round models.ROUND) (string, error) {
	return a.repo_adapter.GetRoundID(ctx, tournamentID, round)
}

func (a *TournamentQuerierAdapter) AddPointsInMultipleCategoryRegistration(ctx context.Context, categoryID string, competitorIDs []string, points int) error {
	categoryOID, err := a.ConvertToObjectID(categoryID)
	if err != nil {
		return err
	}

	competitorOIDs, err := utils.ConvertToObjectIDs(&competitorIDs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.repo_adapter.AddPointsInMultipleCategoryRegistration(ctx, categoryOID, *competitorOIDs, points)
}

func (a *TournamentQuerierAdapter) AddPrizeInMultipleCompetitorStats(ctx context.Context, competitorIDs []string, prize float64) error {
	competitorOIDs, err := utils.ConvertToObjectIDs(&competitorIDs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.repo_adapter.AddPrizeInMultipleCompetitorStats(ctx, *competitorOIDs, prize)
}

func (a *TournamentQuerierAdapter) GetTournamentInfoToFinaliseIt(ctx context.Context, tournamentID string) (*tournament_dto.GetTournamentInfoToFinaliseItDTORes, error) {
	tournamentInfoDAO, err := a.repo_adapter.GetTournamentInfoToFinaliseIt(ctx, tournamentID)
	if err != nil {
		return nil, err
	}

	tournamentInfoDTO := mappers.GetTournamentInfoToFinaliseItDAOtoDTO(tournamentInfoDAO)

	return tournamentInfoDTO, nil
}

func (a *TournamentQuerierAdapter) VerifyTournamentsAlreadyFinished(ctx context.Context, tournamentID string) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.repo_adapter.VerifyTournamentsAlreadyFinished(ctx, tournamentOID)
}

func (a *TournamentQuerierAdapter) VerifyMatchAndRoundCoincidence(ctx context.Context, matchID, roundID string, round models.ROUND) error {
	matchOID, err := a.ConvertToObjectID(matchID)
	if err != nil {
		return err
	}

	roundOID, err := a.ConvertToObjectID(roundID)
	if err != nil {
		return err
	}

	return a.repo_adapter.VerifyMatchAndRoundCoincidence(ctx, matchOID, roundOID, round)
}

func (a *TournamentQuerierAdapter) VerifyMultipleCompetitorsExistsInTournament(ctx context.Context, tournamentID string, competitorIDs []string) error {
	competitorOIDs, err := utils.ConvertToObjectIDs(&competitorIDs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.repo_adapter.VerifyMultipleCompetitorsExistsInTournament(ctx, tournamentOID, *competitorOIDs)
}

func (a *TournamentQuerierAdapter) VerifyCompetitorExistsInTournament(ctx context.Context, tournamentID string, competitorID string) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}

	return a.repo_adapter.VerifyCompetitorExistsInTournament(ctx, tournamentOID, competitorOID)
}

func (a *TournamentQuerierAdapter) AddCompetitorInGroup(ctx context.Context, groupID, competitorID string) error {
	groupOID, err := a.ConvertToObjectID(groupID)
	if err != nil {
		return err
	}

	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}

	return a.repo_adapter.AddCompetitorInGroup(ctx, groupOID, competitorOID)
}

func (a *TournamentQuerierAdapter) AddCompetitorsToTournamentGroups(ctx context.Context, tournamentID string, groupDTOs []*tournament_dto.AddCompetitorsToTournamentGroupsDTOReq) error {
	groupDAOs, tournamentOID, err := mappers.AddCompetitorsToTournamentGroupsDTOtoDAO(groupDTOs, tournamentID, a.ConvertToObjectID)
	if err != nil {
		return nil
	}

	return a.repo_adapter.AddCompetitorsToTournamentGroups(ctx, tournamentOID, groupDAOs)
}

func (a *TournamentQuerierAdapter) AddMatchInTournamentGroup(ctx context.Context, groupID, tournamentID, matchID string) error {
	groupOID, tournamentOID, matchOID, err := mappers.AddMatchInTournamentGroupDTOtoDAO(groupID, tournamentID, matchID, a.ConvertToObjectID)
	if err != nil {
		return nil
	}

	return a.repo_adapter.AddMatchInTournamentGroup(ctx, groupOID, tournamentOID, matchOID)
}

func (a *TournamentQuerierAdapter) VerifyMultipleGroupsInTournament(ctx context.Context, tournamentID string, groupIDs []string) error {
	groupOIDs, err := utils.ConvertToObjectIDs(&groupIDs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.repo_adapter.VerifyMultipleGroupsInTournament(ctx, tournamentOID, *groupOIDs)
}

func (a *TournamentQuerierAdapter) VerifyRoundInTournament(ctx context.Context, roundID, tournamentID string) error {
	roundOID, err := a.ConvertToObjectID(roundID)
	if err != nil {
		return err
	}

	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.repo_adapter.VerifyRoundInTournament(ctx, roundOID, tournamentOID)
}

func (a *TournamentQuerierAdapter) AddMultipleMatchesInTournamentGroup(ctx context.Context, groupID, tournamentID string, matchIDs []string) error {
	groupOID, tournamentOID, matchOIDs, err := mappers.AddMultipleMatchesInTournamentGroupDTOtoDAO(groupID, tournamentID, matchIDs, a.ConvertToObjectID)
	if err != nil {
		return nil
	}

	return a.repo_adapter.AddMultipleMatchesInTournamentGroup(ctx, groupOID, tournamentOID, matchOIDs)
}

func (a *TournamentQuerierAdapter) AddMultipleMatchesInTournament(ctx context.Context, tournamentID string, matchIDs []string) error {
	tournamentOID, matchOIDs, err := mappers.AddMultipleMatchesInTournamentDTOtoDAO(tournamentID, matchIDs, a.ConvertToObjectID)
	if err != nil {
		return nil
	}

	return a.repo_adapter.AddMultipleMatchesInTournament(ctx, tournamentOID, matchOIDs)
}

func (a *TournamentQuerierAdapter) VerifyTournamentGroupInTournament(ctx context.Context, tournamentID string, groupIDs []string) error {
	groupOIDs, err := utils.ConvertToObjectIDs(&groupIDs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.repo_adapter.VerifyTournamentGroupInTournament(ctx, tournamentOID, *groupOIDs)
}

func (a *TournamentQuerierAdapter) GetTournamentGroupMatches(ctx context.Context, groupID string) ([]string, []string, error) {
	return a.repo_adapter.GetTournamentGroupMatches(ctx, groupID)
}

func (a *TournamentQuerierAdapter) RemoveMultipleTournamentMatches(ctx context.Context, tournamentID string, matchesToRemoveIDs []string) error {
	matchesToRemoveOIDs, err := utils.ConvertToObjectIDs(&matchesToRemoveIDs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.repo_adapter.RemoveMultipleTournamentMatches(ctx, tournamentOID, *matchesToRemoveOIDs)
}

func (a *TournamentQuerierAdapter) RemoveMultipleCompetitorStatsMatches(ctx context.Context, competitorIDs, matchesToRemoveIDs []string) error {
	competitorOIDs, err := utils.ConvertToObjectIDs(&competitorIDs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	matchesToRemoveOIDs, err := utils.ConvertToObjectIDs(&matchesToRemoveIDs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.repo_adapter.RemoveMultipleCompetitorStatsMatches(ctx, *competitorOIDs, *matchesToRemoveOIDs)
}

func (a *TournamentQuerierAdapter) DeleteMultipleMatches(ctx context.Context, matchesToRemoveIDs []string) error {
	matchesToRemoveOIDs, err := utils.ConvertToObjectIDs(&matchesToRemoveIDs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.repo_adapter.DeleteMultipleMatches(ctx, *matchesToRemoveOIDs)
}

func (a *TournamentQuerierAdapter) DeleteMultipleCompetitorMatches(ctx context.Context, matchesToRemoveIDs []string) error {
	matchesToRemoveOIDs, err := utils.ConvertToObjectIDs(&matchesToRemoveIDs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.repo_adapter.DeleteMultipleCompetitorMatches(ctx, *matchesToRemoveOIDs)
}

func (a *TournamentQuerierAdapter) VerifyTournamentPot(ctx context.Context, tournamentID, potID string) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	potOID, err := a.ConvertToObjectID(potID)
	if err != nil {
		return err
	}

	return a.repo_adapter.VerifyTournamentPot(ctx, tournamentOID, potOID)
}

func (a *TournamentQuerierAdapter) AddCompetitorInPot(ctx context.Context, potID, competitorID string) error {
	potOID, err := a.ConvertToObjectID(potID)
	if err != nil {
		return err
	}

	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}

	return a.repo_adapter.AddCompetitorInPot(ctx, potOID, competitorOID)
}

func (a *TournamentQuerierAdapter) RemoveCompetitorOfPot(ctx context.Context, potID, competitorID string) error {
	potOID, err := a.ConvertToObjectID(potID)
	if err != nil {
		return err
	}

	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}

	return a.repo_adapter.RemoveCompetitorOfPot(ctx, potOID, competitorOID)
}

func (a *TournamentQuerierAdapter) SetCompetitorsInPots(ctx context.Context, tournamentID string, potDTOs []*tournament_dto.SetPotCompetitorDTOReq) error {
	potDAOs, tournamentOID, err := mappers.SetCompetitorsInPotsDTOtoDAO(potDTOs, tournamentID, a.ConvertToObjectID)
	if err != nil {
		return nil
	}

	return a.repo_adapter.SetCompetitorsInPots(ctx, tournamentOID, potDAOs)
}

func (a *TournamentQuerierAdapter) VerifyMultipleTournamentPot(ctx context.Context, tournamentID string, potIDs []string) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	potOIDs, err := utils.ConvertToObjectIDs(&potIDs, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.repo_adapter.VerifyMultipleTournamentPot(ctx, tournamentOID, *potOIDs)
}

func (a *TournamentQuerierAdapter) AddPotInTournament(ctx context.Context, tournamentID, potID string) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	potOID, err := a.ConvertToObjectID(potID)
	if err != nil {
		return err
	}

	return a.repo_adapter.AddPotInTournament(ctx, tournamentOID, potOID)
}

func (a *TournamentQuerierAdapter) RemovePotToTournament(ctx context.Context, tournamentID string, position int) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.repo_adapter.RemovePotToTournament(ctx, tournamentOID, position)
}

func (a *TournamentQuerierAdapter) GetTournamentPotPositions(ctx context.Context, tournamentID string) ([]*tournament_dto.PotOrGroupPositionDTORes, error) {
	potPositionsDAO, err := a.repo_adapter.GetTournamentPotPositions(ctx, tournamentID)
	if err != nil {
		return nil, err
	}

	potPositionsDTO := mappers.GetTournamentPotPositionsDAOtoDTO(potPositionsDAO)

	return potPositionsDTO, nil
}

func (a *TournamentQuerierAdapter) UpdatePotPositions(ctx context.Context, potID string, position int) error {
	potOID, err := a.ConvertToObjectID(potID)
	if err != nil {
		return err
	}

	return a.repo_adapter.UpdatePotPositions(ctx, potOID, position)
}

func (a *TournamentQuerierAdapter) DeletePotByPosition(ctx context.Context, position int, tournamentID string) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.repo_adapter.DeletePotByPosition(ctx, position, tournamentOID)
}

func (a *TournamentQuerierAdapter) VerifyNumberPotsInTournament(ctx context.Context, tournamentID string, position int) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.repo_adapter.VerifyNumberPotsInTournament(ctx, tournamentOID, position)
}

func (a *TournamentQuerierAdapter) VerifyNumberGroupsInTournament(ctx context.Context, tournamentID string, position int) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.repo_adapter.VerifyNumberGroupsInTournament(ctx, tournamentOID, position)
}

func (a *TournamentQuerierAdapter) AddGroupInTournament(ctx context.Context, tournamentID, groupID string) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	groupOID, err := a.ConvertToObjectID(groupID)
	if err != nil {
		return err
	}

	return a.repo_adapter.AddGroupInTournament(ctx, tournamentOID, groupOID)
}

func (a *TournamentQuerierAdapter) UpdateGroupPositions(ctx context.Context, groupID string, position int) error {
	groupOID, err := a.ConvertToObjectID(groupID)
	if err != nil {
		return err
	}

	return a.repo_adapter.UpdateGroupPositions(ctx, groupOID, position)
}

func (a *TournamentQuerierAdapter) RemoveGroupToTournament(ctx context.Context, tournamentID string, position int) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.repo_adapter.RemoveGroupToTournament(ctx, tournamentOID, position)
}

func (a *TournamentQuerierAdapter) DeleteGroupByPosition(ctx context.Context, position int, tournamentID string) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	return a.repo_adapter.DeleteGroupByPosition(ctx, position, tournamentOID)
}

func (a *TournamentQuerierAdapter) GetTournamentGroupPositions(ctx context.Context, tournamentID string) ([]*tournament_dto.PotOrGroupPositionDTORes, error) {
	groupPositionsDAO, err := a.repo_adapter.GetTournamentGroupPositions(ctx, tournamentID)
	if err != nil {
		return nil, err
	}

	potPositionsDTO := mappers.GetTournamentPotPositionsDAOtoDTO(groupPositionsDAO)

	return potPositionsDTO, nil
}

func (a *TournamentQuerierAdapter) GetTournamentGroupMatchesByPosition(ctx context.Context, position int, tournamentID string) ([]string, []string, error) {
	return a.repo_adapter.GetTournamentGroupMatchesByPosition(ctx, position, tournamentID)
}

func (a *TournamentQuerierAdapter) GetDoubleElimRoundID(ctx context.Context, tournamentOID string, round models.ROUND) (string, error) {
	return a.repo_adapter.GetDoubleElimRoundID(ctx, tournamentOID, round)
}

func (a *TournamentQuerierAdapter) AddMatchInDoubleElim(ctx context.Context, doubleElimID, matchID string) error {
	doubleElimOID, err := a.ConvertToObjectID(doubleElimID)
	if err != nil {
		return err
	}

	matchOID, err := a.ConvertToObjectID(matchID)
	if err != nil {
		return err
	}

	return a.repo_adapter.AddMatchInDoubleElim(ctx, doubleElimOID, matchOID)
}

func (a *TournamentQuerierAdapter) GetDoubleElimID(ctx context.Context, tournamentID string) (string, error) {
	return a.repo_adapter.GetDoubleElimID(ctx, tournamentID)
}

func (a *TournamentQuerierAdapter) GetTournamentRoundNames(ctx context.Context, tournamentID string) ([]models.ROUND, error) {
	return a.repo_adapter.GetTournamentRoundNames(ctx, tournamentID)
}

func (a *TournamentQuerierAdapter) GetAllDoubleElimRoundIDs(ctx context.Context, doubleEliminationID string) ([]string, error) {
	return a.repo_adapter.GetAllDoubleElimRoundIDs(ctx, doubleEliminationID)
}

func (a *TournamentQuerierAdapter) GetDoubleElimInfoToFinaliseIt(ctx context.Context, doubleElimID string) (*tournament_dto.GetDoubleElimInfoToFinaliseItDTORes, error) {
	doubleElimInfoDAO, err := a.repo_adapter.GetDoubleElimInfoToFinaliseIt(ctx, doubleElimID)
	if err != nil {
		return nil, err
	}

	doubleElimInfoDTO := mappers.GetDoubleElimInfoToFinaliseItDAOtoDTO(doubleElimInfoDAO)

	return doubleElimInfoDTO, err
}

func (a *TournamentQuerierAdapter) GetDoubleElimCompetitorChampion(ctx context.Context, doubleElimOID string) (string, error) {
	return a.repo_adapter.GetDoubleElimCompetitorChampion(ctx, doubleElimOID)
}

func (a *TournamentQuerierAdapter) GetCompetitorChampion(ctx context.Context, tournamentOID string) (string, error) {
	return a.repo_adapter.GetCompetitorChampion(ctx, tournamentOID)
}

func (a *TournamentQuerierAdapter) GetMultipleAvailabilitiesByCompetitor(ctx context.Context, competitorIDs []string) ([][]*tournament_dto.GetDailyAvailabilityByIDDTORes, error) {
	availabilityDAOs, err := a.repo_adapter.GetMultipleAvailabilitiesByCompetitor(ctx, competitorIDs)
	if err != nil {
		return nil, err
	}

	availabilityDTOs := mappers.GetMultipleAvailabilitiesByCompetitor(availabilityDAOs)

	return availabilityDTOs, nil
}

func (a *TournamentQuerierAdapter) UpdateMultipleMatchesDate(ctx context.Context, matchesDateDTO []*tournament_dto.MatchDateDTOReq) error {
	matchesDateDAO, err := mappers.UpdateMultipleMatchesDateDTOtoDAO(matchesDateDTO, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.repo_adapter.UpdateMultipleMatchesDate(ctx, matchesDateDAO)
}

func (a *TournamentQuerierAdapter) GetAvailabilityByTournamentID(ctx context.Context, tournamentID string) ([]*tournament_dto.GetDailyAvailabilityByIDDTORes, error) {
	dailyAvailabilityDAO, err := a.repo_adapter.GetAvailabilityByTournamentID(ctx, tournamentID)
	if err != nil {
		return nil, err
	}

	dailyAvailabilityDTO := mappers.GetAvailabilityByTournamentIDDAOtoDTO(dailyAvailabilityDAO)

	return dailyAvailabilityDTO, nil
}

func (a *TournamentQuerierAdapter) GetTournamentAvailavility(ctx context.Context, tournamentID string) (*tournament_dto.TournamentAvailabilityDTO, error) {
	dailyAvailabilityDAO, err := a.repo_adapter.GetTournamentAvailavility(ctx, tournamentID)
	if err != nil {
		return nil, err
	}

	dailyAvailabilityDTO := mappers.GetTournamentAvailavilityDAOtoDTO(dailyAvailabilityDAO)

	return dailyAvailabilityDTO, nil
}

func (a *TournamentQuerierAdapter) CreateAvailability(ctx context.Context, userID, competitorID, tournamentID *string) error {
	userOID, competitorOID, tournamentOID, err := mappers.CreateAvailabilityDTOtODAO(userID, competitorID, tournamentID, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.repo_adapter.CreateAvailability(ctx, userOID, competitorOID, tournamentOID)
}

func (a *TournamentQuerierAdapter) GetAllDatesMatchesFromTournament(ctx context.Context, tournamentID string) ([]time.Time, error) {
	dates, err := a.repo_adapter.GetAllDatesMatchesFromTournament(ctx, tournamentID)
	if err != nil {
		return nil, err
	}

	return dates, nil
}

func (a *TournamentQuerierAdapter) UpdateMatchDate(ctx context.Context, matchID string, date *time.Time) error {
	matchOID, err := a.ConvertToObjectID(matchID)
	if err != nil {
		return err
	}

	return a.repo_adapter.UpdateMatchDate(ctx, matchOID, date)
}

func (a *TournamentQuerierAdapter) CreateMatchChat(ctx context.Context, matchID string, competitorIDs []string, userID string) error {
	return a.chat_adapter.CreateMatchChat(ctx, matchID, competitorIDs, userID)
}

func (a *TournamentQuerierAdapter) VerifyCompetitorIDInCompetitorUser(ctx context.Context, competitorIDs []string) (bool, error) {
	competitorOIDs, err := utils.ConvertToObjectIDs(&competitorIDs, a.ConvertToObjectID)
	if err != nil {
		return false, err
	}

	return a.repo_adapter.VerifyCompetitorIDInCompetitorUser(ctx, *competitorOIDs)
}

func (a *TournamentQuerierAdapter) UpdateTournamentStartDate(ctx context.Context, tournamentID string) error {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil{
		return err
	}

	return a.repo_adapter.UpdateTournamentStartDate(ctx, tournamentOID)
}
