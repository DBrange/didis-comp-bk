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

func (h *Handler) SearchCompetitorInCategory(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	name := c.Query("name")
	categoryID := c.Param("categoryID")
	sport, competitorType, err := searchCompetitorInCategoryValidateQueries(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	competitors, err := h.category.SearchCompetitorInCategory(ctx, categoryID, name, *sport, *competitorType)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": competitors, "status": http.StatusOK, "message": "Competitor finded!"})
}

func searchCompetitorInCategoryValidateQueries(c *gin.Context) (*model_utils.SPORT, *model_utils.COMPETITOR_TYPE, error) {
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
		return nil, nil, customerrors.HandleError(err, categoryErrorHandlers, errMsgTemplate)
	}

	sportParsed, competcompetitorTypeParsed, err := searchCompetitorInCategoryQueriesParser(sport, competitorType)
	if err != nil {
		return nil, nil, err
	}

	return &sportParsed, &competcompetitorTypeParsed, nil
}

func searchCompetitorInCategoryQueriesParser(sport, competitorType string) (model_utils.SPORT, model_utils.COMPETITOR_TYPE, error) {
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
