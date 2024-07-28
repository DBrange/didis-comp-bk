package adapters

import "github.com/DBrange/didis-comp-bk/domains/repository/repository"

type ControlPlaneManagerProxyAdapter struct {
	repository *repository.Repository
}

func NewControlPlaneManagerProxyAdapter(repository *repository.Repository) *ControlPlaneManagerProxyAdapter {
	return &ControlPlaneManagerProxyAdapter{
		repository: repository,
	}
}
