package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/DBrange/didis-comp-bk/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) OrganizeTournament(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	tournamentInfoDTO, err := organizeTournamentBodyData(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if err := h.tournament.OrganizeTournament(ctx, tournamentInfoDTO); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "New tournament created successfully!"})

}

func organizeTournamentBodyData(c *gin.Context) (*dto.OrganizeTournamentDTOReq, error) {
	var tournamentInfoDTO dto.OrganizeTournamentDTOReq
	if err := c.ShouldBindJSON(&tournamentInfoDTO); err != nil {
		err = fmt.Errorf("%w: error getting the json: %v", customerrors.ErrGetJSON, err.Error())
		if errors.Is(err, customerrors.ErrValidationFailed) {
			appErr := customerrors.AppError{
				Code: customerrors.ErrCodeGetJSON,
				Msg:  fmt.Sprintf("error binding json: %v", err),
			}
			return nil, appErr
		}
		return nil, fmt.Errorf("error validation: %w", err)
	}

	// Validar la estructura excepto el campo Location
	err := utils.Validate.StructExcept(tournamentInfoDTO, "Location")
	if err != nil {
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		if errors.Is(err, customerrors.ErrValidationFailed) {
			appErr := customerrors.AppError{
				Code: customerrors.ErrCodeValidationFailed,
				Msg:  fmt.Sprintf("error validation: %v", err),
			}
			return nil, appErr
		}
		return nil, fmt.Errorf("error validation: %w", err)
	}

	// Validar el campo Location si no es nil
	if tournamentInfoDTO.Location != nil {
		err = utils.Validate.Struct(tournamentInfoDTO.Location)
		if err != nil {
			err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
			if errors.Is(err, customerrors.ErrValidationFailed) {
				appErr := customerrors.AppError{
					Code: customerrors.ErrCodeValidationFailed,
					Msg:  fmt.Sprintf("error validation: %v", err),
				}
				return nil, appErr
			}
			return nil, fmt.Errorf("error validation: %w", err)
		}
	}

	return &tournamentInfoDTO, nil
}
