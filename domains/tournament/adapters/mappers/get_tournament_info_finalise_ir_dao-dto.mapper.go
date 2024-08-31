package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/tournament/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

func GetTournamentInfoToFinaliseItDAOtoDTO(tournamentInfoDAO *dao.GetTournamentInfoToFinaliseItDAORes) *dto.GetTournamentInfoToFinaliseItDTORes {
	var categoryID string
	
	if tournamentInfoDAO.CategoryID != nil{
		categoryID = tournamentInfoDAO.CategoryID.Hex()
	}
	
	tournamentInfoDTO := &dto.GetTournamentInfoToFinaliseItDTORes{
		CategoryID: categoryID,
		TotalPrize: tournamentInfoDAO.TotalPrize,
		Points:     tournamentInfoDAO.Points,
	}

	return tournamentInfoDTO
}
