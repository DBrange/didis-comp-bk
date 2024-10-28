package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
)

func GetLocationByIDDAOtoDTO(locationDAO *dao.GetLocationByIDDAORes) *dto.GetLocationByIDDTORes {
	locationDTO := &dto.GetLocationByIDDTORes{
		ID:      locationDAO.ID.Hex(),
		State:   locationDAO.State,
		Country: locationDAO.Country,
		City:    locationDAO.City,
		Lat:     locationDAO.Lat,
		Long:    locationDAO.Long,
	}

	return locationDTO
}
