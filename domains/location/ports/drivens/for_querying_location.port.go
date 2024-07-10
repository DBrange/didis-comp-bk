package ports

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/location/models/dto"
)

type ForQueryingLocation interface {
	CreateLocation(ctx context.Context, userDTO *dto.CreateLocationDTOReq) (string, error)
}