package adapters

import (
	"context"

	user_dto "github.com/DBrange/didis-comp-bk/domains/user/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/user/services"
)

type UserProxyAdapter struct {
	userService *services.UserService
}

func NewUserProxyAdapter(userService *services.UserService) *UserProxyAdapter {
	return &UserProxyAdapter{
		userService: userService,
	}
}

func (a *UserProxyAdapter) RegisterUser(ctx context.Context, userInfoDTO *user_dto.RegisterUserDTOReq) error {
	return a.userService.RegisterUser(ctx, userInfoDTO)
}

