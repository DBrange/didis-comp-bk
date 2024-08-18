package handlers

import (
	"context"
	"fmt"
	"net/http"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	validate_util "github.com/DBrange/didis-comp-bk/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ModifyPots(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	tournamentID := c.Param("tournamentID")
	potID := c.Param("potID")

	competitorID, add, err := modifyPotsBodyData(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if err := h.tournament.ModifyPots(ctx, tournamentID, potID, competitorID, add); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Round succsessfully updated!"})
}

func modifyPotsBodyData(c *gin.Context) (string, bool, error) {
	var result struct {
		CompetitorID string `json:"competitor_id"`
		Add          bool   `json:"add"`
	}

	if err := c.ShouldBindJSON(&result); err != nil {
		err = fmt.Errorf("%w: error binding json: %v", customerrors.ErrGetJSON, err.Error())
		tournamentErrorHandlers := customerrors.CreateErrorHandlers("tournament")
		errMsgTemplate := "error getting tournament"
		return "", false, customerrors.HandleError(err, tournamentErrorHandlers, errMsgTemplate)
	}

	// Validar la estructura excepto el campo Location
	err := validate_util.Validate.Struct(result)
	if err != nil {
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		tournamentErrorHandlers := customerrors.CreateErrorHandlers("tournament")
		errMsgTemplate := "error validation tournament"
		return "", false, customerrors.HandleError(err, tournamentErrorHandlers, errMsgTemplate)
	}

	return result.CompetitorID, result.Add, nil
}
