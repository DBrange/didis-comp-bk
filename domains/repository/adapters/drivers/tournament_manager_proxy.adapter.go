package adapters

import (
	"context"
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	option_models "github.com/DBrange/didis-comp-bk/cmd/api/models/options/tournament"
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	availability_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability/dao"
	double_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/double/dao"
	double_elimination_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/double_elimination/dao"
	guest_user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/guest_user/dao"
	category_registration_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/category_registration/dao"
	competitor_match_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/competitor_match/dao"
	guest_competitor_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/guest_competitor/dao"
	tournament_registration_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/tournament_registration/dao"
	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	match_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/match/dao"
	pot_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/pot/dao"
	round_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/round/dao"
	single_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/single/dao"
	team_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/team/dao"
	tournament_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/tournament/dao"
	tournament_group_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/tournament_group/dao"
	"github.com/DBrange/didis-comp-bk/domains/repository/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TournamentManagerProxyAdapter struct {
	repository *repository.Repository
}

func NewTournamentManagerProxyAdapter(repository *repository.Repository) *TournamentManagerProxyAdapter {
	return &TournamentManagerProxyAdapter{
		repository: repository,
	}
}

func (a *TournamentManagerProxyAdapter) CreateLocation(ctx context.Context, locationInfoDAO *location_dao.CreateLocationDAOReq) (string, error) {
	return a.repository.CreateLocation(ctx, locationInfoDAO)
}

func (a *TournamentManagerProxyAdapter) VerifyOrganizerExists(ctx context.Context, organizerID string) error {
	return a.repository.VerifyOrganizerExists(ctx, organizerID)
}

func (a *TournamentManagerProxyAdapter) CreateTournament(
	ctx context.Context,
	tournamentInfoDAO *tournament_dao.CreateTournamentDAOReq,
	locationID string,
	options *option_models.OrganizeTournamentOptions,
	categoryID *string,
	organizerID string,
) (string, error) {
	return a.repository.CreateTournament(ctx, tournamentInfoDAO, locationID, options, categoryID, organizerID)
}

func (a *TournamentManagerProxyAdapter) VerifyCategoryExists(ctx context.Context, categoryID string) error {
	return a.repository.VerifyCategoryExists(ctx, categoryID)
}

func (a *TournamentManagerProxyAdapter) AddTournamentInCategory(ctx context.Context, categoryID string, tournamentID string) error {
	return a.repository.AddTournamentInCategory(ctx, categoryID, tournamentID)
}

func (a *TournamentManagerProxyAdapter) WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) error) error {
	return a.repository.WithTransaction(ctx, fn)
}

func (a *TournamentManagerProxyAdapter) CreateTournamentGroup(ctx context.Context, tournamentOID *primitive.ObjectID, position int) (string, error) {
	return a.repository.CreateTournamentGroup(ctx, tournamentOID, position)
}

func (a *TournamentManagerProxyAdapter) CreatePot(ctx context.Context, tournamentOID *primitive.ObjectID, position int) (string, error) {
	return a.repository.CreatePot(ctx, tournamentOID, position)
}

func (a *TournamentManagerProxyAdapter) CreateDoubleEliminationEmpty(ctx context.Context) (string, error) {
	return a.repository.CreateDoubleEliminationEmpty(ctx)
}

func (a *TournamentManagerProxyAdapter) TournamentGroupColl() *mongo.Collection {
	return a.repository.TournamentGroupColl()
}

func (a *TournamentManagerProxyAdapter) PotColl() *mongo.Collection {
	return a.repository.TournamentGroupColl()
}

func (a *TournamentManagerProxyAdapter) DoubleEliminationColl() *mongo.Collection {
	return a.repository.TournamentGroupColl()
}

func (a *TournamentManagerProxyAdapter) DeleteByID(ctx context.Context, mc *mongo.Collection, ID string, name string) error {
	return a.repository.DeleteByID(ctx, mc, ID, name)
}

func (a *TournamentManagerProxyAdapter) UpdateTournamentRelations(ctx context.Context, tournamentOID *primitive.ObjectID, tournamentDAO *tournament_dao.UpdateTournamentOptionsDAOReq, add bool) error {
	return a.repository.UpdateTournamentRelations(ctx, tournamentOID, tournamentDAO, add)
}

