package handlers

import (
	"didis-comp-bk/internal/user/models/dto"
	customerrors "didis-comp-bk/pkg/custom_errors"
	"didis-comp-bk/pkg/utils"
	"errors"
	"fmt"

	"log"

	"github.com/gin-gonic/gin"
)

func saveBodyData(c *gin.Context) (*dto.CreateUserDTO, error) {
	var user dto.CreateUserDTO
	if err := c.ShouldBindJSON(&user); err != nil {
		return nil, err
	}

	user.SetTimeStamp()

	err := utils.Validate.Struct(user)

	if err != nil {
		log.Printf("este es el error real: %v", err)
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		if errors.Is(err, customerrors.ErrValidationFailed) {
			appErr := customerrors.AppError{
				Code: customerrors.ErrCodeValidationFailed,
				Msg:  fmt.Sprintf("error validation: %v", err),
			}
			return nil, appErr
		}
		
		return nil, fmt.Errorf("error validation: %w", err)
	}

	return &user, nil
}
