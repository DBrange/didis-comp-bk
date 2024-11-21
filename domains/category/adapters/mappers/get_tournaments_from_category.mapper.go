package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/category/dao"
	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
)

func GetTournamentsFromCategoryDAOtoDTO(tournamentsDAO []*dao.GetTournamentsFromCategoryDAORes) []*dto.GetTournamentsFromCategoryTournamentDTORes {
	tournamentsDTO := make([]*dto.GetTournamentsFromCategoryTournamentDTORes, len(tournamentsDAO))

	for i, tournament := range tournamentsDAO {
		tournamentsDTO[i] = &dto.GetTournamentsFromCategoryTournamentDTORes{
			ID:           tournament.ID.Hex(),
			Name:         tournament.Name,
			Points:       tournament.Points,
			Location:     GetLocationByIDDAOtoDTO(tournament.Location),
			TotalPrize:   tournament.TotalPrize,
			AverageScore: tournament.AverageScore,
			StartDate:    tournament.StartDate,
			FinishtDate:  tournament.FinishtDate,
		}
	}

	return tournamentsDTO
}

func GetLocationByIDDAOtoDTO(locationDAO *location_dao.GetLocationByIDDAORes) *dto.GetLocationByIDDTORes {
	locationDTO := &dto.GetLocationByIDDTORes{
		State:   locationDAO.State,
		Country: locationDAO.Country,
		City:    locationDAO.City,
		Lat:     locationDAO.Lat,
		Long:    locationDAO.Long,
	}

	return locationDTO
}
