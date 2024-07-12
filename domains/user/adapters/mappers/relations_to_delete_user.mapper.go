package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
	user_dto "github.com/DBrange/didis-comp-bk/domains/user/models/dto"
)

func UserRelationsToDeleteDAOtoDTO(userRelationsToDeleteDAO *dao.UserRelationsToDeleteDAO) *user_dto.UserRelationsToDeleteDTO {
	userRelationsToDeleteDTO := &user_dto.UserRelationsToDeleteDTO{
		LocationID: userRelationsToDeleteDAO.LocationID,
		PaymentID:  userRelationsToDeleteDAO.PaymentID,
		ScheduleID: userRelationsToDeleteDAO.ScheduleID,
	}

	return userRelationsToDeleteDTO
}
