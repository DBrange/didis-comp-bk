package ports

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	optional_models "github.com/DBrange/didis-comp-bk/cmd/api/models/options/tournament"
	double_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/double/dao"
	double_elimination_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/double_elimination/dao"
	guest_user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/guest_user/dao"
	guest_competitor_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/guest_competitor/dao"
	tournament_registration_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/tournament_registration/dao"
	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	match_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/match/dao"
	round_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/round/dao"
	single_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/single/dao"
	team_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/team/dao"
	tournament_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/tournament/dao"
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
	CreateTournamentGroup(ctx context.Context, TournamentID string) (string, error)
	CreatePot(ctx context.Context, TournamentID string) (string, error)
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
	ListCompetitorsInTournament(ctx context.Context, tournamentID, categoryID, lastID string, limit int) ([]*tournament_registration_dao.GetCompetitorsInTournamentDAORes, error)
	VerifyCompetitorExists(ctx context.Context, competitorOID *primitive.ObjectID) error
	VerifyTournamentsExists(ctx context.Context, tournamentOID *primitive.ObjectID) error
	CreateCompetitorStats(ctx context.Context, competitorOID *primitive.ObjectID) error
}
