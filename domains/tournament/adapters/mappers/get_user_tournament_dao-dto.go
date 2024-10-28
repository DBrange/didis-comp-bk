package mappers

import (
	tournament_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/competitor_user/dao"
	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

func GetUserTournamentsDAOtoDTO(userTournametsDAO *tournament_dao.GetUserTournamentsDAORes) *dto.GetUserTournamentsDTORes {
return &dto.GetUserTournamentsDTORes{
	Tournaments: getUserTournamentsTournamentDAOtoDTO(userTournametsDAO.Tournaments),
	Total: userTournametsDAO.Total,
}
}


func getUserTournamentsTournamentDAOtoDTO(tournametsDAO []*tournament_dao.GetUserTournamentDAORes) []*dto.GetUserTournamentDTORes {
	tournametsDTO := make([]*dto.GetUserTournamentDTORes, len(tournametsDAO))

	for i, tournamentDAO := range tournametsDAO {
		tournametsDTO[i] = &dto.GetUserTournamentDTORes{
			ID:           tournamentDAO.ID.Hex(),
			Name:         tournamentDAO.Name,
			StartDate:    tournamentDAO.StartDate,
			FinishDate:   tournamentDAO.FinishDate,
			Points:       tournamentDAO.Points,
			Image:       tournamentDAO.Image,
			AverageScore: tournamentDAO.AverageScore,
			TotalPrize: tournamentDAO.TotalPrize,
			Location:     getUserTournamentsLocationDAOtoDTO(tournamentDAO.Location),
			Organizer:    getUserTournamentsOrganizerDAOtoDTO(tournamentDAO.Organizer),
		}
	}

	return tournametsDTO
}

func getUserTournamentsLocationDAOtoDTO(locationDAO *location_dao.GetLocationByIDDAORes) *dto.GetLocationByIDDTORes {
	return &dto.GetLocationByIDDTORes{
		ID:      locationDAO.ID.Hex(),
		State:   locationDAO.State,
		Country: locationDAO.Country,
		City:    locationDAO.City,
		Lat:     locationDAO.Lat,
		Long:    locationDAO.Long,
	}
}

func getUserTournamentsOrganizerDAOtoDTO(organizerDAO *tournament_dao.GetUserTournamentsOrganizerDAO) *dto.GetUserTournamentOrganizerDTORes {
	return &dto.GetUserTournamentOrganizerDTORes{
		ID:        organizerDAO.ID.Hex(),
		FirstName: organizerDAO.FirstName,
		LastName:  organizerDAO.LastName,
	}
}
