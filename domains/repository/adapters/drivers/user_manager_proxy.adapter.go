package adapters

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
	"github.com/DBrange/didis-comp-bk/domains/repository/repository"
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
	return a.repository.CreateUser(ctx, user)
}

func (a *UserManagerProxyAdapter) GetUserByID(ctx context.Context, id string) (*dao.GetUserByIDDAO, error) {
	return a.repository.GetUserByID(ctx, id)
}

func (a *UserManagerProxyAdapter) UpdateUser(ctx context.Context, userID string, newUserInfo *dao.UpdateUserDAOReq) error {
	return a.repository.UpdateUser(ctx, userID, newUserInfo)
}

func (a *UserManagerProxyAdapter) DeleteUser(ctx context.Context, userID string) (*dao.UserRelationsToDeleteDAO, error) {
	return a.repository.DeleteUser(ctx, userID)
}
