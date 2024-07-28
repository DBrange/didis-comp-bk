package drivens

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	option_models "github.com/DBrange/didis-comp-bk/cmd/api/models/options/tournament"
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
	leagueID *string,
	organizerID string,
) (string, error) {
	tournamentDAO := mappers.CreateTournamentDTOtoDAO(tournamentDTO)

	return a.adapter.CreateTournament(ctx, tournamentDAO, locationID, options, leagueID, organizerID)
}

func (a *TournamentQueryerAdapter) VerifyLeagueExists(ctx context.Context, leagueID string) error {
	return a.adapter.VerifyLeagueExists(ctx, leagueID)
}

func (a *TournamentQueryerAdapter) AddTournamentInLeague(ctx context.Context, leagueID string, tournamentID string) error {
	return a.adapter.AddTournamentInLeague(ctx, leagueID, tournamentID)
}

func (a *TournamentQueryerAdapter) WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) error) error {
	return a.adapter.WithTransaction(ctx, fn)
}

func (a *TournamentQueryerAdapter) CreateTournamentGroup(ctx context.Context, TournamentID string) (string, error) {
	return a.adapter.CreateTournamentGroup(ctx, TournamentID)
}

func (a *TournamentQueryerAdapter) CreatePot(ctx context.Context, TournamentID string) (string, error) {
	return a.adapter.CreatePot(ctx, TournamentID)
}

func (a *TournamentQueryerAdapter) CreateDoubleElimination(ctx context.Context) (string, error) {
	return a.adapter.CreateDoubleElimination(ctx)
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

func (a *TournamentQueryerAdapter) UpdateTournamentOptions(
	ctx context.Context,
	tournamentID string,
	tournamentDTO *tournament_dto.UpdateTournamentOptionsDTOReq,
	add bool,
) error {
	tournamentDAO, err := mappers.UpdateTournamentOptionsDTOtoDAO(tournamentDTO, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.adapter.UpdateTournamentOptions(ctx, tournamentID, tournamentDAO, add)
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

func (a *TournamentQueryerAdapter) CreateCompetitorType(ctx context.Context, competitorType models.COMPETITOR_TYPE) (string, error) {
	competitorTypeOID, err := a.adapter.CreateCompetitorType(ctx, competitorType)
	if err != nil {
		return "", err
	}
	return competitorTypeOID.Hex(), nil
}

func (a *TournamentQueryerAdapter) CreateGuestCompetitor(ctx context.Context, guestCompetitorDTO *tournament_dto.CreateGuestCompetitorDTOReq) (string, error) {
	guestCompetitorDAO, err := mappers.CreateGuestCompetitorDTOtoDAO(guestCompetitorDTO, a.ConvertToObjectID)
	if err != nil {
		return "", err
	}

	return a.adapter.CreateGuestCompetitor(ctx, guestCompetitorDAO)
}
