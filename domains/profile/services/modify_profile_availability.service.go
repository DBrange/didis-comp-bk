package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (d *ProfileService) ModifyProfileAvailability(ctx context.Context, availabilityID string, availabilityInfoDTO *dto.ModifyProfileDailyAvailabilityDTOReq) error {
	if err := d.profileQueryer.ModifyProfileAvailability(ctx, availabilityID, availabilityInfoDTO); err != nil {
		profileErrorHandlers := customerrors.CreateErrorHandlers("profile")
		errMsgTemplate := "error updating profile 'availability'"
		return customerrors.HandleError(err, profileErrorHandlers, errMsgTemplate)
	}

	return nil
}
