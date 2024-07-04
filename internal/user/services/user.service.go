package services

import (
	"context"
	"errors"
	"fmt"

	user_dto "github.com/DBrange/didis-comp-bk/internal/user/models/dto"
	ports "github.com/DBrange/didis-comp-bk/internal/user/ports/drivens"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

type UserService struct {
	userQueryer ports.ForQueryingUser
}

func NewUserService(userQueryer ports.ForQueryingUser) *UserService {
	return &UserService{
		userQueryer: userQueryer,
	}
}

func (d *UserService) CreateUser(ctx context.Context, userDTO *user_dto.CreateUserDTO) error {
	err := d.userQueryer.CreateUser(ctx, userDTO)
	if err != nil {
		if errors.Is(err, customerrors.ErrUserDuplicateKey) {
			appErr := customerrors.AppError{
				Code: customerrors.ErrCodeDuplicateKey,
				Msg:  fmt.Sprintf("error inserting user: %v", err),
			}
			return appErr
		}
		return fmt.Errorf("error inserting user: %w", err)
	}

	return nil
}

func (d *UserService) GetUserByID(ctx context.Context, id string) (*user_dto.GetUserByIDDTO, error) {
	userDTO, err := d.userQueryer.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, customerrors.ErrUserNotFound) {
			appErr := customerrors.AppError{
				Code: customerrors.ErrCodeNotFound,
				Msg:  "error getting user: id not exists",
			}
			return nil, appErr
		}
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	return userDTO, nil
}
