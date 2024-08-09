package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *CategoryService) GetCategoryInfo(ctx context.Context, categoryID string) (*dto.GetCategoryInfoByIDDTORes, error) {
	categoryInfoDTO, err := s.categoryQueryer.GetCategoryInfoByID(ctx, categoryID)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "category", "error when getting category info")
	}

	return categoryInfoDTO, nil
}
