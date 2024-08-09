package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/match/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

func CreateMatchDTOtoDAO(matchDTO *dto.CreateMatchDTOReq, convert utils.ConvertToObjectIDFunc) (*dao.CreateMatchDAOReq, error) {
	roundOID, err := convert(matchDTO.RoundID)
	if err != nil {
		return nil, err
	}
	
	tournamentOID, err := convert(matchDTO.TournamentID)
	if err != nil {
		return nil, err
	}

	matchDAO := &dao.CreateMatchDAOReq{
		Sport:        matchDTO.Sport,
		RoundID:      roundOID,
		Result:       matchDTO.Result,
		TournamentID: tournamentOID,
	}

	return matchDAO, nil
}
