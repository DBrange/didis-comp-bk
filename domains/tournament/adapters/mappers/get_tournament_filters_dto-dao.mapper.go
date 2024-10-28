package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/tournament/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

func GetTournamentFiltersDTOtoDAO(filtersDAO *dao.GetTournamentFiltersDAORes) *dto.GetTournamentFiltersDTORes{
	return &dto.GetTournamentFiltersDTORes{
		TotalCompetitors: filtersDAO.TotalCompetitors,
		MaxCapacity: filtersDAO.MaxCapacity,
		Sport: filtersDAO.Sport,
		Surface: filtersDAO.Surface,
		CompetitorType: filtersDAO.CompetitorType,
	}
}