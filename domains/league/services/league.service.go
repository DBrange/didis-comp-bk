package services

import ports "github.com/DBrange/didis-comp-bk/domains/league/ports/drivens"

type LeagueService struct {
	leagueQueryer ports.ForQueryingLeague
}

func NewLeagueService(leagueQueryer ports.ForQueryingLeague) *LeagueService {
	return &LeagueService{
		leagueQueryer: leagueQueryer,
	}
}
