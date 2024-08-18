package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/category_registration/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateCategoryRegistrationCurrentPositionDTOtoDAO(categoryRegistrationDTO []*dto.GetCategoryRegistrationSortedByPointsDTORes, categoryID string, convert utils.ConvertToObjectIDFunc) ([]*dao.GetCategoryRegistrationSortedByPointsDAORes, *primitive.ObjectID, error) {
	categoryRegistrationDAO := make([]*dao.GetCategoryRegistrationSortedByPointsDAORes, len(categoryRegistrationDTO))

	for i, crDTO := range categoryRegistrationDTO {
		competitorOID, err := convert(crDTO.CompetitorID)
		if err != nil {
			return nil, nil, err
		}

		categoryRegistrationDAO[i] = &dao.GetCategoryRegistrationSortedByPointsDAORes{
			CompetitorID:    competitorOID,
			CurrentPosition: crDTO.CurrentPosition,
		}
	}

	categoryOID, err := convert(categoryID)
	if err != nil {
		return nil, nil, err
	}

	return categoryRegistrationDAO, categoryOID, nil
}
