package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/league/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/league/dao"
)

func OrganizeLeagueMapper(leagueInfoDTO *dto.OrganizeLeagueDTOReq) *dao.CreateLeagueDAOReq {
	leagueInfoDAO := &dao.CreateLeagueDAOReq{
		Name:              leagueInfoDTO.Name,
		Genre:             leagueInfoDTO.Genre,
		TotalParticipants: leagueInfoDTO.TotalParticipants,
		RangeMovement:     leagueInfoDTO.RangeMovement,
		Sport:             leagueInfoDTO.Sport,
	}

	return leagueInfoDAO
}
