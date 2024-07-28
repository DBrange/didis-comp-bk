package interfaces

import (
	"context"

	models "github.com/DBrange/didis-comp-bk/cmd/api/models/options/tournament"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

type OrganizeTournament interface {
	OrganizeTournament(ctx context.Context, tournamentDTO *dto.OrganizeTournamentDTOReq, options *models.OrganizeTournamentOptions) error
}