func (a *TournamentManagerProxyAdapter) ConvertToObjectID(ID string) (*primitive.ObjectID, error) {
	return a.repository.ConvertToObjectID(ID)
}

func (a *TournamentManagerProxyAdapter) CreateTournamentRegistration(ctx context.Context, tournamentRegistrationInfoDAO *tournament_registration_dao.CreateTournamentRegistrationDAOReq) error {
	return a.repository.CreateTournamentRegistration(ctx, tournamentRegistrationInfoDAO)
}

func (a *TournamentManagerProxyAdapter) CreateGuestUser(ctx context.Context, guestUserInfoDAO *guest_user_dao.CreateGuestUserDAOReq) (string, error) {
	return a.repository.CreateGuestUser(ctx, guestUserInfoDAO)
}

func (a *TournamentManagerProxyAdapter) CreateCompetitor(ctx context.Context, sport models.SPORT, competitorType models.COMPETITOR_TYPE, OID *primitive.ObjectID) (string, error) {
	return a.repository.CreateCompetitor(ctx, sport, competitorType, OID)
}

func (a *TournamentManagerProxyAdapter) CreateGuestCompetitor(ctx context.Context, guestCompetitorInfoDAO *guest_competitor_dao.CreateGuestCompetitorDAOReq) (string, error) {
	return a.repository.CreateGuestCompetitor(ctx, guestCompetitorInfoDAO)
}

func (a *TournamentManagerProxyAdapter) CreateMatch(ctx context.Context, match *match_dao.CreateMatchDAOReq) (string, error) {
	return a.repository.CreateMatch(ctx, match)

}

func (a *TournamentManagerProxyAdapter) CreateRound(ctx context.Context, round *round_dao.CreateRoundDAOReq) (string, error) {
	return a.repository.CreateRound(ctx, round)
}

func (a *TournamentManagerProxyAdapter) CreateDoubleElimination(ctx context.Context, doubleEliminationDAO *double_elimination_dao.CreateDoubleEliminationDAOReq) (string, error) {
	return a.repository.CreateDoubleElimination(ctx, doubleEliminationDAO)
}

func (a *TournamentManagerProxyAdapter) CreateSingle(ctx context.Context, singleInfoDAO *single_dao.CreateSingleDAOReq) (string, error) {
	return a.repository.CreateSingle(ctx, singleInfoDAO)
}

func (a *TournamentManagerProxyAdapter) CreateDouble(ctx context.Context, doubleInfoDAO *double_dao.CreateDoubleDAOReq) (string, error) {
	return a.repository.CreateDouble(ctx, doubleInfoDAO)
}

func (a *TournamentManagerProxyAdapter) CreateTeam(ctx context.Context, teamInfoDAO *team_dao.CreateTeamDAOReq) (string, error) {
	return a.repository.CreateTeam(ctx, teamInfoDAO)
}

func (a *TournamentManagerProxyAdapter) ListCompetitorsInTournament(
	ctx context.Context,
	tournamentID, categoryID,
	lastID string, limit int,
) ([]*tournament_registration_dao.GetCompetitorsInTournamentDAORes, error) {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return nil, err
	}

	var categoryOID *primitive.ObjectID
	if categoryID != "" {
		categoryOID, err = a.ConvertToObjectID(categoryID)
		if err != nil {
			return nil, err
		}
	} else {
		categoryOID = nil
	}

	var lastOID *primitive.ObjectID
	if lastID != "" {
		lastOID, err = a.ConvertToObjectID(lastID)
		if err != nil {
			return nil, err
		}
	} else {
		lastOID = nil
	}

	return a.repository.GetCompetitorsInTournament(ctx, tournamentOID, categoryOID, lastOID, limit)
}

func (a *TournamentManagerProxyAdapter) VerifyCompetitorExists(ctx context.Context, competitorOID *primitive.ObjectID) error {
	return a.repository.VerifyCompetitorExists(ctx, competitorOID)
}

func (a *TournamentManagerProxyAdapter) VerifyTournamentExists(ctx context.Context, tournamentOID *primitive.ObjectID) error {
	return a.repository.VerifyTournamentExists(ctx, tournamentOID)
}

func (a *TournamentManagerProxyAdapter) CreateCompetitorStats(ctx context.Context, competitorOID *primitive.ObjectID) error {
	return a.repository.CreateCompetitorStats(ctx, competitorOID)
}

