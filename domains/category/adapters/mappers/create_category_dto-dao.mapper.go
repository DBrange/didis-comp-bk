package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/category/dao"
)

func CreateCategoryDTOtoDAO(categoryDTO *dto.CreateCategoryDTOReq) *dao.CreateCategoryDAOReq {
	categoryDAO := &dao.CreateCategoryDAOReq{
		Name:              categoryDTO.Name,
		Genre:             categoryDTO.Genre,
		// TotalParticipants: categoryDTO.TotalParticipants,
		RangeMovement:     categoryDTO.RangeMovement,
		AverageScore:      categoryDTO.AverageScore,
		Sport:             categoryDTO.Sport,
		CompetitorType:    categoryDTO.CompetitorType,
	}

	return categoryDAO
}
