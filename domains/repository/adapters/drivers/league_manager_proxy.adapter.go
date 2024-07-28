package adapters

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/league/dao"
	"github.com/DBrange/didis-comp-bk/domains/repository/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LeagueManagerProxyAdapter struct {
	repository *repository.Repository
}

func NewLeagueManagerProxyAdapter(repository *repository.Repository) *LeagueManagerProxyAdapter {
	return &LeagueManagerProxyAdapter{
		repository: repository,
	}
}

func (a *LeagueManagerProxyAdapter) VerifyOrganizerExists(ctx context.Context, organizerID string) error {
	return a.repository.VerifyOrganizerExists(ctx, organizerID)
}

func (a *LeagueManagerProxyAdapter) CreateLeague(ctx context.Context, organizerID string, leagueInfoDAO *dao.CreateLeagueDAOReq) error {
	return a.repository.CreateLeague(ctx, organizerID, leagueInfoDAO)
}

func (a *LeagueManagerProxyAdapter) ConvertToObjectID(ID string) (*primitive.ObjectID, error) {
	return a.repository.ConvertToObjectID(ID)
}

func (a *LeagueManagerProxyAdapter) AddTournamentInLeague(ctx context.Context, leagueID string, tournamentID string) error {
	return a.repository.AddTournamentInLeague(ctx, leagueID, tournamentID)
}

func (a *LeagueManagerProxyAdapter) AddLeagueInTournament(ctx context.Context, tournamentID string, leagueID string) error {
	return a.repository.AddLeagueInTournament(ctx, tournamentID, leagueID)
}
