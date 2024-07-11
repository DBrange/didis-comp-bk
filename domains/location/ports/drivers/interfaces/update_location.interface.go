package interfaces

import (
	"context"

	location_dto "github.com/DBrange/didis-comp-bk/domains/location/models/dto"
)

type UpdateLocation interface {
	UpdateLocation(ctx context.Context,locationID string, location *location_dto.UpdateLocationDTOReq) error
}
