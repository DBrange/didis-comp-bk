package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
)

func UpdateProfileDTOtoDAO(newProfileInfoDTO *dto.UpdateProfileDTOReq) *dao.UpdateUserDAOReq {
	newProfileInfoDAO := &dao.UpdateUserDAOReq{
		FirstName: newProfileInfoDTO.FirstName,
		LastName:  newProfileInfoDTO.LastName,
		Username:  newProfileInfoDTO.Username,
		Image:     newProfileInfoDTO.Image,
		Phone:     newProfileInfoDTO.Phone,
	}

	return newProfileInfoDAO
}
