package mappers

import (
	location_dto "github.com/DBrange/didis-comp-bk/domains/location/models/dto"
	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
)

func GetLocationByIDDAOtoDTO(locationDAO *location_dao.GetLocationByIDDAORes) *location_dto.GetLocationByIDDTORes {
	locationDTO := &location_dto.GetLocationByIDDTORes{
		State: locationDAO.State,
		Country: locationDAO.Country,
		City: locationDAO.City,
		Lat: locationDAO.Lat,
		Long: locationDAO.Long,
	}

	return locationDTO
}
