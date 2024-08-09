package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/tournament/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

func UpdateTournamentOptionsDTOtoDAO(tournamentDTO *dto.UpdateTournamentOptionsDTOReq, convert utils.ConvertToObjectIDFunc) (*dao.UpdateTournamentOptionsDAOReq, error) {
	tournamentDAO := &dao.UpdateTournamentOptionsDAOReq{}

	if tournamentDTO.DoubleEliminationID != nil && *tournamentDTO.DoubleEliminationID != "" {
		doubleEliminationOID, err := convert(*tournamentDTO.DoubleEliminationID)
		if err != nil {
			return nil, err
		}

		tournamentDAO.DoubleEliminationID = doubleEliminationOID
	}

	if tournamentDTO.Pots != nil {
		potsOID, err := utils.ConvertToObjectIDs(tournamentDTO.Pots, convert)
		if err != nil {
			return nil, err
		}

		tournamentDAO.Pots = potsOID
	}

	if tournamentDTO.Groups != nil {
		groupsOID, err := utils.ConvertToObjectIDs(tournamentDTO.Groups, convert)
		if err != nil {
			return nil, err
		}

		tournamentDAO.Groups = groupsOID
	}

	if tournamentDTO.Matches != nil {
		matchesOID, err := utils.ConvertToObjectIDs(tournamentDTO.Matches, convert)
		if err != nil {
			return nil, err
		}

		tournamentDAO.Matches = matchesOID
	}

	if tournamentDTO.Rounds != nil {
		roundsOID, err := utils.ConvertToObjectIDs(tournamentDTO.Rounds, convert)
		if err != nil {
			return nil, err
		}

		tournamentDAO.Rounds = roundsOID
	}

	return tournamentDAO, nil
}
