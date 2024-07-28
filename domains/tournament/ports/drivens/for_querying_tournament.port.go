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
	UpdateTournamentOptions(ctx context.Context, tournamentID string, tournamentDTO *tournament_dto.UpdateTournamentOptionsDTOReq, add bool) error
	ConvertToObjectID(ID string) (*primitive.ObjectID, error)
	CreateTournamentRegistration(ctx context.Context, tournamentRegistrationDTO *tournament_dto.CreateTournamentRegistrationDTOReq) error
	CreateGuestUser(ctx context.Context, guestUserInfoDTO *tournament_dto.CreateGuestUserDTOReq) (string, error)
	CreateCompetitor(ctx context.Context, sport models.SPORT, competitorType models.COMPETITOR_TYPE, ID string) (string, error)
	CreateCompetitorType(ctx context.Context, competitorType models.COMPETITOR_TYPE) (string, error)
	CreateGuestCompetitor(ctx context.Context, guestCompetitorInfoDTO *tournament_dto.CreateGuestCompetitorDTOReq) (string, error)
}
