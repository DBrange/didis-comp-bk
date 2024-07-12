package services

import (
	"context"
	"errors"
	"fmt"

	user_dto "github.com/DBrange/didis-comp-bk/domains/user/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (driver *UserService) UpdateUser(ctx context.Context, userID string, newUserInfo *user_dto.UpdateUserDTOReq) error {
	if err := driver.userQueryer.UpdateUser(ctx, userID, newUserInfo); err != nil {
		return updateUserHandleError(err)
	}

	return nil
}

type updateUserErrorHandler func(error) customerrors.AppError

var updateUserErrorHandlers = map[error]updateUserErrorHandler{
	customerrors.ErrUserUpdated: func(err error) customerrors.AppError {
		return customerrors.AppError{
			Code: customerrors.ErrCodeUpdated,
			Msg:  fmt.Sprintf("error updating user: %v", err),
		}
	},
}

func updateUserHandleError(err error) error {
	for knownErr, handler := range updateUserErrorHandlers {
		if errors.Is(err, knownErr) {
			return handler(err)
		}
	}
	return fmt.Errorf("error updating user: %w", err)
}
