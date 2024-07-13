package mappers

import (
	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
	user_dto "github.com/DBrange/didis-comp-bk/domains/user/models/dto"
)

func RegisterUserMapper(userInfoDTO *user_dto.RegisterUserDTOReq) (*user_dao.CreateUserDAO, *location_dao.CreateLocationDAOReq) {
	userInfoDAO := &user_dao.CreateUserDAO{
		FirstName: userInfoDTO.FirstName,
		LastName:  userInfoDTO.LastName,
		Username:  userInfoDTO.Username,
		Password:  userInfoDTO.Password,
		Birthdate: userInfoDTO.Birthdate,
		Image:     userInfoDTO.Image,
		Phone:     userInfoDTO.Phone,
		Genre:     userInfoDTO.Genre,
		Email:     userInfoDTO.Email,
	}

	locationInfoDAO := &location_dao.CreateLocationDAOReq{
		State:   userInfoDTO.Location.State,
		City:    userInfoDTO.Location.City,
		Country: userInfoDTO.Location.Country,
		Lat:     userInfoDTO.Location.Lat,
		Long:    userInfoDTO.Location.Long,
	}

	return userInfoDAO, locationInfoDAO
}
