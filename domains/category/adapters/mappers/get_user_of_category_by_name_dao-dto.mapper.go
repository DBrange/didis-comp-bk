package mappers

import (
	"fmt"

	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/category_registration/dao"
)

func GetUsersOfCategoryByNameDAOtoDTO(competitorsDAO []*dao.GetCompetitorsOfCategoryCompetitorDAORes) []*dto.GetCompetitorsOfCategoryCompetitorDTORes {

	if len(competitorsDAO) == 0 {
		return []*dto.GetCompetitorsOfCategoryCompetitorDTORes{}
	}

	competitorsDTO := make([]*dto.GetCompetitorsOfCategoryCompetitorDTORes, len(competitorsDAO))

	for i, competitor := range competitorsDAO {
		competitorDTO := &dto.GetCompetitorsOfCategoryCompetitorDTORes{
			ID:                  competitor.ID.Hex(),
			CurrentPosition:     competitor.CurrentPosition,
			Points:              competitor.Points,
			RegisteredPositions: MapGetUsersOfCategoryByNameRegisteredPositionDAOToDTO(competitor.RegisteredPositions),
			Users:               participantsDAOtoDTO(competitor.Users),
			GuestUsers:          participantsDAOtoDTO(competitor.GuestUsers),
		}

		competitorsDTO[i] = competitorDTO
	}

	return competitorsDTO
}

func participantsDAOtoDTO(participantsDAO []*dao.GetCompetitorsOfCategoryUserDAORes) []*dto.GetCompetitorsOfCategoryUserDTORes {
	paticipantsDTO := make([]*dto.GetCompetitorsOfCategoryUserDTORes, len(participantsDAO))

	for i, paticipant := range participantsDAO {

		paticipantsDTO[i] = &dto.GetCompetitorsOfCategoryUserDTORes{
			ID:        paticipant.ID.Hex(),
			FirstName: paticipant.FirstName,
			LastName:  paticipant.LastName,
			Image:     paticipant.Image,
			Username:     paticipant.Username,
		}
	}

	return paticipantsDTO

}

func MapGetUsersOfCategoryByNameRegisteredPositionDAOToDTO(dao []dao.RegistedPositionDAORes) []dto.RegistedPositionDTORes {
	registeredPositions := make([]dto.RegistedPositionDTORes, len(dao))
fmt.Printf("aaaaaa  %+v", dao)
	for i, rp := range dao {
		registeredPositions[i] = dto.RegistedPositionDTORes{
			Date:     rp.Date,
			Position: rp.Position,
		}
	}

	return registeredPositions
}
