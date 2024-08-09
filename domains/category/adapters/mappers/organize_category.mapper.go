package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
)

func OrganizeCategoryMapper(organizerCategory *dto.OrganizeCategoryDTOReq) *dto.CreateCategoryDTOReq {
	categoryDTO := &dto.CreateCategoryDTOReq{
		Name:              organizerCategory.Name,
		Genre:             organizerCategory.Genre,
		TotalParticipants: organizerCategory.TotalParticipants,
		RangeMovement:     organizerCategory.RangeMovement,
		Sport:             organizerCategory.Sport,
	}

	return categoryDTO
}
