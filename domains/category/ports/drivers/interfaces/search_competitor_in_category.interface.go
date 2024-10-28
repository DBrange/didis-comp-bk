package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
)

type SearchCompetitorInCategory interface {
	SearchCompetitorInCategory(ctx context.Context, categoryID string, name string, sport models.SPORT, competitorType models.COMPETITOR_TYPE, ) ([]*dto.GetCompetitorsOfCategoryDTORes, error)
}
