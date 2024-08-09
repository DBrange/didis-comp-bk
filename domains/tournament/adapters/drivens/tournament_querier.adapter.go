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

func (a *TournamentQueryerAdapter) CreateTournamentGroup(ctx context.Context, TournamentID string) (string, error) {
	return a.adapter.CreateTournamentGroup(ctx, TournamentID)
}

func (a *TournamentQueryerAdapter) CreatePot(ctx context.Context, TournamentID string) (string, error) {
	return a.adapter.CreatePot(ctx, TournamentID)
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