func (a *TournamentManagerProxyAdapter) UpdateCompetitorMatch(ctx context.Context, matchOID *primitive.ObjectID, competitorMatchDAO *competitor_match_dao.UpdateCompetitorMatchDAOReq) error {
	return a.repository.UpdateCompetitorMatch(ctx, matchOID, competitorMatchDAO)
}

func (a *TournamentManagerProxyAdapter) VerifyMatchExists(ctx context.Context, matchOID *primitive.ObjectID) error {
	return a.repository.VerifyMatchExists(ctx, matchOID)
}

func (a *TournamentManagerProxyAdapter) CreateCompetitorMatch(ctx context.Context, competitorMatchDAO *competitor_match_dao.CreateCompetitorMatchDAOReq) error {
	return a.repository.CreateCompetitorMatch(ctx, competitorMatchDAO)
}

func (a *TournamentManagerProxyAdapter) UpdateRoundTotalPrize(ctx context.Context, roundOID *primitive.ObjectID, totalPrize float64) error {
	return a.repository.UpdateRoundTotalPrize(ctx, roundOID, totalPrize)
}

func (a *TournamentManagerProxyAdapter) VerifyRoundExists(ctx context.Context, roundOID *primitive.ObjectID) error {
	return a.repository.VerifyRoundExists(ctx, roundOID)
}

func (a *TournamentManagerProxyAdapter) GetRoundWithMatches(ctx context.Context, roundID, categoryID string) (*round_dao.GetRoundWithMatchesDAORes, error) {
	roundOID, err := a.ConvertToObjectID(roundID)
	if err != nil {
		return nil, err
	}

	var categoryOID *primitive.ObjectID
	if categoryID != "" {
		categoryOID, err = a.ConvertToObjectID(categoryID)
		if err != nil {
			return nil, err
		}
	} else {
		categoryOID = nil
	}

	return a.repository.GetRoundWithMatches(ctx, roundOID, categoryOID)
}
func (a *TournamentManagerProxyAdapter) GetPositionsBracketMatch(ctx context.Context, roundID string, organizeType models.ORGANIZE_TYPE, organizeBracket models.ORGANIZE_BRACKET) ([]*match_dao.GetPositionsBracketMatchDAORes, error) {
	roundOID, err := a.ConvertToObjectID(roundID)
	if err != nil {
		return nil, err
	}

	return a.repository.GetPositionsBracketMatch(ctx, roundOID, organizeType, organizeBracket)
}

func (a *TournamentManagerProxyAdapter) UpdateMultipleCompetitorMatches(ctx context.Context, competitorMatchDAOs []*competitor_match_dao.UpdateCompetitorMatchDAOReq) error {
	return a.repository.UpdateMultipleCompetitorMatches(ctx, competitorMatchDAOs)
}

func (a *TournamentManagerProxyAdapter) VerifyMatchesExist(ctx context.Context, matchOIDs []*primitive.ObjectID) error {
	return a.repository.VerifyMatchesExist(ctx, matchOIDs)
}

func (a *TournamentManagerProxyAdapter) VerifyMultipleCompetitorsExists(ctx context.Context, competitorOIDs []*primitive.ObjectID) error {
	return a.repository.VerifyMultipleCompetitorsExists(ctx, competitorOIDs)
}

func (a *TournamentManagerProxyAdapter) VerifyMatchesInRoundExits(ctx context.Context, roundOID *primitive.ObjectID) (bool, error) {
	return a.repository.VerifyMatchesInRoundExits(ctx, roundOID)
}

func (a *TournamentManagerProxyAdapter) SetWinnerInMatch(ctx context.Context, matchOID, competitorOID *primitive.ObjectID, result string) error {
	return a.repository.SetWinnerInMatch(ctx, matchOID, competitorOID, result)
}

func (a *TournamentManagerProxyAdapter) FindMatchID(ctx context.Context, position int, roundOID *primitive.ObjectID) (string, error) {
	return a.repository.FindMatchID(ctx, position, roundOID)
}

func (a *TournamentManagerProxyAdapter) AddMatchInTournament(ctx context.Context, tournamentOID, matchOID *primitive.ObjectID) error {
	return a.repository.AddMatchInTournament(ctx, tournamentOID, matchOID)
}

func (a *TournamentManagerProxyAdapter) AddMatchInCompetitorStats(ctx context.Context, competitorOID, matchOID *primitive.ObjectID) error {
	return a.repository.AddMatchInCompetitorStats(ctx, competitorOID, matchOID)
}

