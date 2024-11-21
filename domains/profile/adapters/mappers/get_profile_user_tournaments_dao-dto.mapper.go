package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	competitor_user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/competitor_user/dao"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/tournament_registration/dao"
	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
)

// Maps a GetProfileUserTournamentsDAORes to GetProfileUserTournamentsDTORes
func GetProfileUserTournamentsDAOtoDTO(profileTournamentsDAO *competitor_user_dao.GetProfileUserTournamentsDAORes) *dto.GetProfileUserTournamentsDTORes {
	if profileTournamentsDAO == nil {
		return nil
	}

	// Map tournaments
	tournamentsDTO := make([]*dto.GetProfileUserTournamentDTORes, len(profileTournamentsDAO.Tournaments))
	for i, tournament := range profileTournamentsDAO.Tournaments {
		tournamentsDTO[i] = &dto.GetProfileUserTournamentDTORes{
			ID:        tournament.ID.Hex(),
			Name:      tournament.Name,
			Location:  mapLocationDAOtoDTO(tournament.Location),
			Organizer: mapOrganizerDAOtoDTO(tournament.Organizer),
			Matches:   mapMatchesDAOtoDTO(tournament.Matches),
		}
	}

	return &dto.GetProfileUserTournamentsDTORes{
		Tournaments: tournamentsDTO,
		Total:       profileTournamentsDAO.Total,
	}
}

// Helper function to map GetLocationByIDDAORes to GetLocationByIDDTORes
func mapLocationDAOtoDTO(locationDAO *location_dao.GetLocationByIDDAORes) *dto.GetLocationByIDDTORes {
	if locationDAO == nil {
		return nil
	}
	return &dto.GetLocationByIDDTORes{
		State:   locationDAO.State,
		Country: locationDAO.Country,
		City:    locationDAO.City,
		Lat:     locationDAO.Lat,
		Long:    locationDAO.Long,
	}
}

// Helper function to map GetUserTournamentsOrganizerDAO to GetUserTournamentsOrganizerDTO
func mapOrganizerDAOtoDTO(organizerDAO *competitor_user_dao.GetUserTournamentsOrganizerDAO) *dto.GetUserTournamentsOrganizerDTO {
	if organizerDAO == nil {
		return nil
	}
	return &dto.GetUserTournamentsOrganizerDTO{
		ID:        organizerDAO.ID.Hex(),
		FirstName: organizerDAO.FirstName,
		LastName:  organizerDAO.LastName,
	}
}

// Helper function to map matches
func mapMatchesDAOtoDTO(matchesDAO []*competitor_user_dao.GetProfileUserTournamentMatchDAORes) []*dto.GetProfileUserTournamentMatchDTORes {
	matchesDTO := make([]*dto.GetProfileUserTournamentMatchDTORes, len(matchesDAO))
	for i, match := range matchesDAO {

		var winner *string
		if match.Winner != nil {
			winnerToString := match.Winner.Hex()
			winner = &winnerToString
		}
		matchesDTO[i] = &dto.GetProfileUserTournamentMatchDTORes{
			ID:          match.ID.Hex(),
			Result:      match.Result,
			Winner:      winner,
			Date:        match.Date,
			Round:       mapRoundDAOtoDTO(match.Round),
			Competitors: mapCompetitorsDAOtoDTO(match.Competitors),
		}
	}
	return matchesDTO
}

// Helper function to map round
func mapRoundDAOtoDTO(roundDAO *competitor_user_dao.GetProfileUserTournamentRoundDAORes) *dto.GetProfileUserTournamentRoundDTORes {
	if roundDAO == nil {
		return nil
	}
	return &dto.GetProfileUserTournamentRoundDTORes{
		ID:    roundDAO.ID.Hex(),
		Round: roundDAO.Round,
	}
}

// Helper function to map competitors
func mapCompetitorsDAOtoDTO(competitorsDAO []*dao.GetCompetitorsInTournamentDAORes) []*dto.GetCompetitorsInTournamentDTORes {
	competitorsDTO := make([]*dto.GetCompetitorsInTournamentDTORes, len(competitorsDAO))
	for i, competitor := range competitorsDAO {
		var competitorID *string
		if competitor.CompetitorID != nil {
			competitorIDToString := competitor.CompetitorID.Hex()
			competitorID = &competitorIDToString
		}
		competitorsDTO[i] = &dto.GetCompetitorsInTournamentDTORes{
			CompetitorID:    competitorID,
			CurrentPosition: competitor.CurrentPosition,
			Users:           mapUserDAOtoDTO(competitor.Users),
			GuestUsers:      mapUserDAOtoDTO(competitor.GuestUsers),
		}
	}
	return competitorsDTO
}

// Helper function to map users
func mapUserDAOtoDTO(usersDAO []*dao.GetCompetitorsInTournamentUserDAORes) []*dto.GetProfileInfoInCategoryUsersDTORes {
	usersDTO := make([]*dto.GetProfileInfoInCategoryUsersDTORes, len(usersDAO))
	for i, user := range usersDAO {
		usersDTO[i] = &dto.GetProfileInfoInCategoryUsersDTORes{
			ID:        user.ID.Hex(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Image:     user.Image,
		}
	}
	return usersDTO
}
