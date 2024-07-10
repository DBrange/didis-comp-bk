package services

import (
	"context"
	"errors"
	"fmt"

	user_dto "github.com/DBrange/didis-comp-bk/domains/user/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (d *UserService) CreateUser(ctx context.Context, userDTO *user_dto.CreateUserDTOReq) error {
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
