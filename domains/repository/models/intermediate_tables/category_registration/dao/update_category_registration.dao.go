package dao

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/common"
)

type UpdateCategoryRegistrationDAOReq struct {
	Points               *int                      `bson:"points,omitempty"`
	RegisteredPositions  *[]RegistedPositionDAORes `bson:"registered_positions,omitempty"`
	CurrentPosition      *int                      `bson:"current_position,omitempty"`
	common.UpdateBaseDAO `bson:",inline"`
}
