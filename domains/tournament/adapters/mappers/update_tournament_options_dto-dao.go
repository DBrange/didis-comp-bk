package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/tournament/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func UpdateTournamentOptionsDTOtoDAO(tournamentDTO *dto.UpdateTournamentOptionsDTOReq, convert utils.ConvertToObjectIDFunc) (*dao.UpdateTournamentOptionsDAOReq, error) {
	tournamentDAO := &dao.UpdateTournamentOptionsDAOReq{}

	if tournamentDTO.DoubleEliminationID != nil {
		doubleEliminationOID, err := convert(*tournamentDTO.DoubleEliminationID)
		if err != nil {
			return nil, err
		}

		tournamentDAO.DoubleEliminationID = doubleEliminationOID
	}

	if tournamentDTO.Pots != nil {
		potsOID, err := ConvertToObjectIDs(tournamentDTO.Pots, convert)
		if err != nil {
			return nil, err
		}

		tournamentDAO.Pots = potsOID
	}

	if tournamentDTO.Groups != nil {
		groupsOID, err := ConvertToObjectIDs(tournamentDTO.Groups, convert)
		if err != nil {
			return nil, err
		}

		tournamentDAO.Groups = groupsOID
	}

	if tournamentDTO.Rounds != nil {
		roundsOID, err := ConvertToObjectIDs(tournamentDTO.Rounds, convert)
		if err != nil {
			return nil, err
		}
		
		tournamentDAO.Rounds = roundsOID
	}

	return tournamentDAO, nil
}

func ConvertToObjectIDs(IDs *[]string, convert utils.ConvertToObjectIDFunc) (*[]primitive.ObjectID, error) {
	OIDs := make([]primitive.ObjectID, len(*IDs))
	for i, id := range *IDs {
		OID, err := convert(id)
		if err != nil {
			return nil, err
		}
		OIDs[i] = *OID
	}

	return &OIDs, nil
}
