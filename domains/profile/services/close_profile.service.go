package services

import (
	"context"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (d *ProfileService) CloseProfile(ctx context.Context, userID string) error {
	err := d.profileQueryer.CloseProfile(ctx, userID)
	if err != nil {
		profileErrorHandlers := customerrors.CreateErrorHandlers("profile")
		errMsgTemplate := "error updating profile deleted_at"
		return customerrors.HandleError(err, profileErrorHandlers, errMsgTemplate)
	}

	return nil
}
