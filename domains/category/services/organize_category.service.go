package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *CategoryService) OrganizeCategory(ctx context.Context, organizerID string, categoryDTO *dto.CreateCategoryDTOReq) error {
	if err := s.categoryQuerier.VerifyOrganizerExists(ctx, organizerID); err != nil {
		return customerrors.HandleErrMsg(err, "category", "error organizer not exits")
	}

	categoryID, err := s.categoryQuerier.CreateCategory(ctx, organizerID, categoryDTO)
	if err != nil {
		return customerrors.HandleErrMsg(err, "category", "error when creating category")
	}

	if err := s.categoryQuerier.AddCategoryInOrganizer(ctx, organizerID, categoryID); err != nil {
		return customerrors.HandleErrMsg(err, "category", "error when creating category")
	}

	return nil

}
