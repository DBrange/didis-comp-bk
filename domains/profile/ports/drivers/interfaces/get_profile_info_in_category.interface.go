package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
)

type GetProfileInfoInCategory interface {
	GetProfileInfoInCategory(ctx context.Context, categoryID, competitorID string) (*dto.GetProfileInfoInCategoryDTORes, error)
}
