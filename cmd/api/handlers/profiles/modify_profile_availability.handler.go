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

func (h *Handler) ModifyProfileAvailability(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	availabilityID := c.Param("availabilityID")
	
	availabilityInfo, err := modifyProfileAvailabilityBody(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if err = h.profile.ModifyProfileAvailability(ctx, availabilityID, availabilityInfo); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "The availability of the profile has been successfully modified!"})

}

func modifyProfileAvailabilityBody(c *gin.Context) (*profile_dto.UpdateDailyAvailabilityDTOReq, error) {
	var availabilityInfo profile_dto.UpdateDailyAvailabilityDTOReq
	if err := c.ShouldBindJSON(&availabilityInfo); err != nil {
		err = fmt.Errorf("%w: error getting the json: %v", customerrors.ErrGetJSON, err.Error())
		if errors.Is(err, customerrors.ErrValidationFailed) {
			appErr := customerrors.AppError{
				Code: customerrors.ErrCodeGetJSON,
				Msg:  fmt.Sprintf("error binding: %v", err),
			}
			return nil, appErr
		}
		return nil, fmt.Errorf("error getting the json: %w", err)
	}

	err := utils.Validate.Struct(availabilityInfo)
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

	return &availabilityInfo, nil

}
