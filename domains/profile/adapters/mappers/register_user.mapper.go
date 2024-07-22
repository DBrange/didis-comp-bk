package mappers

import (
	profile_dto "github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
)

func RegisterUserMapper(profileInfoDTO *profile_dto.RegisterUserDTOReq) (*user_dao.CreateUserDAO, *location_dao.CreateLocationDAOReq) {
	profileInfoDAO := &user_dao.CreateUserDAO{
		FirstName: profileInfoDTO.FirstName,
		LastName:  profileInfoDTO.LastName,
		Username:  profileInfoDTO.Username,
		Password:  profileInfoDTO.Password,
		Birthdate: profileInfoDTO.Birthdate,
		Image:     profileInfoDTO.Image,
		Phone:     profileInfoDTO.Phone,
		Genre:     profileInfoDTO.Genre,
		Email:     profileInfoDTO.Email,
	}

	locationInfoDAO := &location_dao.CreateLocationDAOReq{
		State:   profileInfoDTO.Location.State,
		City:    profileInfoDTO.Location.City,
		Country: profileInfoDTO.Location.Country,
		Lat:     profileInfoDTO.Location.Lat,
		Long:    profileInfoDTO.Location.Long,
	}

	return profileInfoDAO, locationInfoDAO
}
