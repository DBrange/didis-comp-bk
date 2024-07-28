package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/profile/adapters/mappers"
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *ProfileService) GetPersonalInfoByID(ctx context.Context, userID string) (*dto.GetPersonalInfoByIDDTORes, error) {
	userInfo, err := s.profileQueryer.GetUserByID(ctx, userID)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "profile", "error getting personal info")
	}

	locationInfo, err := s.profileQueryer.GetLocationByID(ctx, userInfo.LocationID)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "profile", "error getting local personal info")
	}

	personalInfo := mappers.GetPersonalInfoByIDMapper(userInfo, locationInfo)

	return personalInfo, nil
}
