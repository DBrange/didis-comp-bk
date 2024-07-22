package ports

import "context"

type ForManagingLeague interface {
	OrganizeLeague(ctx context.Context, leagueInfoDAO any) error
}
