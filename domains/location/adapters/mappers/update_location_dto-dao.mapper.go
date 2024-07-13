package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/location/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
)

func UpdateLocationDTOtoDAO(newLocationInfoDTO *dto.UpdateLocationDTOReq) *dao.UpdateLocationDAOReq {
	newLocationInfoDAO := &dao.UpdateLocationDAOReq{
		ID:      newLocationInfoDTO.ID,
		State:   newLocationInfoDTO.State,
		City:    newLocationInfoDTO.City,
		Country: newLocationInfoDTO.Country,
		Lat:     newLocationInfoDTO.Lat,
		Long:    newLocationInfoDTO.Long,
	}

	return newLocationInfoDAO
}
