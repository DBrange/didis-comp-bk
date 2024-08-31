package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/double_elimination/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

func CreateDoubleEliminationDTOtoDAO(doubleEliminationDTO *dto.CreateDoubleEliminationDTOReq, convert utils.ConvertToObjectIDFunc) (*dao.CreateDoubleEliminationDAOReq, error) {
	var doubleEliminationDAO dao.CreateDoubleEliminationDAOReq
	
	if doubleEliminationDTO.Matches != nil {
		matchesOID, err := utils.ConvertToObjectIDs(&doubleEliminationDTO.Matches, convert)
		if err != nil {
			return nil, err
		}

		doubleEliminationDAO.Matches = *matchesOID
	}

	if doubleEliminationDTO.Rounds != nil {
		roundsOID, err := utils.ConvertToObjectIDs(&doubleEliminationDTO.Rounds, convert)
		if err != nil {
			return nil, err
		}

		doubleEliminationDAO.Rounds = *roundsOID
	}

	doubleEliminationDAO.TotalPrize = doubleEliminationDTO.TotalPrize
	doubleEliminationDAO.Points = doubleEliminationDTO.Points

	return &doubleEliminationDAO,nil
}
