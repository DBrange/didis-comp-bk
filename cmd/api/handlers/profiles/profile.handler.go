package handlers

import (
	profile_ports "github.com/DBrange/didis-comp-bk/domains/profile/ports/drivers"
)

type Handler struct {
	profile profile_ports.ForProfile
}

func NewHandlerProfile(profile profile_ports.ForProfile) *Handler {
	return &Handler{
		profile: profile,
	}
}
