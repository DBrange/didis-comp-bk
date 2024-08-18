package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/match/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

func GetPositionsBracketMatchDAOtoDTO(positionsDAO []*dao.GetPositionsBracketMatchDAORes) []*dto.GetPositionsBracketMatchDTORes {
	positionsDTO := make([]*dto.GetPositionsBracketMatchDTORes, len(positionsDAO))

	for i, positionDAO := range positionsDAO {
		positionsDTO[i] = &dto.GetPositionsBracketMatchDTORes{
			ID:            positionDAO.ID.Hex(),
			PositionMatch: positionDAO.PositionMatch,
			Competitors:   GetPositionsBracketMatchCompetitorDAOtoDTO(positionDAO.Competitors),
		}
	}

	return positionsDTO
}
func GetPositionsBracketMatchCompetitorDAOtoDTO(positionsCompDAO []*dao.GetPositionsBracketMatchCompetitorDAORes) []*dto.GetPositionsBracketMatchCompetitorDTORes {
	positionsCompDTO := make([]*dto.GetPositionsBracketMatchCompetitorDTORes, len(positionsCompDAO))

	for i, positionCompDAO := range positionsCompDAO {
		positionsCompDTO[i] = &dto.GetPositionsBracketMatchCompetitorDTORes{
			ID:              positionCompDAO.ID.Hex(),
			Position:        positionCompDAO.Position,
			CurrentPosition: positionCompDAO.CurrentPosition,
		}
	}

	return positionsCompDTO
}
