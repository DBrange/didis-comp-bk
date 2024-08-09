package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/category_registration/dao"
)

func CreateCategoryRegistrationDTOtoDAO(categoryRegistrationDTO *dto.CreateCategoryRegistrationDTOReq, convert utils.ConvertToObjectIDFunc) (*dao.CreateCategoryRegistrationDAOReq, error) {
	categoryOID, err := convert(categoryRegistrationDTO.CategoryID)
	if err != nil {
		return nil, err
	}

	competitorOID, err := convert(categoryRegistrationDTO.CompetitorID)
	if err != nil {
		return nil, err
	}

	categoryRegistrationDAO := &dao.CreateCategoryRegistrationDAOReq{
		CompetitorID: *competitorOID,
		CategoryID:   *categoryOID,
	}

	return categoryRegistrationDAO, nil
}
