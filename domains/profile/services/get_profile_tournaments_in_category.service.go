package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *ProfileService) GetProfileTournamentsInCategory(ctx context.Context, categoryID, competitorID, lastID string, limit int) ([]*dto.GetTournamentsFromCategoryDTORes, error) {
	tournametsDTO, err := s.profileQueryer.GetCompetitorTournamentsInCategory(ctx, categoryID, competitorID, lastID, limit)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "profile", "error getting competitor profile tournaments from category")
	}

	return tournametsDTO, nil
}
