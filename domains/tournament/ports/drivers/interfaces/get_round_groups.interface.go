package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

type GetRoundGroups interface {
	GetRoundGroups(ctx context.Context, roundID, categoryID string) (*dto.GetRoundGroupsDTORes, error)
}
