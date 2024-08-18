package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	validate_util "github.com/DBrange/didis-comp-bk/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) OrganizeTournamentGroups(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	tournamentID := c.Param("tournamentID")
	roundID := c.Param("roundID")

	competitorDTOs, err := organizeTournamentGroupsBodyData(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	sport, err := organizeTournamentGroupsValidateQueries(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if err := h.tournament.OrganizeTournamentGroups(ctx, tournamentID, roundID, competitorDTOs, *sport); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Matches succsessfully updated!"})

}

func organizeTournamentGroupsBodyData(c *gin.Context) ([]*dto.AddCompetitorsToTournamentGroupsDTOReq, error) {
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

func organizeTournamentGroupsValidateQueries(c *gin.Context) (*models.SPORT, error) {
	sport := c.Query("sport")

	type ValidateQueries struct {
		Sport models.SPORT `json:"sport" validate:"sport"`
	}

	validateQueries := ValidateQueries{Sport: models.SPORT(sport)}

	err := validate_util.Validate.Struct(validateQueries)
	if err != nil {
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		tournamentErrorHandlers := customerrors.CreateErrorHandlers("tournament")
		errMsgTemplate := "error validation tournament"
		return nil, customerrors.HandleError(err, tournamentErrorHandlers, errMsgTemplate)
	}

	return &validateQueries.Sport, nil
}
