package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *CategoryService) GetTournamentsFromCategory(ctx context.Context, categoryID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE, limit int, lastID string) (*dto.GetTournamentsFromCategoryDTORes, error) {
	tournamentsDTO, err := s.categoryQuerier.GetTournamentsFromCategory(ctx, categoryID, sport, competitorType, limit, lastID)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "category", "error when getting tournaments from category")
	}

	tournamentsQuantity, err := s.categoryQuerier.GetTournamentsFromCategoryNumber(ctx, categoryID, sport, competitorType)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "category", "error when getting tournaments quantity")
	}

	tournamentsFromCategory := &dto.GetTournamentsFromCategoryDTORes{
		Tournaments: tournamentsDTO,
		Total:       tournamentsQuantity,
	}

	return tournamentsFromCategory, nil
}
