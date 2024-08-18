package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

type EndMatch interface {
	EndMatch(ctx context.Context, match *dto.EndMatchDTOReq) error
}
