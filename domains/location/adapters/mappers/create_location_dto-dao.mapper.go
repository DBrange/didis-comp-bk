package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/location/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
)

func CreateLocationDTOtoDAO(userDTO *dto.CreateLocationDTOReq) *dao.CreateLocationDAOReq {
	userDAO := &dao.CreateLocationDAOReq{
		State:   userDTO.State,
		Country: userDTO.Country,
		City:    userDTO.City,
		Lat:     userDTO.Lat,
		Long:    userDTO.Long,
	}

	return userDAO
}
