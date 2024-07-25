package ports

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/league/models/dto"
)

type ForQueryingLeague interface {
	OrganizeLeague(ctx context.Context, organizerID string, leagueInfoDTO *dto.OrganizeLeagueDTOReq) error
	AddTournamentInLeague(ctx context.Context, leagueID string, tournamentID string) error
}
