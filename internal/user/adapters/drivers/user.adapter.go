package adapters

import (
	"context"

	user_dto "github.com/DBrange/didis-comp-bk/internal/user/models/dto"
	"github.com/DBrange/didis-comp-bk/internal/user/services"
)

type UserAdapter struct {
	ctx         context.Context
	userService *services.UserService
}

func NewUserAdapter(ctx context.Context, userService *services.UserService) *UserAdapter {
	return &UserAdapter{
		ctx:         ctx,
		userService: userService,
	}
}

func (a *UserAdapter) CreateUser(userDTO *user_dto.CreateUserDTO) error {
	return a.userService.CreateUser(a.ctx, userDTO)

}

func (a *UserAdapter) GetUserByID(id string) (*user_dto.GetUserByIDDTO, error) {
	return a.userService.GetUserByID(a.ctx, id)
}
