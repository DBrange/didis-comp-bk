package handlers

import (
	"context"
	"fmt"
	"net/http"

	models "github.com/DBrange/didis-comp-bk/cmd/api/models/options/tournament"
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	validate_util "github.com/DBrange/didis-comp-bk/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) OrganizeTournament(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	options, err := organizeTournamentValidateQueries(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	tournamentDTO, err := organizeTournamentBodyData(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if err := h.tournament.OrganizeTournament(ctx, tournamentDTO, options); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "New tournament created successfully!"})

}

func organizeTournamentBodyData(c *gin.Context) (*dto.OrganizeTournamentDTOReq, error) {
	var tournamentDTO dto.OrganizeTournamentDTOReq
	if err := c.ShouldBindJSON(&tournamentDTO); err != nil {
		err = fmt.Errorf("%w: error binding json: %v", customerrors.ErrGetJSON, err.Error())
		tournamentErrorHandlers := customerrors.CreateErrorHandlers("tournament")
		errMsgTemplate := "error getting tournament"
		return nil, customerrors.HandleError(err, tournamentErrorHandlers, errMsgTemplate)
	}

	// Validar la estructura excepto el campo Location
	err := validate_util.Validate.StructExcept(tournamentDTO, "location", "double_elimination")
	if err != nil {
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		tournamentErrorHandlers := customerrors.CreateErrorHandlers("tournament")
		errMsgTemplate := "error validation tournament"
		return nil, customerrors.HandleError(err, tournamentErrorHandlers, errMsgTemplate)
	}

	// Validar el campo Location si no es nil
	if tournamentDTO.Location != nil {
		err = validate_util.Validate.Struct(tournamentDTO.Location)
		if err != nil {
			err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
			tournamentErrorHandlers := customerrors.CreateErrorHandlers("tournament")
			errMsgTemplate := "error validation tournament"
			return nil, customerrors.HandleError(err, tournamentErrorHandlers, errMsgTemplate)
		}
	}

	// Validar el campo double_elimination si no es nil
	if tournamentDTO.DoubleElim != nil {
		err = validate_util.Validate.Struct(tournamentDTO.DoubleElim)
		if err != nil {
			err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
			tournamentErrorHandlers := customerrors.CreateErrorHandlers("tournament")
			errMsgTemplate := "error validation tournament"
			return nil, customerrors.HandleError(err, tournamentErrorHandlers, errMsgTemplate)
		}
	}

	return &tournamentDTO, nil
}

func organizeTournamentValidateQueries(c *gin.Context) (*models.OrganizeTournamentOptions, error) {
	pots, err := utils.ParseToInt(c, "quantity_pots")
	if err != nil {
		return nil, err
	}

	groups, err := utils.ParseToInt(c, "quantity_groups")
	if err != nil {
		return nil, err
	}

	validateQueries := &models.OrganizeTournamentOptions{ QuantityPots: pots, QuantityGroups: groups}

	err = validate_util.Validate.Struct(validateQueries)
	if err != nil {
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		tournamentErrorHandlers := customerrors.CreateErrorHandlers("tournament")
		errMsgTemplate := "error validation tournament"
		return nil, customerrors.HandleError(err, tournamentErrorHandlers, errMsgTemplate)
	}

	return validateQueries, nil
}
