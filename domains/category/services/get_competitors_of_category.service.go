package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *CategoryService) GetParticipantsOfCategory(ctx context.Context, categoryID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE, limit int, lastID string) (*dto.GetCompetitorsOfCategoryDTORes, error) {
	if err := s.categoryQuerier.VerifyCategoryExists(ctx, categoryID); err != nil {
		return nil, customerrors.HandleErrMsg(err, "category", "error category not exits")
	}

	competitorDTOs, err := s.categoryQuerier.GetParticipantsOfCategory(ctx, categoryID, sport, competitorType, limit, lastID)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "category", "error when getting competitors")
	}
	competitorsNumber, err := s.categoryQuerier.GetCategoryCompetitorsNumber(ctx, categoryID)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "category", "error when searching competitor name")
	}

	competitors := &dto.GetCompetitorsOfCategoryDTORes{
		Competitors: competitorDTOs,
		Total:       competitorsNumber,
	}

	return competitors, nil
}
