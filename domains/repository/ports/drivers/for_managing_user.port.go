package ports

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
	"go.mongodb.org/mongo-driver/bson"
)

type ForManagingUser interface {
	CreateUser(ctx context.Context, user *dao.CreateUserDAO) error
	GetUserByID(ctx context.Context, id string) (*dao.GetUserByIDDAO, error)
	UpdateUser(ctx context.Context, filter bson.M, update bson.M) error
}
