package services

import (
	ports "github.com/DBrange/didis-comp-bk/domains/profile/ports/drivens"
)

type ProfileService struct {
	profileQueryer ports.ForQueryingProfile
}

func NewProfileService(profileQueryer ports.ForQueryingProfile) *ProfileService {
	return &ProfileService{
		profileQueryer: profileQueryer,
	}
}
