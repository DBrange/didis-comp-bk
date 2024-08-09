package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
)

type OrganizeCategory interface {
	OrganizeCategory(ctx context.Context, organizerID string, categoryDTO *dto.CreateCategoryDTOReq) error
}
