package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/category_registration/dao"
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
)

func GetCategoryRegistrationSortedByPointsDAOtoDTO(categoryRegistrationSortedDAO []*dao.GetCategoryRegistrationSortedByPointsDAORes) []*dto.GetCategoryRegistrationSortedByPointsDTORes {
	categoryRegistrationSortedDTO := make([]*dto.GetCategoryRegistrationSortedByPointsDTORes, len(categoryRegistrationSortedDAO))

	for i, crDAO := range categoryRegistrationSortedDAO {
		categoryRegistrationSortedDTO[i] = &dto.GetCategoryRegistrationSortedByPointsDTORes{
			CompetitorID:    crDAO.CompetitorID.Hex(),
			CurrentPosition: crDAO.CurrentPosition,
		}
	}

	return categoryRegistrationSortedDTO
}
