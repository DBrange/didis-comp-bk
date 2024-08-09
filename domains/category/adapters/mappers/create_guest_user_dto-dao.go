package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/guest_user/dao"
	dto "github.com/DBrange/didis-comp-bk/domains/category/models/dto"
)

func CreateGuestUserDTOtoDAO(guestUserDTO *dto.CreateGuestUserDTOReq) *dao.CreateGuestUserDAOReq {
	guestUserDAO := &dao.CreateGuestUserDAOReq{
		FirstName: guestUserDTO.FirstName,
		LastName:  guestUserDTO.LastName,
		Email:     guestUserDTO.Email,
		Image:     guestUserDTO.Image,
		Genre:     guestUserDTO.Genre,
	}

	return guestUserDAO
}
