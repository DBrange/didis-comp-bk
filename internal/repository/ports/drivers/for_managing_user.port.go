package ports

import (
	"context"

	repo_dto "github.com/DBrange/didis-comp-bk/internal/repository/models/user/dto"
)

type ForManagingUser interface {
	CreateUser(user *repo_dto.CreateUserDTO) error
	GetUserByID(ctx context.Context, id string) (*repo_dto.GetUserByIDDTO, error)
}
