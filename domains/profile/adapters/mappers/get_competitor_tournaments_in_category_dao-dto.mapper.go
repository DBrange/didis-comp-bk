package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/category/dao"
)

func GetCompetitorTournamentsInCategoryDAOtoDTO(tournamentsDAO []*dao.GetTournamentsFromCategoryDAORes) []*dto.GetTournamentsFromCategoryDTORes {
	tournamentsDTO := make([]*dto.GetTournamentsFromCategoryDTORes, len(tournamentsDAO))

	for i, tournamentDAO := range tournamentsDAO {
		tournamentsDTO[i] = &dto.GetTournamentsFromCategoryDTORes{
			ID:          tournamentDAO.ID.Hex(),
			Name:        tournamentDAO.Name,
			Points:      tournamentDAO.Points,
			StartDate:   tournamentDAO.StartDate,
			FinishtDate: tournamentDAO.FinishtDate,
		}
	}

	return tournamentsDTO
}
