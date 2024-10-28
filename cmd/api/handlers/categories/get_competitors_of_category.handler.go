package handlers

import (
	"context"
	"fmt"
	"net/http"

	model_utils "github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	validate_utils "github.com/DBrange/didis-comp-bk/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCompetitorsOfCategory(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	lastID := c.Query("last_id")
	categoryID := c.Param("categoryID")
	sport, competitorType, limit, err := getCompetitorsOfCategoryValidateQueries(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	competitors, err := h.category.GetParticipantsOfCategory(ctx, categoryID, *sport, *competitorType, limit, lastID)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": competitors, "status": http.StatusOK, "message": "Competitor finded!"})
}

func getCompetitorsOfCategoryValidateQueries(c *gin.Context) (*model_utils.SPORT, *model_utils.COMPETITOR_TYPE, int, error) {
	limit, err := utils.ParseToInt(c, "limit")
	if err != nil {
		return nil, nil, 0, err
	}

	sport := c.Query("sport")
	competitorType := c.Query("competitor_type")

	type validateSearchCompetitorForCategoryQueries struct {
		Limit          int    `json:"limit"`
		Sport          string `json:"sport" validate:"sport,required"`
		CompetitorType string `json:"competitor_type" validate:"competitorType,required"`
	}

	validateQueries := &validateSearchCompetitorForCategoryQueries{Sport: sport, CompetitorType: competitorType, Limit: limit}

	err = validate_utils.Validate.Struct(validateQueries)
	if err != nil {
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		categoryErrorHandlers := customerrors.CreateErrorHandlers("category")
		errMsgTemplate := "error validation category"
		return nil, nil, 0, customerrors.HandleError(err, categoryErrorHandlers, errMsgTemplate)
	}

	sportParsed, competcompetitorTypeParsed, err := getCompetitorsOfCategoryQueriesParser(sport, competitorType)
	if err != nil {
		return nil, nil, 0, err
	}

	return &sportParsed, &competcompetitorTypeParsed, validateQueries.Limit, nil
}

func getCompetitorsOfCategoryQueriesParser(sport, competitorType string) (model_utils.SPORT, model_utils.COMPETITOR_TYPE, error) {
	sportParsed, err := model_utils.ParseSport(sport)
	if err != nil {
		return "", "", err
	}

	competitorTypeParsed, err := model_utils.ParseCompetitorType(competitorType)
	if err != nil {
		return "", "", err
	}

	return sportParsed, competitorTypeParsed, nil
}
