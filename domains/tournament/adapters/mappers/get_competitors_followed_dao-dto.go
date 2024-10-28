package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/follower/dao"
)

func GetCompetitorsFollowedDAOtoDTO(competitorsDAO []*dao.GetCompetitorFollowedDAORes) []*dto.GetCompetitorFollowedDTORes {
	competitorsDTO := make([]*dto.GetCompetitorFollowedDTORes, len(competitorsDAO))

	for i, competitorDAO := range competitorsDAO {
		competitorUsers := make([]*dto.GetUserCompetitorFollowedDTORes, len(competitorDAO.Users))
		competitorGuestUsers := make([]*dto.GetUserCompetitorFollowedDTORes, len(competitorDAO.GuestUsers))

		competitorsDTO[i] = &dto.GetCompetitorFollowedDTORes{
			ID:         competitorDAO.ID.Hex(),
			CurrentPosition: competitorDAO.CurrentPosition,
			Users:      competitorUsers,
			GuestUsers: competitorGuestUsers,
		}

		for j, competitorUserDAO := range competitorsDAO[i].Users {
			competitorUsers[j] = &dto.GetUserCompetitorFollowedDTORes{
				ID:        competitorUserDAO.ID.Hex(),
				FirstName: competitorUserDAO.FirstName,
				LastName:  competitorUserDAO.LastName,
				Image:     competitorUserDAO.Image,
				Username:  competitorUserDAO.Username,
			}
		}

		for j, competitorGuestUserDAO := range competitorsDAO[i].GuestUsers {
			competitorGuestUsers[j] = &dto.GetUserCompetitorFollowedDTORes{
				ID:        competitorGuestUserDAO.ID.Hex(),
				FirstName: competitorGuestUserDAO.FirstName,
				LastName:  competitorGuestUserDAO.LastName,
				Image:     competitorGuestUserDAO.Image,
				Username:  competitorGuestUserDAO.Username,
			}
		}
	}

	return competitorsDTO
}
