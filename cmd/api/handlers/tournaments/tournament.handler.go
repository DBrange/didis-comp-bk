package handlers

import ports "github.com/DBrange/didis-comp-bk/domains/tournament/ports/drivers"

type Handler struct {
	tournament ports.ForTournament
}

func NewHandlerTournament(tournament ports.ForTournament) *Handler {
	return &Handler{
		tournament: tournament,
	}
}
