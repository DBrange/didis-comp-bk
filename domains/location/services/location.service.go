package services

import ports "github.com/DBrange/didis-comp-bk/domains/location/ports/drivens"

type LocationService struct {
	locationQuerier ports.ForQueryingLocation
}

func NewLocationService(locationQuerier ports.ForQueryingLocation) *LocationService {
	return &LocationService{
		locationQuerier: locationQuerier,
	}
}
