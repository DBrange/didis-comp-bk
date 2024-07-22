package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (d *ProfileService) GetPersonalInfoByID(ctx context.Context, userID string) (*dto.GetPersonalInfoByIDDTORes, error) {
	personalInfo, err := d.profileQueryer.GetPersonalInfoByID(ctx, userID)
	if err != nil {
		profileErrorHandlers := customerrors.CreateErrorHandlers("profile")
		errMsgTemplate := "error getting profile personal info"
		return nil, customerrors.HandleError(err, profileErrorHandlers, errMsgTemplate)
	}

	return personalInfo, nil
}
