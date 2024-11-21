package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	validate_util "github.com/DBrange/didis-comp-bk/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) OrganizeTournamentGroups(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	tournamentID := c.Param("tournamentID")
	roundID := c.Param("roundID")

	sport, err := organizeTournamentGroupsValidateQueries(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	orderType, top, availableCourts, averageHours, err := OrganizeTournamentGroupsValidateQueries(c, "order_type", "top", "available_courts", "average_hours")
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if err := h.tournament.OrganizeTournamentGroups(ctx, tournamentID, roundID, *sport, orderType, top, availableCourts, averageHours); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Matches succsessfully updated!"})

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

func OrganizeTournamentGroupsValidateQueries(c *gin.Context, orderType, top, availableCourts, averageHours string) (int, int, int, int, error) {
	orderTypeParsed, err := utils.ParseToInt(c, "order_type")
	if err != nil {
		return 0, 0, 0, 0, err
	}

	topParsed, err := utils.ParseToInt(c, "top")
	if err != nil {
		return 0, 0, 0, 0, err
	}

	availableCourtsParsed, err := utils.ParseToInt(c, "available_courts")
	if err != nil {
		return 0, 0, 0, 0, err
	}

	averageHoursParsed, err := utils.ParseToInt(c, "average_hours")
	if err != nil {
		return 0, 0, 0, 0, err
	}

	return orderTypeParsed, topParsed, availableCourtsParsed, averageHoursParsed, nil
}
