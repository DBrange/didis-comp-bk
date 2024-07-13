package ports

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
)

type ForManagingLocation interface {
	CreateLocation(ctx context.Context, locationDAO *dao.CreateLocationDAOReq) (string, error)
	GetLocationByID(ctx context.Context, id string) (*dao.GetLocationByIDDAORes, error)
	UpdateLocation(ctx context.Context, locationID string, newLocationInfoDAO *dao.UpdateLocationDAOReq) error
	DeleteLocation(ctx context.Context, locationID string) error
}
