package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/single/dao"
)

func CreateSingleDTOtoDAO(singleDTO *dto.CreateSingleDTOReq) *dao.CreateSingleDAOReq{
	return &dao.CreateSingleDAOReq{}
}