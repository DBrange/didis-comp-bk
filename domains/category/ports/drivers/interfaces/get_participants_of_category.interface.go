package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
)

type GetParticipantsOfCategory interface {
	GetParticipantsOfCategory(ctx context.Context, categoryID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE, limit int, lastID string) (*dto.GetCompetitorsOfCategoryDTORes, error)
}
