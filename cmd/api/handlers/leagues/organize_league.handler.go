package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DBrange/didis-comp-bk/domains/league/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/DBrange/didis-comp-bk/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) OrganizeLeague(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	organizerID := c.Param("organizerID")

	leagueInfoDTO, err := organizeLeagueBodyData(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if err := h.league.OrganizeLeague(ctx, organizerID, leagueInfoDTO); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "New league created successfully!"})
}

func organizeLeagueBodyData(c *gin.Context) (*dto.CreateLeagueDTOReq, error) {
	var leagueDTO dto.CreateLeagueDTOReq
	if err := c.ShouldBindJSON(&leagueDTO); err != nil {
		err = fmt.Errorf("%w: error binding json: %v", customerrors.ErrGetJSON, err.Error())
		leagueErrorHandlers := customerrors.CreateErrorHandlers("league")
		errMsgTemplate := "error getting league"
		return nil, customerrors.HandleError(err, leagueErrorHandlers, errMsgTemplate)
	}

	err := utils.Validate.Struct(leagueDTO)
	if err != nil {
		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
		leagueErrorHandlers := customerrors.CreateErrorHandlers("league")
		errMsgTemplate := "error validation league"
		return nil, customerrors.HandleError(err, leagueErrorHandlers, errMsgTemplate)
	}

	return &leagueDTO, nil
}
