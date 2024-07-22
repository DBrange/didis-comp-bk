package services

import ports "github.com/DBrange/didis-comp-bk/domains/tournament/ports/drivens"

type TournamentService struct {
	tournamentQueryer ports.ForQueryingTournament
}

func NewTournamentService(tournamentQueryer ports.ForQueryingTournament) *TournamentService {
	return &TournamentService{
		tournamentQueryer: tournamentQueryer,
	}
}
