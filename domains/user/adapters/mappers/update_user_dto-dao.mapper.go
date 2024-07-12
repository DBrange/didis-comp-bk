package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
	"github.com/DBrange/didis-comp-bk/domains/user/models/dto"
)

func UpdateUserDTOtoDAO(newUserInfoDTO *dto.UpdateUserDTOReq) *dao.UpdateUserDAOReq{
	newUserInfoDAO := &dao.UpdateUserDAOReq{
		FirstName: newUserInfoDTO.FirstName,
		LastName: newUserInfoDTO.LastName,
		Username: newUserInfoDTO.Username,
		Image: newUserInfoDTO.Image,
		Phone: newUserInfoDTO.Phone,
	} 

	return newUserInfoDAO
}