package mappers

import (
	profile_dto "github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
)

func GetPersonalInfoByIDMapper(userInfo *user_dao.GetUserByIDDAO, locationInfo *location_dao.GetLocationByIDDAORes) *profile_dto.GetPersonalInfoByIDDTORes {
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
