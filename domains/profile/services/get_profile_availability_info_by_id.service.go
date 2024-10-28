package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (d *ProfileService) GetProfileDailyAvailabilityByID(ctx context.Context, userID string, day string) (*models.GetDailyAvailabilityByIDDTORes,string, error) {
	availabilityInfo, availabilityID,err := d.profileQuerier.GetDailyAvailabilityUserID(ctx, userID, day)
	if err != nil {
		return nil,"", customerrors.HandleErrMsg(err, "profile", "error getting profile 'availability' info")
	}

	return availabilityInfo,availabilityID, nil
}
