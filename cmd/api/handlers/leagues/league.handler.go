package handlers

import ports "github.com/DBrange/didis-comp-bk/domains/league/ports/drivers"

type Handler struct {
	league ports.ForLeague
}

func NewHandlerLeague(league ports.ForLeague) *Handler {
	return &Handler{
		league: league,
	}
}
