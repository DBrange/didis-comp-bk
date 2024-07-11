package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/location/models/dto"
)

type CreateLocation interface {
	CreateLocation(ctx context.Context, location *dto.CreateLocationDTOReq) (string, error)
}
