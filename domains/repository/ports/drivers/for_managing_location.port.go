package ports

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	"go.mongodb.org/mongo-driver/bson"
)

type ForManagingLocation interface {
	CreateLocation(ctx context.Context, locationDAO *dao.CreateLocationDAOReq) (string, error)
	GetLocationByID(ctx context.Context, id string) (*dao.GetLocationByIDDAORes, error)
	UpdateLocation(ctx context.Context, filter bson.M, update bson.M) error
}
