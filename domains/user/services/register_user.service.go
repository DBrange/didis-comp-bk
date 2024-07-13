package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/user/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (d *UserService) RegisterUser(ctx context.Context, userInfoDTO *dto.RegisterUserDTOReq) error {
	if err := d.userQueryer.RegisterUser(ctx, userInfoDTO); err != nil {
		userErrorHandlers := customerrors.CreateErrorHandlers("user")
		errMsgTemplate := "error when registering user"
		return customerrors.HandleError(err, userErrorHandlers, errMsgTemplate)
	}

	return nil
}
