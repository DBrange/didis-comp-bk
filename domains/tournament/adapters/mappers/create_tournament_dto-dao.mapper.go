package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/tournament/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

func CreateTournamentDTOtoDAO(tournamentDTO *dto.CreateTournamentDTOReq) *dao.CreateTournamentDAOReq {
	tournamentDAO := &dao.CreateTournamentDAOReq{
		Name:             tournamentDTO.Name,
		Points:           tournamentDTO.Points,
		TotalPrize:       tournamentDTO.TotalPrize,
		TotalCompetitors: tournamentDTO.TotalCompetitors,
		MaxCapacity:      tournamentDTO.MaxCapacity,
		Genre:            tournamentDTO.Genre,
		Sport:            tournamentDTO.Sport,
		Surface:          tournamentDTO.Surface,
	}

	return tournamentDAO
}
