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

func (h *Handler) ModifyPersonalInfo(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	userID := c.Param("userID")

	personalInfo, err := modifyPersonalInfoBody(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if err := h.profile.ModifyPersonalInfo(ctx, userID, personalInfo); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "The personal info of the profile has been successfully modified!"})

}

func modifyPersonalInfoBody(c *gin.Context) (*profile_dto.ModifyPersonalInfoDTOReq, error) {
	var profileInfoDTO profile_dto.ModifyPersonalInfoDTOReq
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
