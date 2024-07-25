package dao

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/common"
)

type UpdateLeagueRegistrationDAOReq struct {
	Points              *int   `bson:"points,omitempty"`
	RegisteredPositions *[]int `bson:"registered_positions,omitempty"`
	CurrentPosition     *int   `bson:"registered_positions,omitempty"`
	common.UpdateBaseDAO `bson:",inline"`
}
