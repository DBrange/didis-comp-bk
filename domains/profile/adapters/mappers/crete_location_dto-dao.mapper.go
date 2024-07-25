package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
)

func CreateLocationDTOtoDAO(locationDTO *dto.CreateLocationDTOReq) *dao.CreateLocationDAOReq {
	locationDAO := &dao.CreateLocationDAOReq{
		State:   locationDTO.State,
		Country: locationDTO.Country,
		City:    locationDTO.City,
		Lat:     locationDTO.Lat,
		Long:    locationDTO.Long,
	}

	return locationDAO
}
