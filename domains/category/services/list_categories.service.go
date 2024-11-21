package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *CategoryService) ListCategories(ctx context.Context, organizerID string, sport models.SPORT, competitorType *models.COMPETITOR_TYPE) ([]dto.GetCategoriesFromOrganizerDTORes, error) {
	categoriesDTO, err := s.categoryQuerier.GetCategoriesFromOrganizer(ctx, organizerID, sport, competitorType)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "category", "error when getting categories")
	}

	return categoriesDTO, nil
}
