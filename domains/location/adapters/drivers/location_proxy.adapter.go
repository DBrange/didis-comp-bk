package adapters

import (
	"context"

	location_dto "github.com/DBrange/didis-comp-bk/domains/location/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/location/services"
)

type LocationProxyAdapter struct {
	locationService *services.LocationService
}

func NewLocationProxyAdapter(locationService *services.LocationService) *LocationProxyAdapter {
	return &LocationProxyAdapter{
		locationService: locationService,
	}
}

func (a *LocationProxyAdapter) CreateLocation(ctx context.Context, locationDTO *location_dto.CreateLocationDTOReq) (string, error) {
	return a.locationService.CreateLocation(ctx, locationDTO)
}

func (a *LocationProxyAdapter) GetLocationByID(ctx context.Context, id string) (*location_dto.GetLocationByIDDTORes, error) {
	return a.locationService.GetLocationByID(ctx, id)
}

func (a *LocationProxyAdapter) UpdateLocation(ctx context.Context,locationID string, location *location_dto.UpdateLocationDTOReq) error {
	return a.locationService.UpdateLocation(ctx,locationID, location)
}
