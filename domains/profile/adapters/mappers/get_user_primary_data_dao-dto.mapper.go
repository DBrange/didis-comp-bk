package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
)

func GetUserPrimaryDataDAOtoDTO(userDAO *dao.GetUserPrimaryDataDAORes) *dto.GetUserPrimaryDataDTORes {
	return &dto.GetUserPrimaryDataDTORes{
		ID:         userDAO.ID.Hex(),
		FirstName:  userDAO.FirstName,
		LastName:   userDAO.LastName,
		Username:   userDAO.Username,
		Image:      userDAO.Image,
		LocationID: userDAO.LocationID.Hex(),
	}
}
