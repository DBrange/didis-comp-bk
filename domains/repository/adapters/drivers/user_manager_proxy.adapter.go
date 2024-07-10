package adapters

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
	"github.com/DBrange/didis-comp-bk/domains/repository/repository"
)

type UserMangerProxyAdapter struct {
	repository *repository.Repository
}

func NewUserMangerProxyAdapter(repository *repository.Repository) *UserMangerProxyAdapter {
	return &UserMangerProxyAdapter{
		repository: repository,
	}
}

func (a *UserMangerProxyAdapter) CreateUser(ctx context.Context, user *dao.CreateUserDAO) error {
	err := a.repository.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (a *UserMangerProxyAdapter) GetUserByID(ctx context.Context, id string) (*dao.GetUserByIDDAO, error) {
	return a.repository.GetUserByID(ctx, id)
}