func (a *TournamentManagerProxyAdapter) UpdateCompetitorStats(ctx context.Context, competitorOID *primitive.ObjectID, winner bool) error {
	return a.repository.UpdateCompetitorStats(ctx, competitorOID, winner)
}

func (a *TournamentManagerProxyAdapter) IncrementTotalCompetitorsInTournament(ctx context.Context, tournamentOID *primitive.ObjectID) error {
	return a.repository.IncrementTotalCompetitorsInTournament(ctx, tournamentOID)
}

func (a *TournamentManagerProxyAdapter) VerifyTournamentsCapacity(ctx context.Context, tournamentOID *primitive.ObjectID) (bool, error) {
	return a.repository.VerifyTournamentsCapacity(ctx, tournamentOID)
}

func (a *TournamentManagerProxyAdapter) UpdateRoundPoints(ctx context.Context, roundOID *primitive.ObjectID, points int) error {
	return a.repository.UpdateRoundPoints(ctx, roundOID, points)
}

func (a *TournamentManagerProxyAdapter) UpdateTournamentInfo(ctx context.Context, tournamentOID *primitive.ObjectID, tournamentDAO *tournament_dao.UpdateTournamentInfoDAOReq) error {
	return a.repository.UpdateTournamentInfo(ctx, tournamentOID, tournamentDAO)
}

func (a *TournamentManagerProxyAdapter) AddTournamentWonInCompetitorStats(ctx context.Context, competitorOID, tournamentOID *primitive.ObjectID) error {
	return a.repository.AddTournamentWonInCompetitorStats(ctx, competitorOID, tournamentOID)
}

func (a *TournamentManagerProxyAdapter) GetRoundsWithCompetitors(ctx context.Context, tournamentID string) ([]*round_dao.GetRoundWithCompetitorsDAORes, error) {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return nil, err
	}

	return a.repository.GetRoundsWithCompetitors(ctx, tournamentOID)
}

func (a *TournamentManagerProxyAdapter) GetTournamentCompetitorIDs(ctx context.Context, tournamentID string) ([]string, error) {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return nil, err
	}

	return a.repository.GetTournamentCompetitorIDs(ctx, tournamentOID)
}

func (a *TournamentManagerProxyAdapter) GetCompetitorsOutCategory(ctx context.Context, categoryID string, competitorIDs []string) ([]string, error) {
	categoryOID, err := a.ConvertToObjectID(categoryID)
	if err != nil {
		return []string{}, err
	}

	competitorOIDs, err := utils.ConvertToObjectIDs(&competitorIDs, a.ConvertToObjectID)
	if err != nil {
		return []string{}, err
	}

	return a.repository.GetCompetitorsOutCategory(ctx, categoryOID, *competitorOIDs)
}

func (a *TournamentManagerProxyAdapter) CreateCategoryRegistration(ctx context.Context, categoryRegistrationDAO *category_registration_dao.CreateCategoryRegistrationDAOReq) error {
	return a.repository.CreateCategoryRegistration(ctx, categoryRegistrationDAO)
}

func (a *TournamentManagerProxyAdapter) IncrementTotalParticipants(ctx context.Context, categoryOID *primitive.ObjectID) error {
	return a.repository.IncrementTotalParticipants(ctx, categoryOID)
}

func (a *TournamentManagerProxyAdapter) GetCategoryRegistrationSortedByPoints(ctx context.Context, categoryID string) ([]*category_registration_dao.GetCategoryRegistrationSortedByPointsDAORes, error) {
	categoryOID, err := a.ConvertToObjectID(categoryID)
	if err != nil {
		return nil, err
	}

	return a.repository.GetCategoryRegistrationSortedByPoints(ctx, categoryOID)
}

func (a *TournamentManagerProxyAdapter) UpdateCategoryRegistrationCurrentPosition(ctx context.Context, categoryOID *primitive.ObjectID, categoryRegistration []*category_registration_dao.GetCategoryRegistrationSortedByPointsDAORes) error {
	return a.repository.UpdateCategoryRegistrationCurrentPosition(ctx, categoryOID, categoryRegistration)
}

