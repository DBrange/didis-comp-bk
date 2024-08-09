package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/DBrange/didis-comp-bk/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddGuestUserInTournament(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()
	tournamentID := c.Param("tournamentID")
	sport := c.Query("sport")
	competitorType := c.Query("competitor_type")

	guestUsersDTO, err := addGuestUserInTournamentBodyData(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if err := addGuestUserInTournamentValidateQueries(sport, competitorType); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	sportParsed, competitorTypeParsed, err := addGuestUserInTournamentQueriesParser(sport, competitorType)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if err := h.tournament.AddGuestUserInTournament(ctx, tournamentID, guestUsersDTO, sportParsed, competitorTypeParsed); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Guest competitor successfully added!"})

}

func addGuestUserInTournamentBodyData(c *gin.Context) ([]*dto.CreateGuestUserDTOReq, error) {
	var guestUsersDTO struct {
		GuestUsers []*dto.CreateGuestUserDTOReq `json:"guest_users"`
	}

	if err := c.ShouldBindJSON(&guestUsersDTO); err != nil {
		err = fmt.Errorf("%w: error binding json: %v", customerrors.ErrGetJSON, err.Error())
		tournamentErrorHandlers := customerrors.CreateErrorHandlers("tournament")
		errMsgTemplate := "error getting body"
		return nil, customerrors.HandleError(err, tournamentErrorHandlers, errMsgTemplate)
	}

	// Validar la estructura excepto el campo Location
	err := utils.Validate.Struct(guestUsersDTO)
	if err != nil {
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		tournamentErrorHandlers := customerrors.CreateErrorHandlers("tournament")
		errMsgTemplate := "error validation body"
		return nil, customerrors.HandleError(err, tournamentErrorHandlers, errMsgTemplate)
	}

	return guestUsersDTO.GuestUsers, nil
}

func addGuestUserInTournamentValidateQueries(sport, competitorType string) error {
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

func addGuestUserInTournamentQueriesParser(sport, competitorType string) (models.SPORT, models.COMPETITOR_TYPE, error) {
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
