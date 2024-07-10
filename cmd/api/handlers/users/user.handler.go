package handlers

import (
	location_ports "github.com/DBrange/didis-comp-bk/domains/location/ports/drivers"
	user_ports "github.com/DBrange/didis-comp-bk/domains/user/ports/drivers"
)

type Handler struct {
	user     user_ports.ForUser
	location location_ports.ForLocation
}

func NewHandlerUser(user user_ports.ForUser, location location_ports.ForLocation) *Handler {
	return &Handler{
		user:     user,
		location: location,
	}
}
