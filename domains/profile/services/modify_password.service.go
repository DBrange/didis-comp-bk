package services

import (
	"context"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (d *ProfileService) ModifyPassword(ctx context.Context, userID, newPassword, oldPassword string) error {
	err := d.profileQueryer.ModifyPassword(ctx, userID, newPassword, oldPassword)
	if err != nil {
		profileErrorHandlers := customerrors.CreateErrorHandlers("profile")
		errMsgTemplate := "error updating profile password"
		return customerrors.HandleError(err, profileErrorHandlers, errMsgTemplate)
	}

	return nil
}
