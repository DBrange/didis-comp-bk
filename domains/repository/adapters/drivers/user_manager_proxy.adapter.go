package adapters

import (
	"context"

	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
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

func (a *UserManagerProxyAdapter) CreateUserAndLocation(ctx context.Context, userInfoDAO *user_dao.CreateUserDAO, locationInfoDAO *location_dao.CreateLocationDAOReq) error {
	return a.repository.CreateUserAndLocation(ctx, userInfoDAO, locationInfoDAO)
}

func (a *UserManagerProxyAdapter) CreateUser(ctx context.Context, user *user_dao.CreateUserDAO) error {
	return a.repository.CreateUser(ctx, user)
}

func (a *UserManagerProxyAdapter) GetUserByID(ctx context.Context, id string) (*user_dao.GetUserByIDDAO, error) {
	return a.repository.GetUserByID(ctx, id)
}

func (a *UserManagerProxyAdapter) UpdateUser(ctx context.Context, userID string, newUserInfo *user_dao.UpdateUserDAOReq) error {
	return a.repository.UpdateUser(ctx, userID, newUserInfo)
}

func (a *UserManagerProxyAdapter) DeleteUser(ctx context.Context, userID string) (*user_dao.UserRelationsToDeleteDAO, error) {
	return a.repository.DeleteUser(ctx, userID)
}
