package handlers

import (
	"errors"
	"fmt"

	"github.com/DBrange/didis-comp-bk/internal/user/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/DBrange/didis-comp-bk/pkg/utils"

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
