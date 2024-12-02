package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

type ModifyBracketMatch interface {
	ModifyBracketMatch(ctx context.Context, tournamentID string, competitors []*dto.UpdateCompetitorMatchDTOReq) error
}
