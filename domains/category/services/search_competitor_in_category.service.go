package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *CategoryService) SearchCompetitorInCategory(ctx context.Context, categoryID string, name string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) ([]*dto.GetCompetitorsOfCategoryCompetitorDTORes, error) {
	competitorDTOs, err := s.categoryQuerier.GetCompetitorsOfCategoryByName(ctx, categoryID, name, sport, competitorType)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "category", "error when searching competitor name")
	}

	return competitorDTOs, nil
}
