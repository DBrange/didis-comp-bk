package interfaces

import "context"

type OrganizeLeague interface {
	OrganizeLeague(ctx context.Context, leagueDTO any) error
}
