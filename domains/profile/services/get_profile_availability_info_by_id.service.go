package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (d *ProfileService) GetProfileDailyAvailabilityByID(ctx context.Context, userID string, day string) (*dto.GetDailyAvailabilityByIDDTORes, error) {
	availabilityInfo, err := d.profileQuerier.GetDailyAvailabilityUserID(ctx, userID, day)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "profile", "error getting profile 'availability' info")
	}

	return availabilityInfo, nil
}
