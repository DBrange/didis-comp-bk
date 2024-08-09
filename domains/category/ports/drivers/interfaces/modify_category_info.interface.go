package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
)

type ModifyCategoryInfo interface {
	ModifyCategoryInfo(ctx context.Context, categoryID string, categoryInfoDTO *dto.UpdateCategoryDTOReq) error
}
