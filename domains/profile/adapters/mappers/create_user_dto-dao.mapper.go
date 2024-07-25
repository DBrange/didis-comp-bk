package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
)

func CreateUserDTOtoDAO(userDTO *dto.CreateUserDTOReq) *dao.CreateUserDAOReq {
	userDAO := &dao.CreateUserDAOReq{
		FirstName: userDTO.FirstName,
		LastName: userDTO.LastName,
		Username: userDTO.Username,
		Email: userDTO.Email,
		Birthdate: userDTO.Birthdate,
		Password: userDTO.Password,
		Phone: userDTO.Phone,
		Image: userDTO.Image,
		Genre: userDTO.Genre,
	}

	return userDAO
}
