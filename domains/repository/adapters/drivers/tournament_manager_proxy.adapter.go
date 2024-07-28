package adapters

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	option_models "github.com/DBrange/didis-comp-bk/cmd/api/models/options/tournament"
	guest_user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/guest_user/dao"
	guest_competitor_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/guest_competitor/dao"
	tournament_registration_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/tournament_registration/dao"
	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	tournament_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/tournament/dao"
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
	leagueID *string,
	organizerID string,
) (string, error) {
	return a.repository.CreateTournament(ctx, tournamentInfoDAO, locationID, options, leagueID, organizerID)
}

func (a *TournamentManagerProxyAdapter) VerifyLeagueExists(ctx context.Context, leagueID string) error {
	return a.repository.VerifyLeagueExists(ctx, leagueID)
}

func (a *TournamentManagerProxyAdapter) AddTournamentInLeague(ctx context.Context, leagueID string, tournamentID string) error {
	return a.repository.AddTournamentInLeague(ctx, leagueID, tournamentID)
}

func (a *TournamentManagerProxyAdapter) WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) error) error {
	return a.repository.WithTransaction(ctx, fn)
}

func (a *TournamentManagerProxyAdapter) CreateTournamentGroup(ctx context.Context, tournamentID string) (string, error) {
	return a.repository.CreateTournamentGroup(ctx, tournamentID)
}

func (a *TournamentManagerProxyAdapter) CreatePot(ctx context.Context, tournamentID string) (string, error) {
	return a.repository.CreatePot(ctx, tournamentID)
}

func (a *TournamentManagerProxyAdapter) CreateDoubleElimination(ctx context.Context) (string, error) {
	return a.repository.CreateDoubleElimination(ctx)
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

func (a *TournamentManagerProxyAdapter) UpdateTournamentOptions(ctx context.Context, tournamentID string, tournamentDAO *tournament_dao.UpdateTournamentOptionsDAOReq, add bool) error {
	return a.repository.UpdateTournamentOptions(ctx, tournamentID, tournamentDAO, add)
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

func (a *TournamentManagerProxyAdapter) CreateCompetitorType(ctx context.Context, competitorType models.COMPETITOR_TYPE) (*primitive.ObjectID, error) {
	return a.repository.CreateCompetitorType(ctx, competitorType)
}

func (a *TournamentManagerProxyAdapter) CreateGuestCompetitor(ctx context.Context, guestCompetitorInfoDAO *guest_competitor_dao.CreateGuestCompetitorDAOReq) (string, error) {
	return a.repository.CreateGuestCompetitor(ctx, guestCompetitorInfoDAO)
}
