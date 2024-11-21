package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
)

type ListCategories interface {
	ListCategories(ctx context.Context, organizerID string, sport models.SPORT, competitorType *models.COMPETITOR_TYPE) ([]dto.GetCategoriesFromOrganizerDTORes, error)
}
