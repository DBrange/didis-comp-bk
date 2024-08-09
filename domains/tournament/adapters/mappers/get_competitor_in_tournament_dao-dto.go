package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/tournament_registration/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

func GetCompetitorsInTournamentDAOtoDTO(competitorsDAO []*dao.GetCompetitorsInTournamentDAORes) []*dto.GetCompetitorsInTournamentDTORes{
	competitorsDTO := make([]*dto.GetCompetitorsInTournamentDTORes, len(competitorsDAO))
	
	for i, competitor := range competitorsDAO{
		competitorsDTO[i] = &dto.GetCompetitorsInTournamentDTORes{
			CompetitorID: competitor.CompetitorID.Hex(),
			CurrentPosition: competitor.CurrentPosition,
			Users: GetCompetitorsInTournamentUserDAOtoDTO(competitor.Users),
			GuestUsers: GetCompetitorsInTournamentUserDAOtoDTO(competitor.GuestUsers),
		}
	}

	return competitorsDTO
}

func GetCompetitorsInTournamentUserDAOtoDTO(usersDAO []*dao.GetCompetitorsInTournamentUserDAORes) []*dto.GetCompetitorsInTournamentUserDTORes{
	usersDTO := make([]*dto.GetCompetitorsInTournamentUserDTORes, len(usersDAO))

	for i, user := range usersDAO{
		usersDTO[i] = &dto.GetCompetitorsInTournamentUserDTORes{
			ID: user.ID.Hex(),
			FirstName: user.FirstName,
			LastName: user.LastName,
			Image: user.Image,
		}
	}

	return usersDTO
}