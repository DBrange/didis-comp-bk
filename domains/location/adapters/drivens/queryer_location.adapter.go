package drivens

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/location/adapters/mappers"
	location_dto "github.com/DBrange/didis-comp-bk/domains/location/models/dto"
	ports "github.com/DBrange/didis-comp-bk/domains/repository/ports/drivers"
)

type LocationQueryerAdapter struct {
	drivers ports.ForManagingLocation
}

func NewLocationQueryerAdapter(drivers ports.ForManagingLocation) *LocationQueryerAdapter {
	return &LocationQueryerAdapter{
		drivers: drivers,
	}
}

func (a *LocationQueryerAdapter) CreateLocation(ctx context.Context, locationDTO *location_dto.CreateLocationDTOReq) (string, error) {
	locationToDAO := mappers.CreateLocationDTOtoDAO(locationDTO)

	id, err := a.drivers.CreateLocation(ctx, locationToDAO)
	if err != nil {
		return "", err
	}

	return id, nil
}