func (a *TournamentManagerProxyAdapter) VerifyCompetitorsMatch(ctx context.Context, matchOID, competitorOID *primitive.ObjectID) error {
	return a.repository.VerifyCompetitorsMatch(ctx, matchOID, competitorOID)
}

func (a *TournamentManagerProxyAdapter) UpdateTournamentFinishDate(ctx context.Context, tournamentOID *primitive.ObjectID) error {
	return a.repository.UpdateTournamentFinishDate(ctx, tournamentOID)
}

func (a *TournamentManagerProxyAdapter) VerifyMatchPosition(ctx context.Context, matchOID *primitive.ObjectID, position int) error {
	return a.repository.VerifyMatchPosition(ctx, matchOID, position)
}

func (a *TournamentManagerProxyAdapter) GetRoundQuantityMatches(ctx context.Context, roundID string) (int, error) {
	roundOID, err := a.ConvertToObjectID(roundID)
	if err != nil {
		return 0, err
	}

	return a.repository.GetRoundQuantityMatches(ctx, roundOID)
}

func (a *TournamentManagerProxyAdapter) GetMatchPosition(ctx context.Context, matchID string) (int, error) {
	matchOID, err := a.ConvertToObjectID(matchID)
	if err != nil {
		return 0, err
	}

	return a.repository.GetMatchPosition(ctx, matchOID)
}

func (a *TournamentManagerProxyAdapter) GetRoundID(ctx context.Context, tournamentID string, round models.ROUND) (string, error) {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return "", err
	}

	return a.repository.GetRoundID(ctx, tournamentOID, round)
}

func (a *TournamentManagerProxyAdapter) AddPointsInMultipleCategoryRegistration(ctx context.Context, categoryOID *primitive.ObjectID, competitorOIDs []*primitive.ObjectID, points int) error {
	return a.repository.AddPointsInMultipleCategoryRegistration(ctx, categoryOID, competitorOIDs, points)
}

func (a *TournamentManagerProxyAdapter) AddPrizeInMultipleCompetitorStats(ctx context.Context, competitorOIDs []*primitive.ObjectID, prize float64) error {
	return a.repository.AddPrizeInMultipleCompetitorStats(ctx, competitorOIDs, prize)
}

func (a *TournamentManagerProxyAdapter) GetTournamentInfoToFinaliseIt(ctx context.Context, tournamentID string) (*tournament_dao.GetTournamentInfoToFinaliseItDAORes, error) {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return nil, err
	}

	return a.repository.GetTournamentInfoToFinaliseIt(ctx, tournamentOID)
}

func (a *TournamentManagerProxyAdapter) VerifyTournamentsAlreadyFinished(ctx context.Context, tournamentOID *primitive.ObjectID) error {
	return a.repository.VerifyTournamentsAlreadyFinished(ctx, tournamentOID)
}

func (a *TournamentManagerProxyAdapter) VerifyMatchAndRoundCoincidence(ctx context.Context, matchOID, roundOID *primitive.ObjectID, round models.ROUND) error {
	return a.repository.VerifyMatchAndRoundCoincidence(ctx, matchOID, roundOID, round)
}

func (a *TournamentManagerProxyAdapter) VerifyMultipleCompetitorsExistsInTournament(ctx context.Context, tournamentOID *primitive.ObjectID, competitorOIDs []*primitive.ObjectID) error {
	return a.repository.VerifyMultipleCompetitorsExistsInTournament(ctx, tournamentOID, competitorOIDs)
}

func (a *TournamentManagerProxyAdapter) VerifyCompetitorExistsInTournament(ctx context.Context, tournamentOID *primitive.ObjectID, competitorOID *primitive.ObjectID) error {
	return a.repository.VerifyCompetitorExistsInTournament(ctx, tournamentOID, competitorOID)
}

func (a *TournamentManagerProxyAdapter) AddCompetitorInGroup(ctx context.Context, groupOID, competitorOID *primitive.ObjectID) error {
	return a.repository.AddCompetitorInGroup(ctx, groupOID, competitorOID)
}

func (a *TournamentManagerProxyAdapter) AddCompetitorsToTournamentGroups(ctx context.Context, tournamentOID *primitive.ObjectID, groupDTOs []*tournament_group_dao.AddCompetitorsToTournamentGroupsDAOReq) error {
	return a.repository.AddCompetitorsToTournamentGroups(ctx, tournamentOID, groupDTOs)
}

