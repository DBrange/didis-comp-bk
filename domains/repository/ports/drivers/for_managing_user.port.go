package ports

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
)

type ForManagingUser interface {
	CreateUser(ctx context.Context, user *dao.CreateUserDAO) error
	GetUserByID(ctx context.Context, id string) (*dao.GetUserByIDDAO, error)
	UpdateUser(ctx context.Context, userID string, newUserInfo *dao.UpdateUserDAOReq) error
	DeleteUser(ctx context.Context, userID string) (*dao.UserRelationsToDeleteDAO, error)
}
