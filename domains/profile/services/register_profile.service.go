package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (d *ProfileService) RegisterUser(ctx context.Context, profileInfoDTO *dto.RegisterUserDTOReq) error {
	if err := d.profileQueryer.RegisterUser(ctx, profileInfoDTO); err != nil {
		profileErrorHandlers := customerrors.CreateErrorHandlers("profile")
		errMsgTemplate := "error when registering profile"
		return customerrors.HandleError(err, profileErrorHandlers, errMsgTemplate)
	}

	return nil
}
