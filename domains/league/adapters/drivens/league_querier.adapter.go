package adapters

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/league/adapters/mappers"
	"github.com/DBrange/didis-comp-bk/domains/league/models/dto"
	ports "github.com/DBrange/didis-comp-bk/domains/repository/ports/drivers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LeagueQueryerAdapter struct {
	adapter ports.ForManagingLeague
}

func NewLeagueQueryerAdapter(adapter ports.ForManagingLeague) *LeagueQueryerAdapter {
	return &LeagueQueryerAdapter{
		adapter: adapter,
	}
}

func (a *LeagueQueryerAdapter) VerifyOrganizerExists(ctx context.Context, organizerID string) error {
	return a.adapter.VerifyOrganizerExists(ctx, organizerID)
}

func (a *LeagueQueryerAdapter) ConvertToObjectID(ID string) (*primitive.ObjectID, error) {
	return a.adapter.ConvertToObjectID(ID)
}

func (a *LeagueQueryerAdapter) CreateLeague(ctx context.Context, organizerID string, leagueDTO *dto.CreateLeagueDTOReq) error {
	leagueDAO := mappers.CreateLeagueDTOtoDAO(leagueDTO)

	return a.adapter.CreateLeague(ctx, organizerID, leagueDAO)
}

func (a *LeagueQueryerAdapter) AddTournamentInLeague(ctx context.Context, leagueID string, tournamentID string) error {
	return a.adapter.AddTournamentInLeague(ctx, leagueID, tournamentID)
}

func (a *LeagueQueryerAdapter) AddLeagueInTournament(ctx context.Context, tournamentID string, leagueID string) error {
	return a.adapter.AddTournamentInLeague(ctx, tournamentID, leagueID)
}
