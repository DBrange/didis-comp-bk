package dto

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type GetRoleDTOByID struct {
	ID       string           `json:"_id"`
	Name     models.ROLE      `json:"name"`
	RoleType models.ROLE_TYPE `json:"role_type"`
}
