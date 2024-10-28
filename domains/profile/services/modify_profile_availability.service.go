package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (d *ProfileService) ModifyProfileAvailability(ctx context.Context, availabilityID string, availabilityInfoDTO *models.UpdateDailyAvailabilityDTOReq) error {
	if err := d.profileQuerier.UpdateAvailability(ctx, availabilityID, availabilityInfoDTO); err != nil {
		return customerrors.HandleErrMsg(err, "profile", "error updating profile daily availability")
	}

	return nil
}
