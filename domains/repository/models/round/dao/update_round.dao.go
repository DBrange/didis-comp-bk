package dao

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/common"
)

type UpdateRoundDAOReq struct {
	TotalPrize           float64 `bson:"total_prize"`
	common.UpdateBaseDAO `bson:",inline"`
}
