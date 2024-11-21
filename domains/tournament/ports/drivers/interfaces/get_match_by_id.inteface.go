package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

type GetMatchByID interface {
	GetMatchByID(ctx context.Context, matchID string) (*dto.GetMatchDTORes, error)
}
