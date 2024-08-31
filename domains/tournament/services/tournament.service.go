package services

import ports "github.com/DBrange/didis-comp-bk/domains/tournament/ports/drivens"

type TournamentService struct {
	tournamentQuerier ports.ForQueryingTournament
}

func NewTournamentService(tournamentQuerier ports.ForQueryingTournament) *TournamentService {
	return &TournamentService{
		tournamentQuerier: tournamentQuerier,
	}
}
