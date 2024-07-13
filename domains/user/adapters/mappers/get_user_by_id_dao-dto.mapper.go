package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
	user_dto "github.com/DBrange/didis-comp-bk/domains/user/models/dto"
)

func GetUserByIDDAOtoDTO(userDAO *dao.GetUserByIDDAO) *user_dto.GetUserByIDDTORes {
	userDTO := &user_dto.GetUserByIDDTORes{
		ID:          userDAO.ID,
		FirstName:   userDAO.FirstName,
		LastName:    userDAO.LastName,
		Username:    userDAO.Username,
		Password:    userDAO.Password,
		Birthdate:   userDAO.Birthdate,
		LocationID:  userDAO.LocationID,
		PaymentID:   userDAO.PaymentID,
		ScheduleID:  userDAO.ScheduleID,
		Email:       userDAO.Email,
		Phone:       userDAO.Phone,
		Image:       userDAO.Image,
		Active:      userDAO.Active,
		Role:        userDAO.Role,
		Genre:       userDAO.Genre,
		AccessLevel: userDAO.AccessLevel,
	}

	return userDTO
}
