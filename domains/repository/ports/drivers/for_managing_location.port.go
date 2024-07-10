package ports

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
)

type ForManagingLocation interface {
	CreateLocation(ctx context.Context, locationDAO *dao.CreateLocationDAOReq) (string, error)
}
