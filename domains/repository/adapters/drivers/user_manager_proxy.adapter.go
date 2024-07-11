package adapters

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
	"github.com/DBrange/didis-comp-bk/domains/repository/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type UserManagerProxyAdapter struct {
	repository *repository.Repository
}

func NewUserManagerProxyAdapter(repository *repository.Repository) *UserManagerProxyAdapter {
	return &UserManagerProxyAdapter{
		repository: repository,
	}
}

func (a *UserManagerProxyAdapter) CreateUser(ctx context.Context, user *dao.CreateUserDAO) error {
	err := a.repository.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (a *UserManagerProxyAdapter) GetUserByID(ctx context.Context, id string) (*dao.GetUserByIDDAO, error) {
	return a.repository.GetUserByID(ctx, id)
}

func (a *UserManagerProxyAdapter) UpdateUser(ctx context.Context,  filter bson.M, update bson.M) error {
	return a.repository.UpdateUser(ctx, filter, update)
}
