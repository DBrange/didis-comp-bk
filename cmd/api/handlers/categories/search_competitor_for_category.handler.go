package handlers

import (
	"context"
	"fmt"
	"net/http"

	model_utils "github.com/DBrange/didis-comp-bk/cmd/api/models"
	// "github.com/DBrange/didis-comp-bk/cmd/api/utils"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	validate_utils "github.com/DBrange/didis-comp-bk/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SearchCompetitorForCategory(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	userID := c.Param("userID")
	competitorType := c.Query("competitor_type")
	sport := c.Query("sport")
	// team := c.Query("team")
	name := c.Query("name")

	sportParsed, competitorTypeParsed, err := searchCompetitorForCategoryValidateQueries(sport, competitorType)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	competitors, err := h.category.SearchCompetitorForCategory(ctx, userID, name, *sportParsed, *competitorTypeParsed)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": competitors, "status": http.StatusOK, "message": "Competitor finded!"})

}

func searchCompetitorForCategoryValidateQueries(sport, competitorType string) (*model_utils.SPORT, *model_utils.COMPETITOR_TYPE, error) {
	type validateSearchCompetitorForCategoryQueries struct {
		Sport          string `json:"sport" validate:"sport,required"`
		CompetitorType string `json:"competitor_type" validate:"competitorType,required"`
	}

	validateQueries := &validateSearchCompetitorForCategoryQueries{Sport: sport, CompetitorType: competitorType}

	err := validate_utils.Validate.Struct(validateQueries)
	if err != nil {
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		profileErrorHandlers := customerrors.CreateErrorHandlers("profile")
		errMsgTemplate := "error validation profile"
		return nil, nil, customerrors.HandleError(err, profileErrorHandlers, errMsgTemplate)
	}

	sportParsed, competitorTypeParsed, err := searchCompetitorForCategoryQueriesParser(sport, competitorType)
	if err != nil {
		return nil, nil, err
	}
	return &sportParsed, &competitorTypeParsed, err
}

func searchCompetitorForCategoryQueriesParser(sport, competitorType string) (model_utils.SPORT, model_utils.COMPETITOR_TYPE, error) {
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