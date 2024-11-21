package handlers

import (
	"context"
	"fmt"
	"net/http"

	model_utils "github.com/DBrange/didis-comp-bk/cmd/api/models"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	validate_utils "github.com/DBrange/didis-comp-bk/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTournamentsByNameFromCategory(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	name := c.Query("name")
	categoryID := c.Param("categoryID")
	sport, competitorType,  err := getTournamentsByNameFromCategoryValidateQueries(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	competitors, err := h.category.GetTournamentsByNameFromCategory(ctx, categoryID, *sport, *competitorType, name)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": competitors, "status": http.StatusOK, "message": "Competitor finded!"})
}

func getTournamentsByNameFromCategoryValidateQueries(c *gin.Context) (*model_utils.SPORT, *model_utils.COMPETITOR_TYPE, error) {
	sport := c.Query("sport")
	competitorType := c.Query("competitor_type")

	type validateSearchCompetitorForCategoryQueries struct {
		Sport          string `json:"sport" validate:"sport,required"`
		CompetitorType string `json:"competitor_type" validate:"competitorType,required"`
	}

	validateQueries := &validateSearchCompetitorForCategoryQueries{Sport: sport, CompetitorType: competitorType}

	err := validate_utils.Validate.Struct(validateQueries)
	if err != nil {
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		categoryErrorHandlers := customerrors.CreateErrorHandlers("category")
		errMsgTemplate := "error validation category"
		return nil, nil,  customerrors.HandleError(err, categoryErrorHandlers, errMsgTemplate)
	}

	sportParsed, competcompetitorTypeParsed, err := getTournamentsFromCategoryQueriesParser(sport, competitorType)
	if err != nil {
		return nil, nil,  err
	}

	return &sportParsed, &competcompetitorTypeParsed, nil
}

