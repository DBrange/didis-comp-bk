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

	"github.com/gin-gonic/gin"
)

func saveBodyData(c *gin.Context) (*user_dto.RegisterUserDTOReq, error) {
	var userInfoDTO user_dto.RegisterUserDTOReq
	if err := c.ShouldBindJSON(&userInfoDTO); err != nil {
		return nil, err
	}

	// Validar la estructura excepto el campo Location
	err := utils.Validate.StructExcept(userInfoDTO, "Location")
	if err != nil {
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

	// Validar el campo Location si no es nil
	if userInfoDTO.Location != nil {
		err = utils.Validate.Struct(userInfoDTO.Location)
		if err != nil {
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
	}

	return &userInfoDTO, nil
}

func UpdateUserSaveBody(c *gin.Context) (*user_dto.UpdateUserDTOReq, *location_dto.UpdateLocationDTOReq, error) {
	var attributesToUpdate api_dto.UpdateUserDTOReq
	if err := c.ShouldBindJSON(&attributesToUpdate); err != nil {
		return nil, nil, err
	}

	err := utils.Validate.StructExcept(attributesToUpdate, "Location")

	if err != nil {
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

	var onlyUpdateLocation *location_dto.UpdateLocationDTOReq
	if attributesToUpdate.Location != nil {
		err = utils.Validate.Struct(attributesToUpdate.Location)
		if err != nil {
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
		onlyUpdateLocation = mappers.OnlyUpdateLocation(&attributesToUpdate)
	}

	onlyUpdateUser := mappers.OnlyUpdateUser(&attributesToUpdate)

	return onlyUpdateUser, onlyUpdateLocation, nil
}
