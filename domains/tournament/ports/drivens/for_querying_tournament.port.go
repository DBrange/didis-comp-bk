package ports

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

type ForQueryingTournament interface {
	OrganizeTournament(ctx context.Context, tournamentInfoDTO *dto.OrganizeTournamentDTOReq) error
}
