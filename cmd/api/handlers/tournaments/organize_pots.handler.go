package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	validate_util "github.com/DBrange/didis-comp-bk/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) OrganizePots(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	tournamentID := c.Param("tournamentID")

	pots, err := organizePotsBodyData(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if err := h.tournament.OrganizePots(ctx, tournamentID, pots); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Round succsessfully updated!"})
}

func organizePotsBodyData(c *gin.Context) ([]*dto.SetPotCompetitorDTOReq, error) {
	var pots struct {
		Pots []*dto.SetPotCompetitorDTOReq `json:"pots"`
	}

	if err := c.ShouldBindJSON(&pots); err != nil {
		err = fmt.Errorf("%w: error binding json: %v", customerrors.ErrGetJSON, err.Error())
		tournamentErrorHandlers := customerrors.CreateErrorHandlers("tournament")
		errMsgTemplate := "error getting tournament"
		return nil, customerrors.HandleError(err, tournamentErrorHandlers, errMsgTemplate)
	}

	// Validar la estructura excepto el campo Location
	err := validate_util.Validate.Struct(pots)
	if err != nil {
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		tournamentErrorHandlers := customerrors.CreateErrorHandlers("tournament")
		errMsgTemplate := "error validation tournament"
		return nil, customerrors.HandleError(err, tournamentErrorHandlers, errMsgTemplate)
	}

	return pots.Pots, nil
}
