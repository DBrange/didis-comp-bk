package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

type OrganizeBracket interface {
	OrganizeBracket(ctx context.Context,tournamentID string, competitorMatchDTO []*dto.UpdateCompetitorMatchDTOReq) error
}
