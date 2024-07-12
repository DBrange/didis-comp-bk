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

func (a *UserProxyAdapter) CreateUser(ctx context.Context, userDTO *user_dto.CreateUserDTOReq) error {
	return a.userService.CreateUser(ctx, userDTO)

}

func (a *UserProxyAdapter) GetUserByID(ctx context.Context, id string) (*user_dto.GetUserByIDDTO, error) {
	return a.userService.GetUserByID(ctx, id)
}

func (a *UserProxyAdapter) UpdateUser(ctx context.Context, userID string, newUserInfo *user_dto.UpdateUserDTOReq) error {
	return a.userService.UpdateUser(ctx, userID, newUserInfo)
}

func (a *UserProxyAdapter) DeleteUser(ctx context.Context, userID string) (*user_dto.UserRelationsToDeleteDTO, error) {
	return a.userService.DeleteUser(ctx, userID)
}