func (a *TournamentManagerProxyAdapter) AddMatchInTournamentGroup(ctx context.Context, groupOID, tournamentOID, matchOID *primitive.ObjectID) error {
	return a.repository.AddMatchInTournamentGroup(ctx, groupOID, tournamentOID, matchOID)
}

func (a *TournamentManagerProxyAdapter) VerifyMultipleGroupsInTournament(ctx context.Context, tournamentOID *primitive.ObjectID, groupOIDs []*primitive.ObjectID) error {
	return a.repository.VerifyMultipleGroupsInTournament(ctx, tournamentOID, groupOIDs)
}

func (a *TournamentManagerProxyAdapter) VerifyRoundInTournament(ctx context.Context, roundOID, tournamentOID *primitive.ObjectID) error {
	return a.repository.VerifyRoundInTournament(ctx, roundOID, tournamentOID)
}

func (a *TournamentManagerProxyAdapter) AddMultipleMatchesInTournamentGroup(ctx context.Context, groupOID, tournamentOID *primitive.ObjectID, matchOIDs []*primitive.ObjectID) error {
	return a.repository.AddMultipleMatchesInTournamentGroup(ctx, groupOID, tournamentOID, matchOIDs)
}

func (a *TournamentManagerProxyAdapter) AddMultipleMatchesInTournament(ctx context.Context, tournamentOID *primitive.ObjectID, matchOIDs []*primitive.ObjectID) error {
	return a.repository.AddMultipleMatchesInTournament(ctx, tournamentOID, matchOIDs)
}

func (a *TournamentManagerProxyAdapter) VerifyTournamentGroupInTournament(ctx context.Context, tournamentOID *primitive.ObjectID, groupOIDs []*primitive.ObjectID) error {
	return a.repository.VerifyTournamentGroupInTournament(ctx, tournamentOID, groupOIDs)
}

func (a *TournamentManagerProxyAdapter) GetTournamentGroupMatches(ctx context.Context, groupID string) ([]string, []string, error) {
	groupOID, err := a.ConvertToObjectID(groupID)
	if err != nil {
		return nil, nil, err
	}

	return a.repository.GetTournamentGroupMatches(ctx, groupOID)
}

func (a *TournamentManagerProxyAdapter) RemoveMultipleTournamentMatches(ctx context.Context, tournamentOID *primitive.ObjectID, matchesToRemoveOIDs []*primitive.ObjectID) error {
	return a.repository.RemoveMultipleTournamentMatches(ctx, tournamentOID, matchesToRemoveOIDs)
}

func (a *TournamentManagerProxyAdapter) RemoveMultipleCompetitorStatsMatches(ctx context.Context, competitorIDs, matchesToRemove []*primitive.ObjectID) error {
	return a.repository.RemoveMultipleCompetitorStatsMatches(ctx, competitorIDs, matchesToRemove)
}

func (a *TournamentManagerProxyAdapter) DeleteMultipleMatches(ctx context.Context, matchesToRemove []*primitive.ObjectID) error {
	return a.repository.DeleteMultipleMatches(ctx, matchesToRemove)
}

func (a *TournamentManagerProxyAdapter) DeleteMultipleCompetitorMatches(ctx context.Context, matchesToRemove []*primitive.ObjectID) error {
	return a.repository.DeleteMultipleCompetitorMatches(ctx, matchesToRemove)
}

func (a *TournamentManagerProxyAdapter) VerifyTournamentPot(ctx context.Context, tournamentOID, potOID *primitive.ObjectID) error {
	return a.repository.VerifyTournamentPot(ctx, tournamentOID, potOID)
}

func (a *TournamentManagerProxyAdapter) AddCompetitorInPot(ctx context.Context, potOID, competitorOID *primitive.ObjectID) error {
	return a.repository.AddCompetitorInPot(ctx, potOID, competitorOID)
}

func (a *TournamentManagerProxyAdapter) RemoveCompetitorOfPot(ctx context.Context, potOID, competitorOID *primitive.ObjectID) error {
	return a.repository.RemoveCompetitorOfPot(ctx, potOID, competitorOID)
}

func (a *TournamentManagerProxyAdapter) VerifyMultipleTournamentPot(ctx context.Context, tournamentOID *primitive.ObjectID, potOIDs []*primitive.ObjectID) error {
	return a.repository.VerifyMultipleTournamentPot(ctx, tournamentOID, potOIDs)
}

