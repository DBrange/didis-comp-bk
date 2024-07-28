package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/role/dao"
)

func CreateRoleDAOtoDTO(roleDAO *dao.GetRoleDAOByID) *dto.GetRoleDTOByID {
	roleDTO := &dto.GetRoleDTOByID{
		ID:       roleDAO.ID.Hex(),
		Name:     roleDAO.Name,
		RoleType: roleDAO.RoleType,
	}

	return roleDTO
}
