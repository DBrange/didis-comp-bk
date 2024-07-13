package ports

import (
	"context"

	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
)

type ForManagingUser interface {
	CreateUserAndLocation(ctx context.Context, userDAO *user_dao.CreateUserDAO, locationDAO *location_dao.CreateLocationDAOReq) error
	CreateUser(ctx context.Context, userDAO *user_dao.CreateUserDAO) error
	GetUserByID(ctx context.Context, userID string) (*user_dao.GetUserByIDDAO, error)
	UpdateUser(ctx context.Context, userID string, newUserInfo *user_dao.UpdateUserDAOReq) error
	DeleteUser(ctx context.Context, userID string) (*user_dao.UserRelationsToDeleteDAO, error)
}
