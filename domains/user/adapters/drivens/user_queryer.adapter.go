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

func (a *UserQueryerAdapter) CreateUser(ctx context.Context, userDTO *user_dto.CreateUserDTOReq) error {
	mappedUser := mappers.CreateUserDTOReqtoDAO(userDTO)

	err := a.drivers.CreateUser(ctx, mappedUser)
	if err != nil {
		return err
	}

	return nil
}

func (a *UserQueryerAdapter) GetUserByID(ctx context.Context, id string) (*user_dto.GetUserByIDDTO, error) {
	userDTO, err := a.drivers.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	mappedUser := mappers.GetUserByIDDAOtoDTO(userDTO)

	return mappedUser, nil
}

func (a *UserQueryerAdapter) UpdateUser(ctx context.Context, userID string, newUserInfo *user_dto.UpdateUserDTOReq) error {
	newUserInfoDAO := mappers.UpdateUserDTOtoDAO(newUserInfo)

	return a.drivers.UpdateUser(ctx, userID, newUserInfoDAO)
}

func (a *UserQueryerAdapter) DeleteUser(ctx context.Context, userID string) (*user_dto.UserRelationsToDeleteDTO, error) {
	userRelationsToDeleteDAO, err := a.drivers.DeleteUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	mappedUserRelationsToDelete := mappers.UserRelationsToDeleteDAOtoDTO(userRelationsToDeleteDAO)

	return mappedUserRelationsToDelete, nil
}
