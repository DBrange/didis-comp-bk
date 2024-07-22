package services

import (
	"context"
"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (d *ProfileService) GetProfileAvailabilityInfoByID(ctx context.Context, userID string, day string) (*dto.GetProfileDailyAvailabilityInfoByIDDTORes, error) {
	availabilityInfo, err := d.profileQueryer.GetProfileAvailabilityInfoByID(ctx, userID, day)
	if err != nil {
		profileErrorHandlers := customerrors.CreateErrorHandlers("profile")
		errMsgTemplate := "error getting profile 'availability' info"
		return nil, customerrors.HandleError(err, profileErrorHandlers, errMsgTemplate)
	}

	return availabilityInfo, nil
}
