package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

type OrganizeTournament interface {
	OrganizeTournament(ctx context.Context, tournamentDTO *dto.OrganizeTournamentDTOReq) error
}