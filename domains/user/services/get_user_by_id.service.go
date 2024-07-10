package services

import (
	"context"
	"errors"
	"fmt"

	user_dto "github.com/DBrange/didis-comp-bk/domains/user/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

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
