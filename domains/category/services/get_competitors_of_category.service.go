package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *CategoryService) GetParticipantsOfCategory(ctx context.Context, categoryID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE, limit int, lastID string) ([]*dto.GetCompetitorsOfCategoryDTORes, error) {
	if err := s.categoryQueryer.VerifyCategoryExists(ctx,categoryID); err != nil{
		return nil, customerrors.HandleErrMsg(err, "category", "error category not exits")
	}
	
	competitorDTOs, err := s.categoryQueryer.GetParticipantsOfCategory(ctx, categoryID, sport, competitorType, limit, lastID)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "category", "error when getting competitors")
	}

	return competitorDTOs, nil
}
