package adapters

import (
	"context"

	ports "github.com/DBrange/didis-comp-bk/domains/repository/ports/drivers"
	"github.com/DBrange/didis-comp-bk/domains/user/adapters/mappers"
	user_dto "github.com/DBrange/didis-comp-bk/domains/user/models/dto"
)

type UserQueryerAdapter struct {
	drivers ports.ForManagingUser
}

func NewUserQueryerAdapter(drivers ports.ForManagingUser) *UserQueryerAdapter {
	return &UserQueryerAdapter{
		drivers: drivers,
	}
}

func (a *UserQueryerAdapter) RegisterUser(ctx context.Context, userInfoDTO *user_dto.RegisterUserDTOReq) error {
	userInfoDAO, locationInfoDAO := mappers.RegisterUserMapper(userInfoDTO)

	return a.drivers.CreateUserAndLocation(ctx, userInfoDAO, locationInfoDAO)
}
