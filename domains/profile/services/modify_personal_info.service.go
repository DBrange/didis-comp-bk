package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (d *ProfileService) ModifyPersonalInfo(ctx context.Context, userID string, userInfoDTO *dto.ModifyPersonalInfoDTOReq) error {
	if err := d.profileQueryer.ModifyPersonalInfo(ctx, userID, userInfoDTO); err != nil {
		profileErrorHandlers := customerrors.CreateErrorHandlers("profile")
		errMsgTemplate := "error updating profile"
		return customerrors.HandleError(err, profileErrorHandlers, errMsgTemplate)
	}

	return nil
}
