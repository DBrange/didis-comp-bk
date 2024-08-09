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

func (h *Handler) ListCategories(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	organizerID := c.Param("organizerID")
	sport, competitorType, err := listCategoriesValidateQueries(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	categories, err := h.category.ListCategories(ctx, organizerID, *sport, *competitorType)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": categories, "status": http.StatusOK, "message": "List categories finded!"})
}

func listCategoriesValidateQueries(c *gin.Context) (*model_utils.SPORT, *model_utils.COMPETITOR_TYPE, error) {
	sport := c.Query("sport")
	competitorType := c.Query("competitor_type")
fmt.Printf(" asadasd %s %s", sport, competitorType)
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

	sportParsed, competcompetitorTypeParsed, err := listCategoriesQueriesParser(sport, competitorType)
	if err != nil {
		return nil, nil, err
	}

	return &sportParsed, &competcompetitorTypeParsed, nil
}

func listCategoriesQueriesParser(sport, competitorType string) (model_utils.SPORT, model_utils.COMPETITOR_TYPE, error) {
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
