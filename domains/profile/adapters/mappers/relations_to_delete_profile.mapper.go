package mappers

import (
	profile_dto "github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
)

func ProfileRelationsToDeleteDAOtoDTO(profileRelationsToDeleteDAO *dao.UserRelationsToDeleteDAOReq) *profile_dto.ProfileRelationsToDeleteDTO {
	profileRelationsToDeleteDTO := &profile_dto.ProfileRelationsToDeleteDTO{
		LocationID:     profileRelationsToDeleteDAO.LocationID,
	}

	return profileRelationsToDeleteDTO
}
