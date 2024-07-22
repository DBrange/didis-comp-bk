package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
)

func ModifyPersonalInfoMapper(userInfoDTO *dto.ModifyPersonalInfoDTOReq) (*user_dao.UpdateUserDAOReq, *location_dao.UpdateLocationDAOReq) {
	userInfoDAO := &user_dao.UpdateUserDAOReq{
		FirstName: userInfoDTO.FirstName,
		LastName:  userInfoDTO.LastName,
		Username:  userInfoDTO.Username,
		Image:     userInfoDTO.Image,
		Phone:     userInfoDTO.Phone,
	}

	if userInfoDTO.Location != nil {
		locationInfoDAO := &location_dao.UpdateLocationDAOReq{
			ID:      userInfoDTO.Location.ID,
			State:   userInfoDTO.Location.State,
			Country: userInfoDTO.Location.Country,
			City:    userInfoDTO.Location.City,
			Lat:     userInfoDTO.Location.Lat,
			Long:    userInfoDTO.Location.Long,
		}
		return userInfoDAO, locationInfoDAO
	}

	return userInfoDAO, nil
}
