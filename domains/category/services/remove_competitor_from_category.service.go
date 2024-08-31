package services

import (
	"context"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *CategoryService) RemoveCompetitorFromCategory(ctx context.Context, categoryRegistrationID string) error {
	if err := s.categoryQuerier.PermaDeleteCategoryRegistration(ctx, s.categoryQuerier.CategoryRegistrationColl(), categoryRegistrationID); err != nil {
		return customerrors.HandleErrMsg(err, "category", "error when deleting categoryRegistration")
	}

	return nil
}
