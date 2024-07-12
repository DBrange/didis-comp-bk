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
		return nil, getUserByIDHandleError(err)
	}

	return userDTO, nil
}

type getUserByIDErrorHandler func(error) customerrors.AppError

var getUserByIDErrorHandlers = map[error]getUserByIDErrorHandler{
	customerrors.ErrUserInvalidID: func(err error) customerrors.AppError {
		return customerrors.AppError{
			Code: customerrors.ErrCodeDuplicateKey,
			Msg:  fmt.Sprintf("invalid user id format: %v", err),
		}
	},
	customerrors.ErrUserNotFound: func(err error) customerrors.AppError {
		return customerrors.AppError{
			Code: customerrors.ErrCodeSchemaViolation,
			Msg:  fmt.Sprintf("error when searching for user: %v", err),
		}
	},
}

func getUserByIDHandleError(err error) error {
	for knownErr, handler := range getUserByIDErrorHandlers {
		if errors.Is(err, knownErr) {
			return handler(err)
		}
	}
	return fmt.Errorf("error when searching for user: %w", err)
}
