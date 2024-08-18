package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

type GetRoundWithMatches interface {
	GetRoundWithMatches(ctx context.Context, roundID, categoryID string) (*dto.GetRoundWithMatchesDTORes, error)
}
