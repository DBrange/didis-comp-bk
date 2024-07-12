package services

import (
	"context"
	"errors"
	"fmt"

	user_dto "github.com/DBrange/didis-comp-bk/domains/user/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (d *UserService) DeleteUser(ctx context.Context, userID string) (*user_dto.UserRelationsToDeleteDTO, error) {
	userRelationsToDelete, err := d.userQueryer.DeleteUser(ctx, userID)
	if err != nil {
		return nil, deleteUserHandleError(err)
	}

	return userRelationsToDelete, nil
}

type deleteUserErrorHandler func(error) customerrors.AppError

var deleteUserErrorHandlers = map[error]deleteUserErrorHandler{
	customerrors.ErrUserInvalidID: func(err error) customerrors.AppError {
		return customerrors.AppError{
			Code: customerrors.ErrCodeInvalidID,
			Msg:  fmt.Sprintf("invalid id format: %v", err),
		}
	},
	customerrors.ErrUserNotFound: func(err error) customerrors.AppError {
		return customerrors.AppError{
			Code: customerrors.ErrCodeNotFound,
			Msg:  fmt.Sprintf("error when searching for the user: %v", err),
		}
	},
	customerrors.ErrUserUpdated: func(err error) customerrors.AppError {
		return customerrors.AppError{
			Code: customerrors.ErrCodeUpdated,
			Msg:  fmt.Sprintf("error updating user: %v", err),
		}
	},
}

func deleteUserHandleError(err error) error {
	for knownErr, handler := range deleteUserErrorHandlers {
		if errors.Is(err, knownErr) {
			return handler(err)
		}
	}
	return fmt.Errorf("error when searching for user: %w", err)
}
