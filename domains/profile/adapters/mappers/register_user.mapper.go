package mappers

import (
	profile_dto "github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
)

func RegisterUserMapper(profileInfoDTO *profile_dto.RegisterUserDTOReq) (*profile_dto.CreateUserDTOReq, *profile_dto.CreateLocationDTOReq) {
	profileInfoDAO := &profile_dto.CreateUserDTOReq{
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

	if profileInfoDTO.Location != nil {
		locationInfoDAO := &profile_dto.CreateLocationDTOReq{
			State:   profileInfoDTO.Location.State,
			City:    profileInfoDTO.Location.City,
			Country: profileInfoDTO.Location.Country,
			Lat:     profileInfoDTO.Location.Lat,
			Long:    profileInfoDTO.Location.Long,
		}

		return profileInfoDAO, locationInfoDAO
	}

	return profileInfoDAO, &profile_dto.CreateLocationDTOReq{}
}
