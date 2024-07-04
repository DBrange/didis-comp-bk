package handlers

import ports "didis-comp-bk/internal/user/ports/drivers"

type Handler struct {
	user ports.ForUser
}

func NewHandlerUser(user ports.ForUser) *Handler {
	return &Handler{
		user: user,
	}
}
