package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/category_registration/dao"
)

func GetUsersOfCategoryByNameDAOtoDTO(competitorsDAO []*dao.GetCompetitorsOfCategoryDAORes) []*dto.GetCompetitorsOfCategoryDTORes {

	if len(competitorsDAO) == 0 {
		return []*dto.GetCompetitorsOfCategoryDTORes{}
	}

	competitorsDTO := make([]*dto.GetCompetitorsOfCategoryDTORes, len(competitorsDAO))

	for i, competitor := range competitorsDAO {
		competitorDTO := &dto.GetCompetitorsOfCategoryDTORes{
			ID:                  competitor.ID.Hex(),
			CurrentPosition:     competitor.CurrentPosition,
			Points:              competitor.Points,
			RegisteredPositions: competitor.RegisteredPositions,
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
		}
	}

	return paticipantsDTO

}
