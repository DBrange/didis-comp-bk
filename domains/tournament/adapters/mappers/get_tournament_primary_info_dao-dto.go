package mappers

import (
	location_dto "github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	competitor_user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/competitor_user/dao"
	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/tournament/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

func GetTournamentPrimaryInfoDAOtoDTO(tournamentDAO *dao.GetTournamentPrimaryInfoDAORes) *dto.GetTournamentPrimaryInfoDTORes {
	return &dto.GetTournamentPrimaryInfoDTORes{
		ID:               tournamentDAO.ID.Hex(),
		Name:             tournamentDAO.Name,
		StartDate:        tournamentDAO.StartDate,
		FinishDate:       tournamentDAO.FinishDate,
		Points:           tournamentDAO.Points,
		Image:            tournamentDAO.Image,
		TotalPrize:       tournamentDAO.TotalPrize,
		TotalCompetitors: tournamentDAO.TotalCompetitors,
		MaxCapacity:      tournamentDAO.MaxCapacity,
		AverageScore:     tournamentDAO.AverageScore,
		Genre:            tournamentDAO.Genre,
		Sport:            tournamentDAO.Sport,
		CompetitorType:   tournamentDAO.CompetitorType,
		Surface:          tournamentDAO.Surface,
		Availability: &dto.TournamentAvailabilityDTO{
			AvailableCourts: tournamentDAO.Availability.AvailableCourts,
			AverageHours:    tournamentDAO.Availability.AverageHours,
		},
		Rounds:    GetTournamentPrimaryInfoRoundDAOtoDTO(tournamentDAO.Rounds),
		Location:  GetTournamentPrimaryInfoLocationDAOtoDTO(tournamentDAO.Location),
		Organizer: GetTournamentPrimaryInfoOrganizerDAOtoDTO(tournamentDAO.Organizer),
		Category:  GetTournamentPrimaryInfoCategoryDAOtoDTO(tournamentDAO.Category),
	}
}

func GetTournamentPrimaryInfoRoundDAOtoDTO(roundsDAO []*dao.GetTournamentPrimaryInforRoundDAORes) []*dto.GetTournamentPrimaryInforRoundDTORes {
	roundsDTO := make([]*dto.GetTournamentPrimaryInforRoundDTORes, len(roundsDAO))

	for i, roundDAO := range roundsDAO {

		roundsDTO[i] = &dto.GetTournamentPrimaryInforRoundDTORes{
			ID:    roundDAO.ID.Hex(),
			Round: roundDAO.Round,
		}
	}

	return roundsDTO
}

func GetTournamentPrimaryInfoLocationDAOtoDTO(locationDAO *location_dao.GetLocationByIDDAORes) *location_dto.GetLocationByIDDTORes {
	if locationDAO == nil { // Check if locationDAO is nil
		return nil
	}

	locationDTO := &location_dto.GetLocationByIDDTORes{
		ID:      locationDAO.ID.Hex(),
		State:   locationDAO.State,
		Country: locationDAO.Country,
		City:    locationDAO.City,
		Lat:     locationDAO.Lat,
		Long:    locationDAO.Long,
	}

	return locationDTO
}

func GetTournamentPrimaryInfoOrganizerDAOtoDTO(organizerDAO *competitor_user_dao.GetUserTournamentsOrganizerDAO) *dto.GetUserTournamentOrganizerDTORes {
	if organizerDAO == nil { // Check if organizerDAO is nil
		return nil
	}

	return &dto.GetUserTournamentOrganizerDTORes{
		ID:        organizerDAO.ID.Hex(),
		FirstName: organizerDAO.FirstName,
		LastName:  organizerDAO.LastName,
	}
}

func GetTournamentPrimaryInfoCategoryDAOtoDTO(categoryDAO *dao.GetTournamentPrimaryInfoCategoryDAORes) *dto.GetTournamentPrimaryInfoCategoryDTORes {
	if categoryDAO == nil { // Check if categoryDAO is nil
		return nil
	}

	return &dto.GetTournamentPrimaryInfoCategoryDTORes{
		ID:   categoryDAO.ID.Hex(),
		Name: categoryDAO.Name,
	}
}
