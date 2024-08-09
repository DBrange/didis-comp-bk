package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/category/dao"
)

func GetTournamentsFromCategoryDAOtoDTO(tournamentsDAO []dao.GetTournamentsFromCategoryDAORes) []dto.GetTournamentsFromCategoryDTORes {
	tournamentsDTO := make([]dto.GetTournamentsFromCategoryDTORes, len(tournamentsDAO))
	
	for i, tournament := range tournamentsDAO {
		tournamentsDTO[i] = dto.GetTournamentsFromCategoryDTORes{
			ID:          tournament.ID.Hex(),
			Name:        tournament.Name,
			Points:      tournament.Points,
			StartDate:   tournament.StartDate,
			FinishtDate: tournament.FinishtDate,
		}
	}

	return tournamentsDTO
}
