package services

import (
	"context"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *CategoryService) RemoveCompetitorFromCategory(ctx context.Context, categoryID, competitorID string) error {
	if err := s.categoryQuerier.DeleteCategoryRegistration(ctx, categoryID, competitorID); err != nil {
		return customerrors.HandleErrMsg(err, "category", "error when deleting categoryRegistration")
	}
	
	if err := s.UpdateCategoryRanking(ctx, categoryID); err != nil {
		return err
	}
	return nil
}
