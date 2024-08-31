package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/double_elimination/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

func GetDoubleElimInfoToFinaliseItDAOtoDTO(doubleElimInfoDAO *dao.GetDoubleElimInfoToFinaliseItDAORes) *dto.GetDoubleElimInfoToFinaliseItDTORes {
	doubleElimInfoDTO := &dto.GetDoubleElimInfoToFinaliseItDTORes{
		TotalPrize: doubleElimInfoDAO.TotalPrize,
		Points:     doubleElimInfoDAO.Points,
	}

	return doubleElimInfoDTO
}
