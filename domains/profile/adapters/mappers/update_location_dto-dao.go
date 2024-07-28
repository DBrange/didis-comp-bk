package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
)

func UpdateLocationDTOtoDAO(locationDTO *dto.UpdateLocationDTOReq) *dao.UpdateLocationDAOReq {
	locationDAO := &dao.UpdateLocationDAOReq{
		ID:      locationDTO.ID,
		State:   locationDTO.State,
		Country: locationDTO.Country,
		City:    locationDTO.City,
		Lat:     locationDTO.Lat,
		Long:    locationDTO.Long,
	}

	return locationDAO
}
