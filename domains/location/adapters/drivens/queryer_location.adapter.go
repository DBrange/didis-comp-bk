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

func (a *LocationQueryerAdapter) GetLocationByID(ctx context.Context, id string) (*location_dto.GetLocationByIDDTORes, error) {
	locationDTO, err := a.drivers.GetLocationByID(ctx, id)
	if err != nil {
		return nil, err
	}
	mappedLocation := mappers.GetLocationByIDDAOtoDTO(locationDTO)

	return mappedLocation, nil
}

func (a *LocationQueryerAdapter) UpdateLocation(ctx context.Context, locationID string, newLocationInfoDTO *location_dto.UpdateLocationDTOReq) error {
	newLocationInfoDAO := mappers.UpdateLocationDTOtoDAO(newLocationInfoDTO)

	return a.drivers.UpdateLocation(ctx, locationID, newLocationInfoDAO)
}

func (a *LocationQueryerAdapter) DeleteLocation(ctx context.Context, locationID string) error {
	return a.drivers.DeleteLocation(ctx, locationID)
}
