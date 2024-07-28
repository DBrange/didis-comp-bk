package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
)

func GetUserRolesDTOtoDAO(loginDTO *dto.LoginDTOReq) *dao.LoginDAOReq{
	loginDAO := &dao.LoginDAOReq{
		Username: loginDTO.Username,
		Password: loginDTO.Password,
	}

	return loginDAO
}