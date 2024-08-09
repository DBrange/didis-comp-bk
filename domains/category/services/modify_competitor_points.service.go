package services

import (
	"context"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *CategoryService) ModifyCompetitorPoints(ctx context.Context, categoryID, competitorID string, points int) error {
	if err := s.categoryQueryer.UpdateCompetitorPoints(ctx, categoryID, competitorID, points); err != nil {
		return customerrors.HandleErrMsg(err, "category", "error when modifing competitor points")
	}

	return nil
}