func (a *TournamentManagerProxyAdapter) SetCompetitorsInPots(ctx context.Context, tournamentOID *primitive.ObjectID, potDTOs []*pot_dao.SetPotCompetitorDAOReq) error {
	return a.repository.SetCompetitorsInPots(ctx, tournamentOID, potDTOs)
}

func (a *TournamentManagerProxyAdapter) AddPotInTournament(ctx context.Context, tournamentOID, potOID *primitive.ObjectID) error {
	return a.repository.AddPotInTournament(ctx, tournamentOID, potOID)
}

func (a *TournamentManagerProxyAdapter) RemovePotToTournament(ctx context.Context, tournamentOID *primitive.ObjectID, position int) error {
	return a.repository.RemovePotToTournament(ctx, tournamentOID, position)
}

func (a *TournamentManagerProxyAdapter) GetTournamentPotPositions(ctx context.Context, tournamentID string) ([]*pot_dao.PotOrGroupPositionDAORes, error) {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return nil, err
	}

	return a.repository.GetTournamentPotPositions(ctx, tournamentOID)
}

func (a *TournamentManagerProxyAdapter) UpdatePotPositions(ctx context.Context, potOID *primitive.ObjectID, position int) error {
	return a.repository.UpdatePotPositions(ctx, potOID, position)
}

func (a *TournamentManagerProxyAdapter) DeletePotByPosition(ctx context.Context, position int, tournamentOID *primitive.ObjectID) error {
	return a.repository.DeletePotByPosition(ctx, position, tournamentOID)
}

func (a *TournamentManagerProxyAdapter) VerifyNumberPotsInTournament(ctx context.Context, tournamentOID *primitive.ObjectID, position int) error {
	return a.repository.VerifyNumberPotsInTournament(ctx, tournamentOID, position)
}

func (a *TournamentManagerProxyAdapter) VerifyNumberGroupsInTournament(ctx context.Context, tournamentOID *primitive.ObjectID, position int) error {
	return a.repository.VerifyNumberGroupsInTournament(ctx, tournamentOID, position)
}

func (a *TournamentManagerProxyAdapter) UpdateGroupPositions(ctx context.Context, groupOID *primitive.ObjectID, position int) error {
	return a.repository.UpdateGroupPositions(ctx, groupOID, position)
}

func (a *TournamentManagerProxyAdapter) AddGroupInTournament(ctx context.Context, tournamentOID, groupOID *primitive.ObjectID) error {
	return a.repository.AddGroupInTournament(ctx, tournamentOID, groupOID)
}

func (a *TournamentManagerProxyAdapter) RemoveGroupToTournament(ctx context.Context, tournamentOID *primitive.ObjectID, position int) error {
	return a.repository.RemoveGroupToTournament(ctx, tournamentOID, position)
}

func (a *TournamentManagerProxyAdapter) DeleteGroupByPosition(ctx context.Context, position int, tournamentOID *primitive.ObjectID) error {
	return a.repository.DeleteGroupByPosition(ctx, position, tournamentOID)
}

func (a *TournamentManagerProxyAdapter) GetTournamentGroupPositions(ctx context.Context, tournamentID string) ([]*pot_dao.PotOrGroupPositionDAORes, error) {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return nil, err
	}

	return a.repository.GetTournamentGroupPositions(ctx, tournamentOID)
}

func (a *TournamentManagerProxyAdapter) GetTournamentGroupMatchesByPosition(ctx context.Context, position int, tournamentID string) ([]string, []string, error) {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return nil, nil, err
	}

	return a.repository.GetTournamentGroupMatchesByPosition(ctx, position, tournamentOID)
}

func (a *TournamentManagerProxyAdapter) GetDoubleElimRoundID(ctx context.Context, tournamentID string, round models.ROUND) (string, error) {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return "", err
	}

	return a.repository.GetDoubleElimRoundID(ctx, tournamentOID, round)
}

func (a *TournamentManagerProxyAdapter) AddMatchInDoubleElim(ctx context.Context, doubleElimOID, matchOID *primitive.ObjectID) error {
	return a.repository.AddMatchInDoubleElim(ctx, doubleElimOID, matchOID)
}

func (a *TournamentManagerProxyAdapter) GetDoubleElimID(ctx context.Context, tournamentID string) (string, error) {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return "", err
	}

	return a.repository.GetDoubleElimID(ctx, tournamentOID)
}

