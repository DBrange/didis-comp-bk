package adapters

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	"github.com/DBrange/didis-comp-bk/domains/repository/repository"
)

type LocationManagerProxyAdapter struct {
	repository *repository.Repository
}

func NewLocationManagerProxyAdapter(repository *repository.Repository) *LocationManagerProxyAdapter {
	return &LocationManagerProxyAdapter{
		repository: repository,
	}
}

func (a *LocationManagerProxyAdapter) CreateLocation(ctx context.Context, location *dao.CreateLocationDAOReq) (string, error) {
	return a.repository.CreateLocation(ctx, location)
}

func (a *LocationManagerProxyAdapter) GetLocationByID(ctx context.Context, id string) (*dao.GetLocationByIDDAORes, error) {
	return a.repository.GetLocationByID(ctx, id)
}

func (a *LocationManagerProxyAdapter) UpdateLocation(ctx context.Context, locationID string, newLocationInfoDAO *dao.UpdateLocationDAOReq) error {
	return a.repository.UpdateLocation(ctx, locationID, newLocationInfoDAO)
}

func (a *LocationManagerProxyAdapter) DeleteLocation(ctx context.Context, locationID string) error {
	return a.repository.DeleteLocation(ctx, locationID)
}
