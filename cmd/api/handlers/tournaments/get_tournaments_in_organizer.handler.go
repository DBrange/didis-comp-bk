package handlers

import (
	// "context"
	"context"
	"fmt"
	"net/http"

	model_utils "github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	validate_utils "github.com/DBrange/didis-comp-bk/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTournamentsInOrganizer(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	organizerID := c.Param("organizerID")
	lastID := c.Query("last_id")
	sport, limit, err := getTournamentsInOrganizerValidateQueries(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	categories, err := h.tournament.GetTournamentsInOrganizer(ctx, organizerID,sport,limit,lastID)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories, "status": http.StatusOK, "message": "Tournaments finded!"})
}

func getTournamentsInOrganizerValidateQueries(c *gin.Context) (model_utils.SPORT, int, error) {
	limit, err := utils.ParseToInt(c, "limit")
	if err != nil {
		return "", 0, err
	}

	sport := c.Query("sport")

	type validateSearchCompetitorForCategoryQueries struct {
		Limit int    `json:"limit"`
		Sport string `json:"sport" validate:"sport,required"`
	}

	validateQueries := &validateSearchCompetitorForCategoryQueries{Sport: sport, Limit: limit}

	err = validate_utils.Validate.Struct(validateQueries)
	if err != nil {
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		categoryErrorHandlers := customerrors.CreateErrorHandlers("category")
		errMsgTemplate := "error validation category"
		return "", 0, customerrors.HandleError(err, categoryErrorHandlers, errMsgTemplate)
	}

	sportParsed, err := getTournamentsInOrganizerQueriesParser(sport)
	if err != nil {
		return "", 0, err
	}

	return sportParsed, validateQueries.Limit, nil
}

func getTournamentsInOrganizerQueriesParser(sport string) (model_utils.SPORT, error) {
	sportParsed, err := model_utils.ParseSport(sport)
	if err != nil {
		return "", err
	}

	return sportParsed, nil
}