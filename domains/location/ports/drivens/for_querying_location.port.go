package ports

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/location/models/dto"
)

type ForQueryingLocation interface {
	CreateLocation(ctx context.Context, userDTO *dto.CreateLocationDTOReq) (string, error)
	GetLocationByID(ctx context.Context, id string) (*dto.GetLocationByIDDTORes, error)
	UpdateLocation(ctx context.Context,locationID string, location *dto.UpdateLocationDTOReq) error
}