func (a *TournamentManagerProxyAdapter) GetTournamentRoundNames(ctx context.Context, tournamentID string) ([]models.ROUND, error) {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return nil, err
	}

	return a.repository.GetTournamentRoundNames(ctx, tournamentOID)
}

func (a *TournamentManagerProxyAdapter) GetAllDoubleElimRoundIDs(ctx context.Context, doubleEliminationID string) ([]string, error) {
	doubleEliminationOID, err := a.ConvertToObjectID(doubleEliminationID)
	if err != nil {
		return nil, err
	}

	return a.repository.GetAllDoubleElimRoundIDs(ctx, doubleEliminationOID)
}

func (a *TournamentManagerProxyAdapter) GetDoubleElimInfoToFinaliseIt(ctx context.Context, doubleElimID string) (*double_elimination_dao.GetDoubleElimInfoToFinaliseItDAORes, error) {
	doubleElimOID, err := a.ConvertToObjectID(doubleElimID)
	if err != nil {
		return nil, err
	}

	return a.repository.GetDoubleElimInfoToFinaliseIt(ctx, doubleElimOID)
}

func (a *TournamentManagerProxyAdapter) GetDoubleElimCompetitorChampion(ctx context.Context, doubleElimID string) (string, error) {
	doubleElimOID, err := a.ConvertToObjectID(doubleElimID)
	if err != nil {
		return "", err
	}

	return a.repository.GetDoubleElimCompetitorChampion(ctx, doubleElimOID)
}

func (a *TournamentManagerProxyAdapter) GetCompetitorChampion(ctx context.Context, tournamentID string) (string, error) {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return "", err
	}

	return a.repository.GetCompetitorChampion(ctx, tournamentOID)
}

func (a *TournamentManagerProxyAdapter) GetMultipleAvailabilitiesByCompetitor(ctx context.Context, competitorIDs []string) ([][]*availability_dao.GetDailyAvailabilityByIDDAORes, error) {
	competitorOIDs, err := utils.ConvertToObjectIDs(&competitorIDs, a.ConvertToObjectID)
	if err != nil {
		return [][]*availability_dao.GetDailyAvailabilityByIDDAORes{}, err
	}

	return a.repository.GetMultipleAvailabilitiesByCompetitor(ctx, *competitorOIDs)
}

func (a *TournamentManagerProxyAdapter) UpdateMultipleMatchesDate(ctx context.Context, matchDates []*match_dao.MatchDateDAOReq) error {
	return a.repository.UpdateMultipleMatchesDate(ctx, matchDates)
}

func (a *TournamentManagerProxyAdapter) GetAvailabilityByTournamentID(ctx context.Context, tournamentID string) ([]*availability_dao.GetDailyAvailabilityByIDDAORes, error) {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return nil, err
	}

	return a.repository.GetAvailabilityByTournamentID(ctx, tournamentOID)
}

func (a *TournamentManagerProxyAdapter) GetTournamentAvailavility(ctx context.Context, tournamentID string) (*tournament_dao.TournamentAvailabilityDAO, error) {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return nil, err
	}

	return a.repository.GetTournamentAvailavility(ctx, tournamentOID)
}

func (a *TournamentManagerProxyAdapter) CreateAvailability(ctx context.Context, userOID, competitorOID, tournamentOID *primitive.ObjectID) error {
	return a.repository.CreateAvailability(ctx, userOID, competitorOID, tournamentOID)
}

func (a *TournamentManagerProxyAdapter) GetAllDatesMatchesFromTournament(ctx context.Context, tournamentID string) ([]time.Time, error) {
	tournamentOID, err := a.ConvertToObjectID(tournamentID)
	if err != nil {
		return nil, err
	}

	return a.repository.GetAllDatesMatchesFromTournament(ctx, tournamentOID)
}

func (a *TournamentManagerProxyAdapter) UpdateMatchDate(ctx context.Context, matchID *primitive.ObjectID, date *time.Time) error {
	return a.repository.UpdateMatchDate(ctx, matchID, date)
}

func (a *TournamentManagerProxyAdapter) VerifyCompetitorIDInCompetitorUser(ctx context.Context, competitorIDs []*primitive.ObjectID) (bool, error) {
	return a.repository.VerifyCompetitorIDInCompetitorUser(ctx, competitorIDs)
}
