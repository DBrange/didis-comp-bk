package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *ProfileService) GetProfileInfoInCategory(ctx context.Context, categoryID, competitorID string) (*dto.GetProfileInfoInCategoryDTORes, error) {
	profileInfoDTO, err := s.profileQueryer.GetProfileInfoInCategory(ctx, categoryID, competitorID)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "profile", "error getting profile info from category")
	}

	return profileInfoDTO, nil
}
