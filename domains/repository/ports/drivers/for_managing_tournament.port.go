package ports

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	optional_models "github.com/DBrange/didis-comp-bk/cmd/api/models/options/tournament"
	guest_user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/guest_user/dao"
	guest_competitor_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/guest_competitor/dao"
	tournament_registration_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/tournament_registration/dao"
	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
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
		leagueID *string,
		organizerID string,
	) (string, error)
	VerifyLeagueExists(ctx context.Context, leagueID string) error
	AddTournamentInLeague(ctx context.Context, leagueID string, tournamentID string) error
	WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) error) error
	CreateTournamentGroup(ctx context.Context, TournamentID string) (string, error)
	CreatePot(ctx context.Context, TournamentID string) (string, error)
	CreateDoubleElimination(ctx context.Context) (string, error)
	TournamentGroupColl() *mongo.Collection
	PotColl() *mongo.Collection
	DoubleEliminationColl() *mongo.Collection
	DeleteByID(ctx context.Context, mc *mongo.Collection, ID string, name string) error
	UpdateTournamentOptions(ctx context.Context, tournamentID string, tournamentDTO *tournament_dao.UpdateTournamentOptionsDAOReq, add bool) error
	ConvertToObjectID(ID string) (*primitive.ObjectID, error)
	CreateTournamentRegistration(ctx context.Context, tournamentRegistrationInfoDTO *tournament_registration_dao.CreateTournamentRegistrationDAOReq) error
	CreateGuestUser(ctx context.Context, guestUserInfoDAO *guest_user_dao.CreateGuestUserDAOReq) (string, error)
	CreateCompetitor(ctx context.Context, sport models.SPORT, competitorType models.COMPETITOR_TYPE, OID *primitive.ObjectID) (string, error)
	CreateCompetitorType(ctx context.Context, competitorType models.COMPETITOR_TYPE) (*primitive.ObjectID, error)
	CreateGuestCompetitor(ctx context.Context, guestCompetitorInfoDAO *guest_competitor_dao.CreateGuestCompetitorDAOReq) (string, error)
}
