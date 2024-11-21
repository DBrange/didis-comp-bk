package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/DBrange/didis-comp-bk/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ModifyTournamentAvailability(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	availabilityID := c.Param("availabilityID")

	availabilityInfo, err := modifyTournamentAvailabilityBody(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if err = h.tournament.ModifyTournamentAvailability(ctx, availabilityID, availabilityInfo); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "The availability of the Tournament has been successfully modified!"})

}

func modifyTournamentAvailabilityBody(c *gin.Context) (*models.UpdateDailyAvailabilityDTOReq, error) {
	var availabilityInfo models.UpdateDailyAvailabilityDTOReq
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
