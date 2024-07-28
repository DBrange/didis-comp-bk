package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (d *ProfileService) ModifyProfileAvailability(ctx context.Context, availabilityID string, availabilityInfoDTO *dto.UpdateDailyAvailabilityDTOReq) error {
	if err := d.profileQueryer.UpdateAvailability(ctx, availabilityID, availabilityInfoDTO); err != nil {
		return customerrors.HandleErrMsg(err, "profile", "error updating profile daily availability")
	}

	return nil
}
