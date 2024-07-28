package ports

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/league/models/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ForQueryingLeague interface {
	VerifyOrganizerExists(ctx context.Context, organizerID string) error
	CreateLeague(ctx context.Context, organizerID string, leagueDtO *dto.CreateLeagueDTOReq) error
	ConvertToObjectID(ID string) (*primitive.ObjectID, error)
	AddTournamentInLeague(ctx context.Context, leagueID string, tournamentID string) error
	AddLeagueInTournament(ctx context.Context, tournamentID string, leagueID string) error
}
