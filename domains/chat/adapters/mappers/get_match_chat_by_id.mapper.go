package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/chat/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/chat/dao"
)

func GetMatchChatByIDDAOtoDTO(chatDAO *dao.GetMatchChatByIDDAORes) *dto.GetMatchChatByIDDTORes {
	chatDTO := &dto.GetMatchChatByIDDTORes{
		ID:                 chatDAO.ID.Hex(),
		AvailabilityStatus: chatDAO.AvailabilityStatus,
		MatchID:            chatDAO.MatchID.Hex(),
		Users:               GetMatchChatByIDUsersDAOtoDTO(chatDAO.Users),
		Competitors:        GetMatchChatByIDCompetitorDAOtoDTO(chatDAO.Competitors),
	}

	return chatDTO
}

func GetMatchChatByIDUserDAOtoDTO(userDAO *dao.GetMatchChatByIDUserdDAORes) *dto.GetMatchChatByIDUserdDTORes {
	userDTO := &dto.GetMatchChatByIDUserdDTORes{
		ID:        userDAO.ID.Hex(),
		FirstName: userDAO.FirstName,
		LastName:  userDAO.LastName,
		Image:     userDAO.Image,
	}

	return userDTO
}

func GetMatchChatByIDCompetitorDAOtoDTO(competitorDAOs []*dao.GetMatchChatByIDCompetitordDAORes) []*dto.GetMatchChatByIDCompetitordDTORes {
	competitorDTOs := make([]*dto.GetMatchChatByIDCompetitordDTORes, len(competitorDAOs))

	for i, competitorDAO := range competitorDAOs {
		competitorDTOs[i] = &dto.GetMatchChatByIDCompetitordDTORes{
			ID:                 competitorDAO.ID.Hex(),
			AvailabilityStatus: competitorDAO.AvailabilityStatus,
			Users:              GetMatchChatByIDUsersDAOtoDTO(competitorDAO.Users),
		}
	}

	return competitorDTOs
}

func GetMatchChatByIDUsersDAOtoDTO(userDAOs []*dao.GetMatchChatByIDUserdDAORes) []*dto.GetMatchChatByIDUserdDTORes {
	userDTOs := make([]*dto.GetMatchChatByIDUserdDTORes, len(userDAOs))

	for i, userDAO := range userDAOs {
		userDTOs[i] = GetMatchChatByIDUserDAOtoDTO(userDAO)
	}

	return userDTOs
}
