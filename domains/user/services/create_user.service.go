package services

import (
	"context"
	"errors"
	"fmt"

	user_dto "github.com/DBrange/didis-comp-bk/domains/user/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (d *UserService) CreateUser(ctx context.Context, userDTO *user_dto.CreateUserDTOReq) error {
	if err := d.userQueryer.CreateUser(ctx, userDTO); err != nil {
		return createUserHandleError(err)
	}
	return nil
}

type createUserErrorHandler func(error) customerrors.AppError

var createUserErrorHandlers = map[error]createUserErrorHandler{
	customerrors.ErrUserDuplicateKey: func(err error) customerrors.AppError {
		return customerrors.AppError{
			Code: customerrors.ErrCodeDuplicateKey,
			Msg:  fmt.Sprintf("error duplicate key for user: %v", err),
		}
	},
	customerrors.ErrSchemaViolation: func(err error) customerrors.AppError {
		return customerrors.AppError{
			Code: customerrors.ErrCodeSchemaViolation,
			Msg:  fmt.Sprintf("error user schema type: %v", err),
		}
	},
}

func createUserHandleError(err error) error {
	for knownErr, handler := range createUserErrorHandlers {
		if errors.Is(err, knownErr) {
			return handler(err)
		}
	}
	return fmt.Errorf("error inserting user: %w", err)
}

