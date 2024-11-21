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

func (h *Handler) ModifyTournamentGroups(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	tournamentID := c.Param("tournamentID")
	roundID := c.Param("roundID")

	competitorDTOs, err := modifyTournamentGroupsBodyData(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	sport, err := organizeTournamentGroupsValidateQueries(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if err := h.tournament.ModifyTournamentGroups(ctx, tournamentID, roundID, competitorDTOs, *sport); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Groups succsessfully updated!"})

}

func modifyTournamentGroupsBodyData(c *gin.Context) ([]*dto.AddCompetitorsToTournamentGroupsDTOReq, error) {
	var competitors struct {
		Groups []*dto.AddCompetitorsToTournamentGroupsDTOReq `json:"groups"`
	}

	if err := c.ShouldBindJSON(&competitors); err != nil {
		err = fmt.Errorf("%w: error binding json: %v", customerrors.ErrGetJSON, err.Error())
		tournamentErrorHandlers := customerrors.CreateErrorHandlers("tournament")
		errMsgTemplate := "error getting tournament"
		return nil, customerrors.HandleError(err, tournamentErrorHandlers, errMsgTemplate)
	}

	// Validar la estructura excepto el campo Location
	err := validate_util.Validate.Struct(competitors)
	if err != nil {
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		tournamentErrorHandlers := customerrors.CreateErrorHandlers("tournament")
		errMsgTemplate := "error validation tournament"
		return nil, customerrors.HandleError(err, tournamentErrorHandlers, errMsgTemplate)
	}

	return competitors.Groups, nil
}
