package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/DBrange/didis-comp-bk/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterCompetitor(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	sport := c.Query("sport")
	competitorType := c.Query("competitor_type")

	if err := registerCompetitorValidateQueries(sport, competitorType); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	sportParsed, competitorTypeParsed, err := registerCompetitorQueriesParser(sport, competitorType)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	usersIDs, err := registerCompetitorBodyData(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if err := h.profile.RegisterCompetitor(ctx, usersIDs, sportParsed, competitorTypeParsed); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "New competitor created successfully!"})
}

func registerCompetitorBodyData(c *gin.Context) ([]string, error) {
	var Users struct {
		Users []string `json:"users"`
	}

	if err := c.ShouldBindJSON(&Users); err != nil {
		err = fmt.Errorf("%w: error binding json: %v", customerrors.ErrGetJSON, err.Error())
		profileErrorHandlers := customerrors.CreateErrorHandlers("profile")
		errMsgTemplate := "error getting profile"
		return nil, customerrors.HandleError(err, profileErrorHandlers, errMsgTemplate)
	}

	// Validar la estructura excepto el campo Location
	err := utils.Validate.Struct(Users)
	if err != nil {
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		profileErrorHandlers := customerrors.CreateErrorHandlers("profile")
		errMsgTemplate := "error validation profile"
		return nil, customerrors.HandleError(err, profileErrorHandlers, errMsgTemplate)
	}

	return Users.Users, nil
}

func registerCompetitorValidateQueries(sport, competitorType string) error {
	type validateRegisterCompetitorQueries struct {
		Sport          string `json:"sport" validate:"sport,required"`
		CompetitorType string `json:"competitor_type" validate:"competitorType,required"`
	}

	validateQueries := &validateRegisterCompetitorQueries{Sport: sport, CompetitorType: competitorType}

	err := utils.Validate.Struct(validateQueries)
	if err != nil {
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		profileErrorHandlers := customerrors.CreateErrorHandlers("profile")
		errMsgTemplate := "error validation profile"
		return customerrors.HandleError(err, profileErrorHandlers, errMsgTemplate)
	}

	return err
}

func registerCompetitorQueriesParser(sport, competitorType string) (models.SPORT, models.COMPETITOR_TYPE, error) {
	sportParsed, err := models.ParseSport(sport)
	if err != nil {
		return "", "", err
	}

	competitorTypeParsed, err := models.ParseCompetitorType(competitorType)
	if err != nil {
		return "", "", err
	}

	return sportParsed, competitorTypeParsed, nil
}
