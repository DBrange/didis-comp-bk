package ports

import (
	"context"
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	optional_models "github.com/DBrange/didis-comp-bk/cmd/api/models/options/tournament"
	availability_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability/dao"
	double_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/double/dao"
	double_elimination_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/double_elimination/dao"
	guest_user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/guest_user/dao"
	category_registration_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/category_registration/dao"
	competitor_match_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/competitor_match/dao"
	competitor_user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/competitor_user/dao"
	follower_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/follower/dao"
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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ForManagingTournament interface {
	CreateLocation(ctx context.Context, locationInfoDAO *location_dao.CreateLocationDAOReq) (string, error)
	VerifyOrganizerExists(ctx context.Context, organizerID string) error
	CreateTournament(
		ctx context.Context,
		tournamentInfoDAO *tournament_dao.CreateTournamentDAOReq,
		locationID string,
		options *optional_models.OrganizeTournamentOptions,
		categoryID *string,
		organizerID string,
	) (string, error)
	VerifyCategoryExists(ctx context.Context, categoryID string) error
	AddTournamentInCategory(ctx context.Context, categoryID string, tournamentID string) error
	WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) error) error
	CreateTournamentGroup(ctx context.Context, TournamentOID *primitive.ObjectID, position int) (string, error)
	CreatePot(ctx context.Context, TournamentOID *primitive.ObjectID, position int) (string, error)
	CreateDoubleEliminationEmpty(ctx context.Context) (string, error)
	TournamentGroupColl() *mongo.Collection
	PotColl() *mongo.Collection
	DoubleEliminationColl() *mongo.Collection
	DeleteByID(ctx context.Context, mc *mongo.Collection, ID string, name string) error
	UpdateTournamentRelations(ctx context.Context, tournamentOID *primitive.ObjectID, tournamentDAO *tournament_dao.UpdateTournamentOptionsDAOReq, add bool) error
	ConvertToObjectID(ID string) (*primitive.ObjectID, error)
	CreateTournamentRegistration(ctx context.Context, tournamentRegistrationInfoDTO *tournament_registration_dao.CreateTournamentRegistrationDAOReq) error
	CreateGuestUser(ctx context.Context, guestUserInfoDAO *guest_user_dao.CreateGuestUserDAOReq) (string, error)
	CreateCompetitor(ctx context.Context, sport models.SPORT, competitorType models.COMPETITOR_TYPE, OID *primitive.ObjectID) (string, error)
	CreateGuestCompetitor(ctx context.Context, guestCompetitorInfoDAO *guest_competitor_dao.CreateGuestCompetitorDAOReq) (string, error)
	CreateMatch(ctx context.Context, match *match_dao.CreateMatchDAOReq) (string, error)
	CreateRound(ctx context.Context, round *round_dao.CreateRoundDAOReq) (string, error)
	CreateDoubleElimination(ctx context.Context, doubleEliminationDAO *double_elimination_dao.CreateDoubleEliminationDAOReq) (string, error)
	CreateSingle(ctx context.Context, singleInfoDAO *single_dao.CreateSingleDAOReq) (string, error)
	CreateDouble(ctx context.Context, doubleInfoDAO *double_dao.CreateDoubleDAOReq) (string, error)
	CreateTeam(ctx context.Context, teamInfoDAO *team_dao.CreateTeamDAOReq) (string, error)
	ListCompetitorsInTournament(ctx context.Context, tournamentID, categoryID, lastID string, limit int, getAll bool) ([]*tournament_registration_dao.GetCompetitorsInTournamentDAORes, error)
	VerifyCompetitorExists(ctx context.Context, competitorOID *primitive.ObjectID) error
	VerifyTournamentExists(ctx context.Context, tournamentOID *primitive.ObjectID) error
	CreateCompetitorStats(ctx context.Context, competitorOID *primitive.ObjectID) error
	UpdateCompetitorMatch(ctx context.Context, matchOID *primitive.ObjectID, competitorMatchDAO *competitor_match_dao.UpdateCompetitorMatchDAOReq) error
	VerifyMatchExists(ctx context.Context, matchOID *primitive.ObjectID) error
	CreateCompetitorMatch(ctx context.Context, competitorMatchDAO *competitor_match_dao.CreateCompetitorMatchDAOReq) error
	UpdateRoundTotalPrize(ctx context.Context, roundOID *primitive.ObjectID, totalPrize float64) error
	VerifyRoundExists(ctx context.Context, roundOID *primitive.ObjectID) error
	GetRoundWithMatches(ctx context.Context, roundID, categoryID string) (*round_dao.GetRoundWithMatchesDAORes, error)
	GetPositionsBracketMatch(ctx context.Context, roundOID string, organizeType models.ORGANIZE_TYPE, organizeBracket models.ORGANIZE_BRACKET) ([]*match_dao.GetPositionsBracketMatchDAORes, error)
	UpdateMultipleCompetitorMatches(ctx context.Context, competitorMatchDAOs []*competitor_match_dao.UpdateCompetitorMatchDAOReq) error
	VerifyMatchesExist(ctx context.Context, matchOIDs []*primitive.ObjectID) error
	VerifyMultipleCompetitorsExists(ctx context.Context, competitorOIDs []*primitive.ObjectID) error
	SetWinnerInMatch(ctx context.Context, matchOID, competitorOID *primitive.ObjectID, result string) error
	VerifyMatchesInRoundExits(ctx context.Context, roundOID *primitive.ObjectID) (bool, error)
	FindMatchID(ctx context.Context, position int, roundOID *primitive.ObjectID) (string, error)
	AddMatchInTournament(ctx context.Context, tournamentOID, matchOID *primitive.ObjectID) error
	AddMatchInCompetitorStats(ctx context.Context, competitorOID, matchOID *primitive.ObjectID) error
	UpdateCompetitorStats(ctx context.Context, competitorOID *primitive.ObjectID, winner bool) error
	IncrementTotalCompetitorsInTournament(ctx context.Context, tournamentOID *primitive.ObjectID) error
	VerifyTournamentsCapacity(ctx context.Context, tournamentOID *primitive.ObjectID) (bool, error)
	UpdateRoundPoints(ctx context.Context, roundOID *primitive.ObjectID, points int) error
	UpdateTournamentInfo(ctx context.Context, tournamentOID *primitive.ObjectID, tournamentDAO *tournament_dao.UpdateTournamentInfoDAOReq) error
	AddTournamentWonInCompetitorStats(ctx context.Context, competitorOID, tournamentOID *primitive.ObjectID) error
	GetTournamentCompetitorIDs(ctx context.Context, tournamentID string) ([]string, error)
	CreateCategoryRegistration(ctx context.Context, categoryRegistrationDAO *category_registration_dao.CreateCategoryRegistrationDAOReq) error
	IncrementTotalParticipants(ctx context.Context, categoryOID *primitive.ObjectID) error
	GetCategoryRegistrationSortedByPoints(ctx context.Context, categoryID string) ([]*category_registration_dao.GetCategoryRegistrationSortedByPointsDAORes, error)
	UpdateCategoryRegistrationCurrentPosition(ctx context.Context, categoryOID *primitive.ObjectID, categoryRegistration []*category_registration_dao.GetCategoryRegistrationSortedByPointsDAORes) error
	VerifyCompetitorsMatch(ctx context.Context, matchOID, competitorOID *primitive.ObjectID) error
	UpdateTournamentFinishDate(ctx context.Context, tournamentOID *primitive.ObjectID) error
	VerifyMatchPosition(ctx context.Context, matchOID *primitive.ObjectID, position int) error
	GetRoundQuantityMatches(ctx context.Context, roundOID string) (int, error)
	GetMatchPosition(ctx context.Context, matchOID string) (int, error)
	GetRoundID(ctx context.Context, tournamentOID string, nextRound models.ROUND) (string, error)
	GetRoundsWithCompetitors(ctx context.Context, tournamentID string) ([]*round_dao.GetRoundWithCompetitorsDAORes, error)
	AddPointsInMultipleCategoryRegistration(ctx context.Context, categoryOID *primitive.ObjectID, competitorOIDs []*primitive.ObjectID, points int) error
	AddPrizeInMultipleCompetitorStats(ctx context.Context, competitorOIDs []*primitive.ObjectID, prize float64) error
	GetCompetitorsOutCategory(ctx context.Context, categoryID string, competitorIDs []string) ([]string, error)
	GetTournamentInfoToFinaliseIt(ctx context.Context, tournamentOID string) (*tournament_dao.GetTournamentInfoToFinaliseItDAORes, error)
	VerifyTournamentsAlreadyFinished(ctx context.Context, tournamentOID *primitive.ObjectID) error
	VerifyMatchAndRoundCoincidence(ctx context.Context, matchOID, roundOID *primitive.ObjectID, round models.ROUND) error
	VerifyMultipleCompetitorsExistsInTournament(ctx context.Context, tournamentOID *primitive.ObjectID, competitorOIDs []*primitive.ObjectID) error
	VerifyCompetitorExistsInTournament(ctx context.Context, tournamentOID *primitive.ObjectID, competitorOID *primitive.ObjectID) error
	AddCompetitorInGroup(ctx context.Context, groupOID, competitorOID *primitive.ObjectID) error
	AddCompetitorsToTournamentGroups(ctx context.Context, tournamentOID *primitive.ObjectID, groupDTOs []*tournament_group_dao.AddCompetitorsToTournamentGroupsDAOReq) error
	AddMatchInTournamentGroup(ctx context.Context, groupOID, tournamentOID, matchOID *primitive.ObjectID) error
	VerifyMultipleGroupsInTournament(ctx context.Context, tournamentOID *primitive.ObjectID, groupOIDs []*primitive.ObjectID) error
	VerifyRoundInTournament(ctx context.Context, roundOID, tournamentOID *primitive.ObjectID) error
	AddMultipleMatchesInTournamentGroup(ctx context.Context, groupOID, tournamentOID *primitive.ObjectID, matchOIDs []*primitive.ObjectID) error
	AddMultipleMatchesInTournament(ctx context.Context, tournamentOID *primitive.ObjectID, matchOIDs []*primitive.ObjectID) error
	VerifyTournamentGroupInTournament(ctx context.Context, tournamentOID *primitive.ObjectID, groupOIDs []*primitive.ObjectID) error
	GetTournamentGroupMatches(ctx context.Context, groupID string) ([]string, []string, error)
	RemoveMultipleTournamentMatches(ctx context.Context, tournamentOID *primitive.ObjectID, matchesToRemoveOIDs []*primitive.ObjectID) error
	RemoveMultipleCompetitorStatsMatches(ctx context.Context, competitorIDs, matchesToRemove []*primitive.ObjectID) error
	DeleteMultipleMatches(ctx context.Context, matchesToRemove []*primitive.ObjectID) error
	DeleteMultipleCompetitorMatches(ctx context.Context, matchesToRemove []*primitive.ObjectID) error
	VerifyTournamentPot(ctx context.Context, tournamentOID, potOID *primitive.ObjectID) error
	AddCompetitorInPot(ctx context.Context, potOID, competitorOID *primitive.ObjectID) error
	RemoveCompetitorOfPot(ctx context.Context, potOID, competitorOID *primitive.ObjectID) error
	SetCompetitorsInPots(ctx context.Context, tournamentOID *primitive.ObjectID, potDTOs []*pot_dao.SetPotCompetitorDAOReq) error
	VerifyMultipleTournamentPot(ctx context.Context, tournamentOID *primitive.ObjectID, potOIDs []*primitive.ObjectID) error
	AddPotInTournament(ctx context.Context, tournamentOID, potOID *primitive.ObjectID) error
	RemovePotToTournament(ctx context.Context, tournamentOID *primitive.ObjectID, position int) error
	GetTournamentPotPositions(ctx context.Context, tournamentOID string) ([]*pot_dao.PotOrGroupPositionDAORes, error)
	UpdatePotPositions(ctx context.Context, potOID *primitive.ObjectID, position int) error
	DeletePotByPosition(ctx context.Context, position int, tournamentOID *primitive.ObjectID) error
	VerifyNumberPotsInTournament(ctx context.Context, tournamentOID *primitive.ObjectID, position int) error
	VerifyNumberGroupsInTournament(ctx context.Context, tournamentOID *primitive.ObjectID, position int) error
	AddGroupInTournament(ctx context.Context, tournamentOID, groupOID *primitive.ObjectID) error
	UpdateGroupPositions(ctx context.Context, groupOID *primitive.ObjectID, position int) error
	RemoveGroupToTournament(ctx context.Context, tournamentOID *primitive.ObjectID, position int) error
	DeleteGroupByPosition(ctx context.Context, position int, tournamentOID *primitive.ObjectID) error
	GetTournamentGroupPositions(ctx context.Context, tournamentOID string) ([]*pot_dao.PotOrGroupPositionDAORes, error)
	GetTournamentGroupMatchesByPosition(ctx context.Context, position int, tournamentOID string) ([]string, []string, error)
	GetDoubleElimRoundID(ctx context.Context, tournamentOID string, round models.ROUND) (string, error)
	AddMatchInDoubleElim(ctx context.Context, doubleElimOID, matchOID *primitive.ObjectID) error
	GetDoubleElimID(ctx context.Context, tournamentID string) (string, error)
	GetTournamentRoundNames(ctx context.Context, tournamentOID string) ([]models.ROUND, error)
	GetAllDoubleElimRoundIDs(ctx context.Context, doubleEliminationOID string) ([]string, error)
	GetDoubleElimInfoToFinaliseIt(ctx context.Context, doubleElimOID string) (*double_elimination_dao.GetDoubleElimInfoToFinaliseItDAORes, error)
	GetDoubleElimCompetitorChampion(ctx context.Context, doubleElimOID string) (string, error)
	GetCompetitorChampion(ctx context.Context, tournamentOID string) (string, error)
	GetMultipleAvailabilitiesByCompetitor(ctx context.Context, competitorOIDs []string) ([][]*availability_dao.GetDailyAvailabilityByIDDAORes, error)
	UpdateMultipleMatchesDate(ctx context.Context, matchDates []*match_dao.MatchDateDAOReq) error
	GetAvailabilityByTournamentID(ctx context.Context, tournamentID string) ([]*availability_dao.GetDailyAvailabilityByIDDAORes, error)
	GetTournamentAvailavility(ctx context.Context, tournamentID string) (*tournament_dao.TournamentAvailabilityDAO, error)
	CreateAvailability(ctx context.Context, userOID, competitorOID, tournamentOID *primitive.ObjectID) error
	GetAllDatesMatchesFromTournament(ctx context.Context, tournamentID string) ([]time.Time, error)
	UpdateMatchDate(ctx context.Context, matchID *primitive.ObjectID, date *time.Time) error
	VerifyCompetitorIDInCompetitorUser(ctx context.Context, competitorIDs []*primitive.ObjectID) (bool, error)
	UpdateTournamentStartDate(ctx context.Context, tournamentOID *primitive.ObjectID) error
	GetUserTournaments(
		ctx context.Context,
		userID string,
		sport models.SPORT,
		limit int,
		lastID string,
	) (*competitor_user_dao.GetUserTournamentsDAORes, error)
	GetTournamentPrimaryInfo(ctx context.Context, tournamentID string) (*tournament_dao.GetTournamentPrimaryInfoDAORes, error)
	GetCompetitorsByNameInTournament(
		ctx context.Context,
		tournamentOID, categoryOID string,
		name string,
		limit int,
	) ([]*tournament_registration_dao.GetCompetitorsInTournamentDAORes, error)
	GetTournamentTotalCompetitors(ctx context.Context, tournamentOID *primitive.ObjectID) (int, error)
	GetCategoryIDOfTournament(ctx context.Context, tournamentID string) (*primitive.ObjectID, error)
	GetCompetitorsFollowed(ctx context.Context, userOID string, name string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) ([]*follower_dao.GetCompetitorFollowedDAORes, error)
	VerifyUserExists(ctx context.Context, userOID *primitive.ObjectID) error
	GetAvailabilityDailySlice(ctx context.Context, userID, competitorID string) ([]*availability_dao.GetDailyAvailabilityByIDDAORes, error)
	CreateAvailabilityForCompetitor(ctx context.Context, competitorID string, dailyAvailability []*availability_dao.CreateDailyAvailability) error
	CreateCompetitorUser(ctx context.Context, userOID *primitive.ObjectID, competitorOID *primitive.ObjectID) error
	GetTournamentFilters(ctx context.Context, tournamentID string) (*tournament_dao.GetTournamentFiltersDAORes, error)
	GetTournamentsInOrganizer(
		ctx context.Context,
		organizerID string,
		sport models.SPORT,
		limit int,
		lastID string, // El último ID desde el que comenzar la siguiente página
	) (*competitor_user_dao.GetUserTournamentsDAORes, error)
	AddTournamentInOrganizer(ctx context.Context, organizerOID, tournamentOID *primitive.ObjectID) error
	DeleteTournamentRegistration(ctx context.Context, tournamentRegistrationID string) error
	DecrementTotalCompetitorsInTournament(ctx context.Context, tournamentOID *primitive.ObjectID) error
	GetTournamentRegistrationByCompetitorAndTournamentID(ctx context.Context, tournamentOID, competitorOID *primitive.ObjectID) (string, error)
	GetTournamentMatchesByID(ctx context.Context, tournamentID string) ([]*primitive.ObjectID, error)
	GetCompetitorIDsFromMatches(ctx context.Context, matches []string) ([]*primitive.ObjectID, error)
	GetCompetitorIDByMatchAndPosition(ctx context.Context, matchID string, position int) (*primitive.ObjectID, error)
	GetRoundGroups(ctx context.Context, roundID, categoryID string) (*round_dao.GetRoundGroupsDAORes, error)
	GetDailyAvailabilityTournamentID(ctx context.Context, tournamentID string, day string) (*availability_dao.GetDailyAvailabilityByIDDAORes, *primitive.ObjectID, error)
	UpdateAvailability(ctx context.Context, availabilityID string, availabilityInfoDAO *availability_dao.UpdateDailyAvailabilityDAOReq) error
	GetTournamentGroupsIDs(ctx context.Context, tournamentID string) ([]*primitive.ObjectID, error)
	UpdateTournamentAvailability(
		ctx context.Context,
		tournamentOID *primitive.ObjectID,
		availableCourts int,
		averageHours int,
	) error
	GetTournamentSportsInOrganizer(ctx context.Context, organizerID string) ([]models.SPORT, error)
	 GetMatchByID(ctx context.Context, matchID string, categoryID string) (*match_dao.GetMatchDAORes, error)
	 GetMatchCategoryID(ctx context.Context, matchID string) (*primitive.ObjectID, error)
	 GetTournamentCompetitorIDsInMatches(ctx context.Context, tournamentID string) ([]string, error)
}
