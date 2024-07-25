package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/league/models/dto"
)

type OrganizeLeague interface {
	OrganizeLeague(ctx context.Context, organizerID string, leagueDTO *dto.OrganizeLeagueDTOReq) error
}
