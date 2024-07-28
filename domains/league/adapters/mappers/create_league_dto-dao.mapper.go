package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/league/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/league/dao"
)

func CreateLeagueDTOtoDAO(leagueDTO *dto.CreateLeagueDTOReq) (*dao.CreateLeagueDAOReq) {
	leagueDAO := &dao.CreateLeagueDAOReq{
		Name:              leagueDTO.Name,
		Genre:             leagueDTO.Genre,
		TotalParticipants: leagueDTO.TotalParticipants,
		RangeMovement:     leagueDTO.RangeMovement,
		AverageScore:      leagueDTO.AverageScore,
		Sport:             leagueDTO.Sport,
	}

	return leagueDAO
}
