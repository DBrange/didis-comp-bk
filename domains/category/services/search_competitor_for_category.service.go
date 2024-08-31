package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
)

func (s *CategoryService) SearchCompetitorForCategory(ctx context.Context, userID string, name string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) ([]*dto.GetCompetitorFollowedDTORes, error) {
	return s.categoryQuerier.GetCompetitorsFollowed(ctx, userID, name, sport, competitorType)
}
