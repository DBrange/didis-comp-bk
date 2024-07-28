package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
)

func GetUserByIDDAOtoDTO(userDAO *dao.GetUserByIDDAORes) *dto.GetUserByIDDTORes{
	userDTO := &dto.GetUserByIDDTORes{
		ID: userDAO.ID,
		FirstName: userDAO.FirstName,
		LastName: userDAO.LastName,
		Username: userDAO.Username,
		Birthdate: userDAO.Birthdate,
		Email: userDAO.Email,
		Phone: userDAO.Phone,
		Image: userDAO.Image,
		Genre: userDAO.Genre,
		LocationID: userDAO.LocationID,
	}

	return userDTO
}