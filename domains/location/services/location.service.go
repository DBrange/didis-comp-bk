package services

import ports "github.com/DBrange/didis-comp-bk/domains/location/ports/drivens"

type LocationService struct {
	locationQueryer ports.ForQueryingLocation
}

func NewLocationService(locationQueryer ports.ForQueryingLocation) *LocationService {
	return &LocationService{
		locationQueryer: locationQueryer,
	}
}
