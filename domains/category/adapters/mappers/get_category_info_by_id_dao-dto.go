package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/category/dao"
)

func GetCategoryInfoByIDDAOtoDTO(categoryInfoDAO *dao.GetCategoryInfoByIDDAORes) *dto.GetCategoryInfoByIDDTORes {
	categoryInfoDTO := &dto.GetCategoryInfoByIDDTORes{
		ID:                categoryInfoDAO.ID.Hex(),
		Name:              categoryInfoDAO.Name,
		Genre:             categoryInfoDAO.Genre,
		TotalParticipants: categoryInfoDAO.TotalParticipants,
		RangeMovement:     categoryInfoDAO.RangeMovement,
		Sport:             categoryInfoDAO.Sport,
		CompetitorType:    categoryInfoDAO.CompetitorType,
		Organizer: dto.GetCategoryInfoOrganizerByIDDTORes{
			ID:        categoryInfoDAO.Organizer.ID.Hex(),
			FirstName: categoryInfoDAO.Organizer.FirstName,
			LastName:  categoryInfoDAO.Organizer.LastName,
		},
	}

	return categoryInfoDTO
}
