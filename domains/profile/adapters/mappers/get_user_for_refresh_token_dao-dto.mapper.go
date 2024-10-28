package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
)

func GetUserForRefreshTokenDAOtoDTO(userDAO *dao.GetUserForRefreshTokenDAO) *dto.GetUserForRefreshTokenDTO {
	userDTO := &dto.GetUserForRefreshTokenDTO{
		ID:        userDAO.ID,
		Roles:     userDAO.Roles,
	}

	return userDTO
}
