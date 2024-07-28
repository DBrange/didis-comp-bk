package ports

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/league/dao"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ForManagingLeague interface {
	VerifyOrganizerExists(ctx context.Context, organizerID string) error
	CreateLeague(ctx context.Context, organizerID string, leagueInfoDAO *dao.CreateLeagueDAOReq) error
	ConvertToObjectID(ID string) (*primitive.ObjectID, error)
	AddTournamentInLeague(ctx context.Context, leagueID string, tournamentID string) error
	AddLeagueInTournament(ctx context.Context, tournamentID string, leagueID string) error
}
