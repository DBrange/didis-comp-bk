package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *CategoryService) AddCompetitorInCategory(ctx context.Context, categoryID, competitorID string) error {
	// Verify if category and competitor exists
	if err := s.categoryQuerier.VerifyCategoryExists(ctx, categoryID); err != nil {
		return customerrors.HandleErrMsg(err, "category", "error verify category exists")
	}

	if err := s.categoryQuerier.VerifyCompetitorExists(ctx, competitorID); err != nil {
		return customerrors.HandleErrMsg(err, "category", "error verify competitor exists")
	}

	categoryRegistrationDTO := &dto.CreateCategoryRegistrationDTOReq{
		CompetitorID: competitorID,
		CategoryID:   categoryID,
	}

	// Verify if category and competitor exists
	if err := s.categoryQuerier.VerifyCategoryExistsRelation(ctx, categoryRegistrationDTO); err != nil {
		return customerrors.HandleErrMsg(err, "category", "error verify category exists")
	}

	if err := s.categoryQuerier.CreateCategoryRegistration(ctx, categoryRegistrationDTO); err != nil {
		return customerrors.HandleErrMsg(err, "category", "error when creating categoryRegistration")
	}

	if err := s.categoryQuerier.IncrementTotalParticipants(ctx, categoryID); err != nil {
		return customerrors.HandleErrMsg(err, "category", "error when increment total participants")
	}

	return nil
}
