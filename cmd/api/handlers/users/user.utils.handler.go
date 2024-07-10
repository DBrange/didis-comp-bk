package handlers

import (
	"errors"
	"fmt"

	api_dto "github.com/DBrange/didis-comp-bk/cmd/api/handlers/users/dto"
	"github.com/DBrange/didis-comp-bk/cmd/api/handlers/users/mappers"
	location_dto "github.com/DBrange/didis-comp-bk/domains/location/models/dto"
	user_dto "github.com/DBrange/didis-comp-bk/domains/user/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/DBrange/didis-comp-bk/pkg/utils"

	"log"

	"github.com/gin-gonic/gin"
)

func saveBodyData(c *gin.Context) (*user_dto.CreateUserDTOReq, *location_dto.CreateLocationDTOReq, error) {
	var user api_dto.CreateUserDTOReq
	if err := c.ShouldBindJSON(&user); err != nil {
		return nil, nil, err
	}

	err := utils.Validate.Struct(user)

	if err != nil {
		log.Printf("este es el error real: %v", err)
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		if errors.Is(err, customerrors.ErrValidationFailed) {
			appErr := customerrors.AppError{
				Code: customerrors.ErrCodeValidationFailed,
				Msg:  fmt.Sprintf("error validation: %v", err),
			}
			return nil, nil, appErr
		}

		return nil, nil, fmt.Errorf("error validation: %w", err)
	}

	onlyUser := mappers.OnlyUser(user)

	onlyLocation := mappers.OnlyLocation(user)

	return onlyUser, onlyLocation, nil
}
