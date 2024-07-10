package adapters

import (
	"context"

	user_dto "github.com/DBrange/didis-comp-bk/domains/user/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/user/services"
)

type UserProxyAdapter struct {
	ctx         context.Context
	userService *services.UserService
}

func NewUserProxyAdapter(ctx context.Context, userService *services.UserService) *UserProxyAdapter {
	return &UserProxyAdapter{
		ctx:         ctx,
		userService: userService,
	}
}

func (a *UserProxyAdapter) CreateUser(userDTO *user_dto.CreateUserDTOReq) error {
	return a.userService.CreateUser(a.ctx, userDTO)

}

func (a *UserProxyAdapter) GetUserByID(id string) (*user_dto.GetUserByIDDTO, error) {
	return a.userService.GetUserByID(a.ctx, id)
}
