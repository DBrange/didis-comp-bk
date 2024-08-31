package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *CategoryService) ModifyCategoryInfo(ctx context.Context, categoryID string, categoryInfoDTO *dto.UpdateCategoryDTOReq) error {
	if err := s.categoryQuerier.UpdateCategory(ctx, categoryID, categoryInfoDTO); err != nil {
		return customerrors.HandleErrMsg(err, "category", "error when modifing category info")
	}

	return nil
}
