package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
)

func ModifyPersonalInfoMapper(profileDTO *dto.ModifyPersonalInfoDTOReq) (*dto.UpdateUserDTOReq, *dto.UpdateLocationDTOReq) {
	userInfoDAO := &dto.UpdateUserDTOReq{
		FirstName: profileDTO.FirstName,
		LastName:  profileDTO.LastName,
		Username:  profileDTO.Username,
		Image:     profileDTO.Image,
		Phone:     profileDTO.Phone,
	}

	if profileDTO.Location != nil {
		locationInfoDAO := &dto.UpdateLocationDTOReq{
			ID:      profileDTO.Location.ID,
			State:   profileDTO.Location.State,
			Country: profileDTO.Location.Country,
			City:    profileDTO.Location.City,
			Lat:     profileDTO.Location.Lat,
			Long:    profileDTO.Location.Long,
		}
		return userInfoDAO, locationInfoDAO
	}

	return userInfoDAO, nil
}
