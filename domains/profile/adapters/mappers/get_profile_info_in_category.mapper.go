package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/category_registration/dao"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProfileInfoInCategoryDAOtoDTO(profileDAO *dao.GetProfileInfoInCategoryDAORes) *dto.GetProfileInfoInCategoryDTORes {
	profileDTO := &dto.GetProfileInfoInCategoryDTORes{
		ID:              profileDAO.ID.Hex(),
		Points:          profileDAO.Points,
		CurrentPosition: profileDAO.CurrentPosition,
		Users:           getProfileInfoInCategoryUserDAOtoDTO(profileDAO.Users),
		GuestUsers:      getProfileInfoInCategoryUserDAOtoDTO(profileDAO.GuestUsers),
		CompetitorStats: getProfileInfoInCategoryStatsDAOtoDTO(profileDAO.CompetitorStats),
	}

	return profileDTO
}

func getProfileInfoInCategoryUserDAOtoDTO(profileUsersDAO []*dao.GetProfileInfoInCategoryUsersDAORes) []*dto.GetProfileInfoInCategoryUsersDTORes {
	if len(profileUsersDAO) == 0{
		return []*dto.GetProfileInfoInCategoryUsersDTORes{}
	}

	profileUsersDTO := make([]*dto.GetProfileInfoInCategoryUsersDTORes, len(profileUsersDAO))

	for i, user := range profileUsersDAO {
		profileUsersDTO[i] = &dto.GetProfileInfoInCategoryUsersDTORes{
			ID:        user.ID.Hex(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Image:     user.Image,
		}

	}

	return profileUsersDTO
}

func getProfileInfoInCategoryStatsDAOtoDTO(profileStatsDAO *dao.GetProfileInfoInCategoryStatsDAORes) *dto.GetProfileInfoInCategoryStatsDTORes {
	profileStatsDTO := dto.GetProfileInfoInCategoryStatsDTORes{
		ID:             profileStatsDAO.ID.Hex(),
		TotalWins:      profileStatsDAO.TotalWins,
		TotalLosses:    profileStatsDAO.TotalLosses,
		MoneyEarned:    profileStatsDAO.MoneyEarned,
		TournamentsWon: tournamentsWonMapper(profileStatsDAO.TournamentsWon),
	}

	return &profileStatsDTO
}

func tournamentsWonMapper(tournaments []*primitive.ObjectID) []string {
	tournamentsStr := make([]string, len(tournaments))
	for i, t := range tournaments {
		tournamentsStr[i] = t.Hex()
	}

	return tournamentsStr
}
