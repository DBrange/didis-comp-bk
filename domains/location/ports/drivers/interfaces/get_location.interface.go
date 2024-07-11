package interfaces

import (
	"context"

	location_dto "github.com/DBrange/didis-comp-bk/domains/location/models/dto"
)

type GetLocation interface {
	GetLocationByID(ctx context.Context,id string) (*location_dto.GetLocationByIDDTORes, error)
}
