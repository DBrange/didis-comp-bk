package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/follower/dao"
)

func GetUserFollowersDAOtoDTO(followersDAO *dao.GetUserFollowersDAORes) *dto.GetUserFollowersDTORes {
	usersDTO := make([]*dto.GetUserFollowersUserDTORes, len(followersDAO.Followers))

	for i, userDAO := range followersDAO.Followers {
		usersDTO[i] = &dto.GetUserFollowersUserDTORes{
			ID:        userDAO.ID.Hex(),
			FirstName: userDAO.FirstName,
			LastName:  userDAO.LastName,
			Image:     userDAO.Image,
			Username:  userDAO.Username,
		}
	}

	followersDTO := &dto.GetUserFollowersDTORes{
		LastCreatedAt: followersDAO.LastCreatedAt,
		Followers:     usersDTO,
		Total: followersDAO.Total,
	}

	return followersDTO

}
