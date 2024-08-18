package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/pot/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

func GetTournamentPotPositionsDAOtoDTO(potDAOs []*dao.PotOrGroupPositionDAORes) []*dto.PotOrGroupPositionDTORes {
	potDTOs := make([]*dto.PotOrGroupPositionDTORes, len(potDAOs))

	for i, potDAO := range potDAOs {
		potDTOs[i] = &dto.PotOrGroupPositionDTORes{
			ID:       potDAO.ID.Hex(),
			Position: potDAO.Position,
		}
	}

	return potDTOs
}
