package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/category/dao"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateCategoryDTOtoDAO(categoryDTO *dto.UpdateCategoryDTOReq, categoryID string, convert utils.ConvertToObjectIDFunc) (*dao.UpdateCategoryDAOReq, *primitive.ObjectID, error) {
	categoryDAO := &dao.UpdateCategoryDAOReq{
		Name:          categoryDTO.Name,
		RangeMovement: categoryDTO.RangeMovement,
	}

	categoryOID, err := convert(categoryID)
	if err != nil {
		return nil, nil, err
	}

	return categoryDAO, categoryOID, nil
}
