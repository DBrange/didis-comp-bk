package services

import (
	ports "github.com/DBrange/didis-comp-bk/domains/profile/ports/drivens"
)

type ProfileService struct {
	profileQuerier ports.ForQueryingProfile
}

func NewProfileService(profileQuerier ports.ForQueryingProfile) *ProfileService {
	return &ProfileService{
		profileQuerier: profileQuerier,
	}
}
