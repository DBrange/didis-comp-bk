package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
)

type GetProfileTournamentsInCategory interface {
	GetProfileTournamentsInCategory(ctx context.Context, categoryID, competitorID, lastID string, limit int) ([]*dto.GetTournamentsFromCategoryDTORes, error)
}
