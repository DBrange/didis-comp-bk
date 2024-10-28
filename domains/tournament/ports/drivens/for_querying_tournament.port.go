package ports

import (
	"context"
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	option_models "github.com/DBrange/didis-comp-bk/cmd/api/models/options/tournament"
	tournament_dto "github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ForQueryingTournament interface {
	CreateLocation(ctx context.Context, locationInfoDTO *tournament_dto.CreateLocationDTOReq) (string, error)
	VerifyOrganizerExists(ctx context.Context, organizerID string) error
	CreateTournament(
		ctx context.Context,
		tournamentInfoDTO *tournament_dto.CreateTournamentDTOReq,
		locationID string,
		options *option_models.OrganizeTournamentOptions,
		categoryID *string,
		organizerID string,
	) (string, error)
	VerifyCategoryExists(ctx context.Context, categoryID string) error
	AddTournamentInCategory(ctx context.Context, categoryID string, tournamentID string) error
	WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) error) error
	CreateTournamentGroup(ctx context.Context, TournamentID string, position int) (string, error)
	CreatePot(ctx context.Context, TournamentID string, position int) (string, error)
	CreateDoubleEliminationEmpty(ctx context.Context) (string, error)
	TournamentGroupColl() *mongo.Collection
	PotColl() *mongo.Collection
	DoubleEliminationColl() *mongo.Collection
	DeleteByID(ctx context.Context, mc *mongo.Collection, ID string, name string) error
	UpdateTournamentRelations(ctx context.Context, tournamentID string, tournamentDTO *tournament_dto.UpdateTournamentOptionsDTOReq, add bool) error
	ConvertToObjectID(ID string) (*primitive.ObjectID, error)
	CreateTournamentRegistration(ctx context.Context, tournamentRegistrationDTO *tournament_dto.CreateTournamentRegistrationDTOReq) error
	CreateGuestUser(ctx context.Context, guestUserInfoDTO *tournament_dto.CreateGuestUserDTOReq) (string, error)
	CreateCompetitor(ctx context.Context, sport models.SPORT, competitorType models.COMPETITOR_TYPE, ID string) (string, error)
	CreateGuestCompetitor(ctx context.Context, guestCompetitorInfoDTO *tournament_dto.CreateGuestCompetitorDTOReq) (string, error)
	CreateMatch(ctx context.Context, match *tournament_dto.CreateMatchDTOReq) (string, error)
	CreateRound(ctx context.Context, round *tournament_dto.CreateRoundDTOReq) (string, error)
	CreateDoubleElimination(ctx context.Context, doubleEliminationDTO *tournament_dto.CreateDoubleEliminationDTOReq) (string, error)
	CreateSingle(ctx context.Context, singleInfoDTO *tournament_dto.CreateSingleDTOReq) (string, error)
	CreateDouble(ctx context.Context, doubleInfoDTO *tournament_dto.CreateDoubleDTOReq) (string, error)
	CreateTeam(ctx context.Context, teamInfoDTO *tournament_dto.CreateTeamDTOReq) (string, error)
	GetCompetitorsInTournament(
		ctx context.Context,
		tournamentID,
		categoryID,
		lastID string,
		limit int,
		gerAll bool,
	) ([]*tournament_dto.GetCompetitorsInTournamentCompetitorDTORes, error)
	VerifyCompetitorExists(ctx context.Context, competitorID string) error
	VerifyTournamentExists(ctx context.Context, tournamentID string) error
	CreateCompetitorStats(ctx context.Context, competitorID string) error
	UpdateCompetitorMatch(ctx context.Context, matchID string, competitorMatchDTO *tournament_dto.UpdateCompetitorMatchDTOReq) error
	VerifyMatchExists(ctx context.Context, matchID string) error
	CreateCompetitorMatch(ctx context.Context, competitorMatchDTO *tournament_dto.CreateCompetitorMatchDTOReq) error
	UpdateRoundTotalPrize(ctx context.Context, roundID string, totalPrize float64) error
	VerifyRoundExists(ctx context.Context, roundID string) error
	GetRoundWithMatches(ctx context.Context, roundID, categoryID string) (*tournament_dto.GetRoundWithMatchesDTORes, error)
	GetPositionsBracketMatch(ctx context.Context, roundID string, organizeType models.ORGANIZE_TYPE, organizeBracket models.ORGANIZE_BRACKET) ([]*tournament_dto.GetPositionsBracketMatchDTORes, error)
	UpdateMultipleCompetitorMatches(ctx context.Context, competitorsDTOs []*tournament_dto.UpdateCompetitorMatchDTOReq) error
	VerifyMatchesExist(ctx context.Context, matchIDs []string) error
	VerifyMultipleCompetitorsExists(ctx context.Context, competitorIDs []string) error
	SetWinnerInMatch(ctx context.Context, matchID, competitorID, result string) error
	VerifyMatchesInRoundExits(ctx context.Context, roundID string) (bool, error)
	FindMatchID(ctx context.Context, position int, roundID string) (string, error)
	AddMatchInTournament(ctx context.Context, tournamentID, matchID string) error
	AddMatchInCompetitorStats(ctx context.Context, competitorID, matchID string) error
	UpdateCompetitorStats(ctx context.Context, competitorID string, winner bool) error
	IncrementTotalCompetitorsInTournament(ctx context.Context, tournamentID string) error
	VerifyTournamentsCapacity(ctx context.Context, tournamentID string) (bool, error)
	UpdateRoundPoints(ctx context.Context, roundID string, points int) error
	UpdateTournamentInfo(ctx context.Context, tournamentID string, tournamentDTO *tournament_dto.UpdateTournamentInfoDTOReq) error
	AddTournamentWonInCompetitorStats(ctx context.Context, competitorID, tournamentID string) error
	GetTournamentCompetitorIDs(ctx context.Context, tournamentID string) ([]string, error)
	CreateCategoryRegistration(ctx context.Context, categoryRegistrationDAO *tournament_dto.CreateCategoryRegistrationDTOReq) error
	IncrementTotalParticipants(ctx context.Context, categorID string) error
	GetCategoryRegistrationSortedByPoints(ctx context.Context, categoryID string) ([]*tournament_dto.GetCategoryRegistrationSortedByPointsDTORes, error)
	UpdateCategoryRegistrationCurrentPosition(ctx context.Context, categoryID string, categoryRegistration []*tournament_dto.GetCategoryRegistrationSortedByPointsDTORes) error
	VerifyCompetitorsMatch(ctx context.Context, matchID, competitorID string) error
	UpdateTournamentFinishDate(ctx context.Context, tournamentID string) error
	VerifyMatchPosition(ctx context.Context, matchID string, position int) error
	GetRoundQuantityMatches(ctx context.Context, roundID string) (int, error)
	GetMatchPosition(ctx context.Context, matchID string) (int, error)
	GetRoundID(ctx context.Context, tournamentID string, round models.ROUND) (string, error)
	GetRoundsWithCompetitors(ctx context.Context, tournamentID string) ([]*tournament_dto.GetRoundWithCompetitorsDTORes, error)
	AddPointsInMultipleCategoryRegistration(ctx context.Context, categoryID string, competitorIDs []string, points int) error
	AddPrizeInMultipleCompetitorStats(ctx context.Context, competitorIDs []string, prize float64) error
	GetCompetitorsOutCategory(ctx context.Context, categoryID string, competitorIDs []string) ([]string, error)
	GetTournamentInfoToFinaliseIt(ctx context.Context, tournamentID string) (*tournament_dto.GetTournamentInfoToFinaliseItDTORes, error)
	VerifyTournamentsAlreadyFinished(ctx context.Context, tournamentID string) error
	VerifyMatchAndRoundCoincidence(ctx context.Context, matchID, roundID string, round models.ROUND) error
	VerifyMultipleCompetitorsExistsInTournament(ctx context.Context, tournamentID string, competitorIDs []string) error
	VerifyCompetitorExistsInTournament(ctx context.Context, tournamentID string, competitorID string) error
	AddCompetitorInGroup(ctx context.Context, groupID, competitorID string) error
	AddCompetitorsToTournamentGroups(ctx context.Context, tournamentID string, groupDTOs []*tournament_dto.AddCompetitorsToTournamentGroupsDTOReq) error
	AddMatchInTournamentGroup(ctx context.Context, groupID, tournamentID, matchID string) error
	VerifyMultipleGroupsInTournament(ctx context.Context, tournamentID string, groupIDs []string) error
	VerifyRoundInTournament(ctx context.Context, roundID, tournamentID string) error
	AddMultipleMatchesInTournamentGroup(ctx context.Context, groupID, tournamentID string, matchIDs []string) error
	AddMultipleMatchesInTournament(ctx context.Context, tournamentID string, matchIDs []string) error
	VerifyTournamentGroupInTournament(ctx context.Context, tournamentID string, groupIDs []string) error
	GetTournamentGroupMatches(ctx context.Context, groupID string) ([]string, []string, error)
	RemoveMultipleTournamentMatches(ctx context.Context, tournamentID string, matchesToRemoveIDs []string) error
	RemoveMultipleCompetitorStatsMatches(ctx context.Context, competitorIDs, matchesToRemove []string) error
	DeleteMultipleMatches(ctx context.Context, matchesToRemove []string) error
	DeleteMultipleCompetitorMatches(ctx context.Context, matchesToRemove []string) error
	VerifyTournamentPot(ctx context.Context, tournamentID, potID string) error
	AddCompetitorInPot(ctx context.Context, potID, competitorID string) error
	RemoveCompetitorOfPot(ctx context.Context, potID, competitorID string) error
	SetCompetitorsInPots(ctx context.Context, tournamentID string, potDTOs []*tournament_dto.SetPotCompetitorDTOReq) error
	VerifyMultipleTournamentPot(ctx context.Context, tournamentID string, potIDs []string) error
	AddPotInTournament(ctx context.Context, tournamentID, potID string) error
	RemovePotToTournament(ctx context.Context, tournamentID string, position int) error
	GetTournamentPotPositions(ctx context.Context, tournamentID string) ([]*tournament_dto.PotOrGroupPositionDTORes, error)
	UpdatePotPositions(ctx context.Context, potID string, position int) error
	DeletePotByPosition(ctx context.Context, position int, tournamentID string) error
	VerifyNumberPotsInTournament(ctx context.Context, tournamentID string, position int) error
	VerifyNumberGroupsInTournament(ctx context.Context, tournamentID string, position int) error
	AddGroupInTournament(ctx context.Context, tournamentID, groupID string) error
	UpdateGroupPositions(ctx context.Context, groupID string, position int) error
	RemoveGroupToTournament(ctx context.Context, tournamentID string, position int) error
	DeleteGroupByPosition(ctx context.Context, position int, tournamentID string) error
	GetTournamentGroupPositions(ctx context.Context, tournamentID string) ([]*tournament_dto.PotOrGroupPositionDTORes, error)
	GetTournamentGroupMatchesByPosition(ctx context.Context, position int, tournamentID string) ([]string, []string, error)
	GetDoubleElimRoundID(ctx context.Context, tournamentID string, round models.ROUND) (string, error)
	AddMatchInDoubleElim(ctx context.Context, doubleElimID, matchID string) error
	GetDoubleElimID(ctx context.Context, tournamentID string) (string, error)
	GetTournamentRoundNames(ctx context.Context, tournamentID string) ([]models.ROUND, error)
	GetAllDoubleElimRoundIDs(ctx context.Context, doubleEliminationID string) ([]string, error)
	GetDoubleElimInfoToFinaliseIt(ctx context.Context, doubleElimID string) (*tournament_dto.GetDoubleElimInfoToFinaliseItDTORes, error)
	GetDoubleElimCompetitorChampion(ctx context.Context, doubleElimOID string) (string, error)
	GetCompetitorChampion(ctx context.Context, tournamentOID string) (string, error)
	GetMultipleAvailabilitiesByCompetitor(ctx context.Context, competitorIDs []string) ([][]*models.GetDailyAvailabilityByIDDTORes, error)
	UpdateMultipleMatchesDate(ctx context.Context, matchDates []*tournament_dto.MatchDateDTOReq) error
	GetAvailabilityByTournamentID(ctx context.Context, tournamentID string) ([]*models.GetDailyAvailabilityByIDDTORes, error)
	GetTournamentAvailavility(ctx context.Context, tournamentID string) (*tournament_dto.TournamentAvailabilityDTO, error)
	CreateAvailability(ctx context.Context, userID, competitorID, tournamentID *string) error
	GetAllDatesMatchesFromTournament(ctx context.Context, tournamentID string) ([]time.Time, error)
	UpdateMatchDate(ctx context.Context, matchID string, date *time.Time) error
	CreateMatchChat(ctx context.Context, matchID string, competitorIDs []string, userID string) error
	VerifyCompetitorIDInCompetitorUser(ctx context.Context, competitorIDs []string) (bool, error)
	UpdateTournamentStartDate(ctx context.Context, tournamentID string) error
	GetUserTournaments(
		ctx context.Context,
		userID string,
		sport models.SPORT,
		limit int,
		lastID string,
	) (*tournament_dto.GetUserTournamentsDTORes, error)
	GetTournamentPrimaryInfo(ctx context.Context, tournamentID string) (*tournament_dto.GetTournamentPrimaryInfoDTORes, error)
	GetCompetitorsByNameInTournament(
		ctx context.Context,
		tournamentID, categoryID string,
		name string,
		limit int,
	) ([]*tournament_dto.GetCompetitorsInTournamentCompetitorDTORes, error)
	GetTournamentTotalCompetitors(ctx context.Context, tournamentID string) (int, error)
	GetCategoryIDOfTournament(ctx context.Context, tournamentID string) (string, error)
	GetCompetitorsFollowed(ctx context.Context, userID string, name string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) ([]*tournament_dto.GetCompetitorFollowedDTORes, error)
	VerifyUserExists(ctx context.Context, userID string) error
	CreateCompetitorUser(ctx context.Context, userID, competitorID string) error
	CreateAvailabilityForCompetitor(ctx context.Context, competitorID string, dailyAvailability []*models.GetDailyAvailabilityByIDDTORes) error
	GetAvailabilityDailySlice(ctx context.Context, userID, competitorID string) ([]*models.GetDailyAvailabilityByIDDTORes, error)
	GetTournamentFilters(ctx context.Context, tournamentID string) (*tournament_dto.GetTournamentFiltersDTORes, error)
	GetTournamentsInOrganizer(
		ctx context.Context,
		organizerID string,
		sport models.SPORT,
		limit int,
		lastID string,
	) (*tournament_dto.GetUserTournamentsDTORes, error)
	AddTournamentInOrganizer(ctx context.Context, organizerOID, tournamentOID string) error
	DeleteTournamentRegistration(ctx context.Context, tournamentRegistrationID string) error
	DecrementTotalCompetitorsInTournament(ctx context.Context, tournamentID string) error
	GetTournamentRegistrationByCompetitorAndTournamentID(ctx context.Context, tournamentID, competitorID string) (string, error)
	GetTournamentMatchesByID(ctx context.Context, tournamentID string) ([]string, error)
	GetCompetitorIDsFromMatches(ctx context.Context, matchIDs []string) ([]string, error)
	GetCompetitorIDByMatchAndPosition(ctx context.Context, matchID string, position int) (string, error)
	GetRoundGroups(ctx context.Context, roundID, categoryID string) (*tournament_dto.GetRoundGroupsDTORes, error)
	GetDailyAvailabilityTournamentID(ctx context.Context, tournamentID string, day string) (*models.GetDailyAvailabilityByIDDTORes, string, error)
	UpdateAvailability(ctx context.Context, availabilityID string, availabilityInfoDAO *models.UpdateDailyAvailabilityDTOReq) error
	GetTournamentGroupsIDs(ctx context.Context, tournamentID string) ([]string, error)
}
