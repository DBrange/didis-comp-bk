package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/organizer/dao"
)

func GetCategoriesFromOrganizerDAOtoDTO(daoSlice []dao.GetCategoriesFromOrganizerDAORes) []dto.GetCategoriesFromOrganizerDTORes {
	dtoSlice := make([]dto.GetCategoriesFromOrganizerDTORes, len(daoSlice))
	for i, daoItem := range daoSlice {
		dtoSlice[i] = MapGetCategoriesFromOrganizerDAOToDTO(daoItem)
	}
	return dtoSlice
}

func MapGetCategoriesFromOrganizerDAOToDTO(dao dao.GetCategoriesFromOrganizerDAORes) dto.GetCategoriesFromOrganizerDTORes {
	return dto.GetCategoriesFromOrganizerDTORes{
		CategoryID:        dao.CategoryID.Hex(),
		Competitors:       MapGetCategoriesFromOrganizerCompetitorDAOToDTOSlice(dao.Competitors),
		TotalParticipants: dao.TotalParticipants,
	}
}

func MapGetCategoriesFromOrganizerCompetitorDAOToDTOSlice(daoSlice []dao.GetCategoriesFromOrganizerCompetitorDAORes) []dto.GetCategoriesFromOrganizerCompetitorDTORes {
	dtoSlice := make([]dto.GetCategoriesFromOrganizerCompetitorDTORes, len(daoSlice))
	for i, daoItem := range daoSlice {
		dtoSlice[i] = MapGetCategoriesFromOrganizerCompetitorDAOToDTO(daoItem)
	}
	return dtoSlice
}

func MapGetCategoriesFromOrganizerCompetitorDAOToDTO(dao dao.GetCategoriesFromOrganizerCompetitorDAORes) dto.GetCategoriesFromOrganizerCompetitorDTORes {
	return dto.GetCategoriesFromOrganizerCompetitorDTORes{
		CompetitorID:        dao.CompetitorID.Hex(),
		CurrentPosition:     dao.CurrentPosition,
		RegisteredPositions: MapGetCategoriesFromOrganizerRegisteredPositionDAOToDTO(dao.RegisteredPositions),
		Points:              dao.Points,
		Users:               MapGetCategoriesFromOrganizerUserDAOToDTOSlice(dao.Users),
	}
}

func MapGetCategoriesFromOrganizerUserDAOToDTOSlice(daoSlice []dao.GetCategoriesFromOrganizerUserDAORes) []dto.GetCategoriesFromOrganizerUserDTORes {
	dtoSlice := make([]dto.GetCategoriesFromOrganizerUserDTORes, len(daoSlice))
	for i, daoItem := range daoSlice {
		dtoSlice[i] = MapGetCategoriesFromOrganizerUserDAOToDTO(daoItem)
	}
	return dtoSlice
}

func MapGetCategoriesFromOrganizerUserDAOToDTO(dao dao.GetCategoriesFromOrganizerUserDAORes) dto.GetCategoriesFromOrganizerUserDTORes {
	return dto.GetCategoriesFromOrganizerUserDTORes{
		UserID:    dao.UserID.Hex(),
		FirstName: dao.FirstName,
		LastName:  dao.LastName,
		Image:     dao.Image,
	}
}

func MapGetCategoriesFromOrganizerRegisteredPositionDAOToDTO(dao []dao.RegisteredPositionsDAORes) []dto.RegistedPositionDTORes {
	registeredPositions := make([]dto.RegistedPositionDTORes, len(dao))

	for i, rp := range registeredPositions {
		registeredPositions[i] = dto.RegistedPositionDTORes{
			Date:     rp.Date,
			Position: rp.Position,
		}
	}

	return registeredPositions
}
