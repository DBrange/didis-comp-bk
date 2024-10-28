package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *ProfileService) GetProfileAvailabilityInCategory(ctx context.Context, competitorID, day string) (*models.GetDailyAvailabilityByIDDTORes,string, error) {
	profileInfoDTO,availabilityID, err := s.profileQuerier.GetDailyAvailabilityCompetitorID(ctx, competitorID, day)
	if err != nil {
		return nil,"", customerrors.HandleErrMsg(err, "profile", "error getting profile info from category")
	}

	return profileInfoDTO,availabilityID, nil
}
