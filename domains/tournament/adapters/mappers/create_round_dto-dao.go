package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/round/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

func CreateRoundDTOtoDAO(roundDTO *dto.CreateRoundDTOReq, convert utils.ConvertToObjectIDFunc) (*dao.CreateRoundDAOReq, error) {
	tournamentDAO, err := convert(roundDTO.TournamentID)
	if err != nil {
		return nil, err
	}

	roundDAO := &dao.CreateRoundDAOReq{
		TournamentID: *tournamentDAO,
		Name:         roundDTO.Name,
		TotalPrize:   roundDTO.TotalPrize,
	}

	return roundDAO, nil
}
