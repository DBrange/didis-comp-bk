package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/tournament/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

func GetTournamentInfoToFinaliseItDAOtoDTO(tournamentInfoDAO *dao.GetTournamentInfoToFinaliseItDAORes) *dto.GetTournamentInfoToFinaliseItDTORes {
	tournamentInfoDTO := &dto.GetTournamentInfoToFinaliseItDTORes{
		CategoryID: tournamentInfoDAO.CategoryID.Hex(),
		TotalPrize: tournamentInfoDAO.TotalPrize,
		Points:     tournamentInfoDAO.Points,
	}

	return tournamentInfoDTO
}
