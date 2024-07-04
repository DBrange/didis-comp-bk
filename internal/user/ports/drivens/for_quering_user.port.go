package ports

import (
	"context"

	dto_user "github.com/DBrange/didis-comp-bk/internal/user/models/dto"
)

type ForQueryingUser interface {
	CreateUser(ctx context.Context, userDTO *dto_user.CreateUserDTO) error
	GetUserByID(ctx context.Context, id string) (*dto_user.GetUserByIDDTO, error)
}
