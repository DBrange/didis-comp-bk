package services

import (
	"context"
	"errors"
	"fmt"

	user_dto "github.com/DBrange/didis-comp-bk/domains/user/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (driver *UserService) UpdateUser(ctx context.Context, userID string, newUser *user_dto.UpdateUserDTOReq) error {
	err := driver.userQueryer.UpdateUser(ctx, userID, newUser)

	if err != nil {
		if errors.Is(err, customerrors.ErrUserInsertionFailed) {
			appErr := customerrors.AppError{
				Code: customerrors.ErrCodeInsertionFailed,
				Msg:  fmt.Sprintf("error inserting user: %v", err),
			}
			return appErr
		}
		return fmt.Errorf("error inserting user: %w", err)
	}

	return nil
}
