package ports

import "context"

type ForQueryingLeague interface {
	OrganizeLeague(ctx context.Context, leagueInfoDTO any) error
}