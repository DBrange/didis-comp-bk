package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/DBrange/didis-comp-bk/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddGuestUserInCategory(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()
	categoryID := c.Param("categoryID")
	sport := c.Query("sport")
	competitorType := c.Query("competitor_type")

	guestUsersDTO, err := addGuestUserInCategoryBodyData(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if err := addGuestUserInCategoryValidateQueries(sport, competitorType); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	sportParsed, competitorTypeParsed, err := addGuestUserInCategoryQueriesParser(sport, competitorType)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if err := h.category.AddGuestUserInCategory(ctx, categoryID, guestUsersDTO, sportParsed, competitorTypeParsed); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Guest competitor successfully added!"})

}

func addGuestUserInCategoryBodyData(c *gin.Context) ([]*dto.CreateGuestUserDTOReq, error) {
	var guestUsersDTO struct {
		GuestUsers []*dto.CreateGuestUserDTOReq `json:"guest_users"`
	}

	if err := c.ShouldBindJSON(&guestUsersDTO); err != nil {
		err = fmt.Errorf("%w: error binding json: %v", customerrors.ErrGetJSON, err.Error())
		categoryErrorHandlers := customerrors.CreateErrorHandlers("category")
		errMsgTemplate := "error getting body"
		return nil, customerrors.HandleError(err, categoryErrorHandlers, errMsgTemplate)
	}

	// Validar la estructura excepto el campo Location
	err := utils.Validate.Struct(guestUsersDTO)
	if err != nil {
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		categoryErrorHandlers := customerrors.CreateErrorHandlers("category")
		errMsgTemplate := "error validation body"
		return nil, customerrors.HandleError(err, categoryErrorHandlers, errMsgTemplate)
	}

	return guestUsersDTO.GuestUsers, nil
}

func addGuestUserInCategoryValidateQueries(sport, competitorType string) error {
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

func addGuestUserInCategoryQueriesParser(sport, competitorType string) (models.SPORT, models.COMPETITOR_TYPE, error) {
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
