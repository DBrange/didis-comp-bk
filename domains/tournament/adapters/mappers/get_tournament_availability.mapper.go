package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/tournament/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

func GetTournamentAvailavilityDAOtoDTO(tournamentAvilabilityDAO *dao.TournamentAvailabilityDAO) *dto.TournamentAvailabilityDTO {
	tournamentAvilabilityDTO := &dto.TournamentAvailabilityDTO{
		AvailableCourts: tournamentAvilabilityDAO.AvailableCourts,
		AverageHours:    tournamentAvilabilityDAO.AverageHours,
	}

	return tournamentAvilabilityDTO
}
