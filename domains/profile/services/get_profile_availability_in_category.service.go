package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *ProfileService) GetProfileAvailabilityInCategory(ctx context.Context, competitorID, day string) (*dto.GetDailyAvailabilityByIDDTORes, error) {
	profileInfoDTO, err := s.profileQuerier.GetDailyAvailabilityCompetitorID(ctx, competitorID, day)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "profile", "error getting profile info from category")
	}

	return profileInfoDTO, nil
}
