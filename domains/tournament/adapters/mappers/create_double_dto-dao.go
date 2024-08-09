package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/double/dao"
)

func CreateDoubleDTOtoDAO(doubleDTO *dto.CreateDoubleDTOReq) *dao.CreateDoubleDAOReq {
	return &dao.CreateDoubleDAOReq{}
}
