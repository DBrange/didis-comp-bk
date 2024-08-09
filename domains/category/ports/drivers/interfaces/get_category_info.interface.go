package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
)

type GetCategoryInfo interface {
	GetCategoryInfo(ctx context.Context, categoryID string) (*dto.GetCategoryInfoByIDDTORes, error)
}
