package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
)

func CreateUserDTOtoDAO(userDTO *dto.CreateUserDTOReq, convert utils.ConvertToObjectIDFunc) (*dao.CreateUserDAOReq, error) {
	locationIOD, err := convert(*userDTO.LocationID)
	if err != nil {
		return nil, nil
	}

	userDAO := &dao.CreateUserDAOReq{
		FirstName: userDTO.FirstName,
		LastName:  userDTO.LastName,
		Username:  userDTO.Username,
		Email:     userDTO.Email,
		Birthdate: userDTO.Birthdate,
		Password:  userDTO.Password,
		Phone:     userDTO.Phone,
		Image:     userDTO.Image,
		Genre:     userDTO.Genre,
		LocationID: locationIOD,
	}


	if userDTO.Roles != nil {
		rolesOID, err := utils.ConvertToObjectIDs(&userDTO.Roles, convert)
		if err != nil {
			return nil, err
		}

		userDAO.Roles = *rolesOID
	}

	return userDAO, nil
}
