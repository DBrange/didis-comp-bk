package adapters

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	option_models "github.com/DBrange/didis-comp-bk/cmd/api/models/options/tournament"
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

func (a *TournamentManagerProxyAdapter) CreateTournamentGroup(ctx context.Context, tournamentID string) (string, error) {
	return a.repository.CreateTournamentGroup(ctx, tournamentID)
}

func (a *TournamentManagerProxyAdapter) CreatePot(ctx context.Context, tournamentID string) (string, error) {
	return a.repository.CreatePot(ctx, tournamentID)
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

func (a *TournamentManagerProxyAdapter) VerifyTournamentsExists(ctx context.Context, tournamentOID *primitive.ObjectID) error {
	return a.repository.VerifyTournamentsExists(ctx, tournamentOID)
}

func (a *TournamentManagerProxyAdapter) CreateCompetitorStats(ctx context.Context, competitorOID *primitive.ObjectID) error {
	return a.repository.CreateCompetitorStats(ctx, competitorOID)

}
