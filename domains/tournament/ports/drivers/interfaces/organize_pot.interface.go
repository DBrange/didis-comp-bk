package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

type OrganizePots interface {
	OrganizePots(ctx context.Context, tournamentID string, potDTOs []*dto.SetPotCompetitorDTOReq) error
}
