package mappers

import (
	profile_dto "github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
)

func GetPersonalInfoByIDMapper(userInfo *profile_dto.GetUserByIDDTORes, locationInfo *profile_dto.GetLocationByIDDTORes) *profile_dto.GetPersonalInfoByIDDTORes {
	profileInfo := &profile_dto.GetPersonalInfoByIDDTORes{
		ID:        userInfo.ID,
		FirstName: userInfo.FirstName,
		LastName:  userInfo.LastName,
		Username:  userInfo.Username,
		Email:     userInfo.Email,
		Birthdate: userInfo.Birthdate,
		Image:     userInfo.Image,
		Phone:     userInfo.Phone,
		Genre:     userInfo.Genre,
		Location: &profile_dto.GetPersonalInfoLocationByIDRes{
			ID:      locationInfo.ID,
			State:   locationInfo.State,
			City:    locationInfo.City,
			Country: locationInfo.Country,
			Lat:     locationInfo.Lat,
			Long:    locationInfo.Long,
		},
	}

	return profileInfo
}
