package adapters

import (
	"context"

	location_dto "github.com/DBrange/didis-comp-bk/domains/location/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/location/services"
)

type LocationProxyAdapter struct {
	ctx             context.Context
	locationService *services.LocationService
}

func NewLocationProxyAdapter(ctx context.Context, locationService *services.LocationService) *LocationProxyAdapter {
	return &LocationProxyAdapter{
		ctx:             ctx,
		locationService: locationService,
	}
}

func (a *LocationProxyAdapter) CreateLocation(locationDTO *location_dto.CreateLocationDTOReq) (string, error) {
	return a.locationService.CreateLocation(a.ctx, locationDTO)
}
