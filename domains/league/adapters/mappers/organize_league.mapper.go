package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/league/models/dto"
)

func OrganizeLeagueMapper(organizerLeague *dto.OrganizeLeagueDTOReq) *dto.CreateLeagueDTOReq {
	leagueDTO := &dto.CreateLeagueDTOReq{
		Name:              organizerLeague.Name,
		Genre:             organizerLeague.Genre,
		TotalParticipants: organizerLeague.TotalParticipants,
		RangeMovement:     organizerLeague.RangeMovement,
		Sport:             organizerLeague.Sport,
	}

	return leagueDTO
}
