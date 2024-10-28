package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/competitor_user/dao"
)

// GetUserCategoriesDAOtoDTO mapea un array de GetUserCategoriesCategoryDAO a GetUserCategoriesCategoryDTO de manera eficiente
func GetUserCategoriesDAOtoDTO(daoCategories []*dao.GetUserCategoriesCategoryDAO) []*dto.GetUserCategoriesCategoryDTO {
	// Inicializar el slice de salida con la capacidad exacta del input
	dtoCategories := make([]*dto.GetUserCategoriesCategoryDTO, len(daoCategories))

	// Iterar sobre cada elemento de la lista DAO
	for i, daoCategory := range daoCategories {
		dtoCategory := &dto.GetUserCategoriesCategoryDTO{
			ID:   daoCategory.ID.Hex(), // Convertir ObjectID a string
			Name: daoCategory.Name,
		}

		// Mapear CompetitorData si existe
		if daoCategory.CompetitorData != nil {
			dtoCategory.CompetitorData = &dto.GetUserCategoriesCompetitorDataDTO{
				Points:          daoCategory.CompetitorData.Points,
				CurrentPosition: daoCategory.CompetitorData.CurrentPosition,
				Users:           mapUsersDAOtoDTO(daoCategory.CompetitorData.Users), // Mapear usuarios
			}
		}

		// Mapear Organizer si existe
		if daoCategory.Organizer != nil {
			dtoCategory.Organizer = &dto.GetUserCategoriesOrganizerDTO{
				ID:        daoCategory.Organizer.ID.Hex(), // Convertir ObjectID a string
				FirstName: daoCategory.Organizer.FirstName,
				LastName:  daoCategory.Organizer.LastName,
			}
		}

		// Asignar la categoría mapeada directamente a la posición correspondiente en el slice
		dtoCategories[i] = dtoCategory
	}

	return dtoCategories
}

// mapUsersDAOtoDTO mapea un slice de GetUserCategoriesUserDAO a GetUserCategoriesUserDTO
func mapUsersDAOtoDTO(daoUsers []*dao.GetUserCategoriesUserDAO) []*dto.GetUserCategoriesUserDTO {
	if len(daoUsers) == 0 {
		return []*dto.GetUserCategoriesUserDTO{} // Retornar un slice vacío en vez de nil
	}

	dtoUsers := make([]*dto.GetUserCategoriesUserDTO, len(daoUsers))
	for i, daoUser := range daoUsers {
		dtoUsers[i] = &dto.GetUserCategoriesUserDTO{
			ID:        daoUser.ID.Hex(), // Convertir ObjectID a string
			FirstName: daoUser.FirstName,
			LastName:  daoUser.LastName,
		}
	}

	return dtoUsers
}
