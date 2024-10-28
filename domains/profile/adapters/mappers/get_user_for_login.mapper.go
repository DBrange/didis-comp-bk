package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
)

func GetUserForLoginDAOtoDTO(userDAO *dao.GetUserForLoginDAO) *dto.GetUserForLoginDTO {
	userDTO := &dto.GetUserForLoginDTO{
		ID:        userDAO.ID,
		FirstName: userDAO.FirstName,
		LastName:  userDAO.LastName,
		Password:  userDAO.Password,
		Username:  userDAO.Username,
		Image:     userDAO.Image,
		Roles:     userDAO.Roles,
	}

	return userDTO
}
