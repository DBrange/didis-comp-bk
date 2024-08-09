package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/follower/dao"
)

func GetCompetitorsFollowedDAOtoDTO(competitorsDAO []*dao.GetCompetitorFollowedDAORes) []*dto.GetCompetitorFollowedDTORes {
	competitorsDTO := make([]*dto.GetCompetitorFollowedDTORes, len(competitorsDAO))

	for i, competitorDAO := range competitorsDAO {
		competitorUsers := make([]*dto.GetUserCompetitorFollowedDTORes, len(competitorDAO.Users))

		competitorsDTO[i] = &dto.GetCompetitorFollowedDTORes{
			ID:    competitorDAO.ID.Hex(),
			Users: competitorUsers,
		}

		for j, competitorUserDAO := range competitorsDAO[i].Users {
			competitorUsers[j] = &dto.GetUserCompetitorFollowedDTORes{
				ID:        competitorUserDAO.ID.Hex(),
				FirstName: competitorUserDAO.FirstName,
				LastName:  competitorUserDAO.LastName,
				Image:     competitorUserDAO.Image,
			}

		}
	}

	return competitorsDTO
}
