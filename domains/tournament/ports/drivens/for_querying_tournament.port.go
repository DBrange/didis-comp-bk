package ports

import (
	"context"

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
	CreateTournamentGroup(ctx context.Context, TournamentID string) (string, error)
	CreatePot(ctx context.Context, TournamentID string) (string, error)
	CreateDoubleEliminationEmpty(ctx context.Context) (string, error)
	TournamentGroupColl() *mongo.Collection
	PotColl() *mongo.Collection
	DoubleEliminationColl() *mongo.Collection
	DeleteByID(ctx context.Context, mc *mongo.Collection, ID string, name string) error
	UpdateTournamentRelations(ctx context.Context, tournamentOID string, tournamentDTO *tournament_dto.UpdateTournamentOptionsDTOReq, add bool) error
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
	) ([]*tournament_dto.GetCompetitorsInTournamentDTORes, error)
	VerifyCompetitorExists(ctx context.Context, competitorID string) error
	VerifyTournamentsExists(ctx context.Context, tournamentOID string) error
	CreateCompetitorStats(ctx context.Context, competitorID string) error
}
