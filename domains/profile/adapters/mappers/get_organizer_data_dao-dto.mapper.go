package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/organizer/dao"
)

func GetOrganizerDataDAOtoDTO(organizerDAO *dao.GetOrganizerDataDAORes) *dto.GetOrganizerDataDTORes {
	return &dto.GetOrganizerDataDTORes{
		ID:                     organizerDAO.ID.Hex(),
		AverageScore:           organizerDAO.AverageScore,
		AverageTournamentScore: organizerDAO.AverageTournamentScore,
		TotalCategories:        organizerDAO.TotalCategories,
		TotalTournaments:       organizerDAO.TotalTournaments,
	}
}
