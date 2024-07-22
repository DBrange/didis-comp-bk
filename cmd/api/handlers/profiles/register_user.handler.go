package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	profile_dto "github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/DBrange/didis-comp-bk/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterUser(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	profileInfoDTO, err := registerUserBodyData(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if err := h.profile.RegisterUser(ctx, profileInfoDTO); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "New profile created successfully!"})
}

func registerUserBodyData(c *gin.Context) (*profile_dto.RegisterUserDTOReq, error) {
	var profileInfoDTO profile_dto.RegisterUserDTOReq
	if err := c.ShouldBindJSON(&profileInfoDTO); err != nil {
		err = fmt.Errorf("%w: error getting the json: %v", customerrors.ErrGetJSON, err.Error())
		if errors.Is(err, customerrors.ErrValidationFailed) {
			appErr := customerrors.AppError{
				Code: customerrors.ErrCodeGetJSON,
				Msg:  fmt.Sprintf("error binding json: %v", err),
			}
			return nil, appErr
		}
		return nil, fmt.Errorf("error validation: %w", err)
	}

	// Validar la estructura excepto el campo Location
	err := utils.Validate.StructExcept(profileInfoDTO, "Location")
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
	if profileInfoDTO.Location != nil {
		err = utils.Validate.Struct(profileInfoDTO.Location)
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

	return &profileInfoDTO, nil
}
