package adapters

import (
	"context"
	repo_dto "didis-comp-bk/internal/repository/models/user/dto"
	ports "didis-comp-bk/internal/repository/ports/drivers"
	user_dto "didis-comp-bk/internal/user/models/dto"
)

type UserQueryerAdapter struct {
	ctx     context.Context
	drivers ports.ForManagingUser
}

func NewUserQueryerAdapter(ctx context.Context, drivers ports.ForManagingUser) *UserQueryerAdapter {
	return &UserQueryerAdapter{
		ctx:     ctx,
		drivers: drivers,
	}
}

func (a *UserQueryerAdapter) CreateUser(ctx context.Context, userDTO *user_dto.CreateUserDTO) error {
	mappedDTO := &repo_dto.CreateUserDTO{
		FirstName: userDTO.FirstName,
		LastName:  userDTO.LastName,
	}
	mappedDTO.SetTimeStamp()

	err := a.drivers.CreateUser(mappedDTO)
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
	mappedDTO := &user_dto.GetUserByIDDTO{
		FirstName: userDTO.FirstName,
		LastName:  userDTO.LastName,
	}

	return mappedDTO, nil
}
