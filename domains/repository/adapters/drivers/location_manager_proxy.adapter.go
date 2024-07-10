package adapters

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	"github.com/DBrange/didis-comp-bk/domains/repository/repository"
)

type LocationManagerProxyAdapter struct {
	repository *repository.Repository
}

func NewLocationMangerProxyAdapter(repository *repository.Repository) *LocationManagerProxyAdapter {
	return &LocationManagerProxyAdapter{
		repository: repository,
	}
}

func (a *LocationManagerProxyAdapter) CreateLocation(ctx context.Context, location *dao.CreateLocationDAOReq) (string, error) {
	id, err := a.repository.CreateLocation(ctx, location)
	if err != nil {
		return "", err
	}

	return id, nil
}
