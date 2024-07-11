package mappers

import (
	user_dto "github.com/DBrange/didis-comp-bk/domains/user/models/dto"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
)

func CreateUserDTOReqtoDAO(userDTO *user_dto.CreateUserDTOReq) *dao.CreateUserDAO {
	userDAO := &dao.CreateUserDAO{
		FirstName:   userDTO.FirstName,
		LastName:    userDTO.LastName,
		Username:    userDTO.Username,
		Password:    userDTO.Password,
		Birthdate:   userDTO.Birthdate,
		LocationID:  userDTO.LocationID,
		PaymentID:   userDTO.PaymentID,
		ScheduleID:  userDTO.ScheduleID,
		Email:       userDTO.Email,
		Phone:       userDTO.Phone,
		Image:       userDTO.Image,
		Active:      userDTO.Active,
		Role:        userDTO.Role,
		Genre:       userDTO.Genre,
		AccessLevel: userDTO.AccessLevel,
	}

	return userDAO
}
