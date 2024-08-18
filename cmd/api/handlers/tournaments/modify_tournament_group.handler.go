package handlers

import (
	"context"
	"net/http"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ModifyTournamentGroups(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	tournamentID := c.Param("tournamentID")
	roundID := c.Param("roundID")

	competitorDTOs, err := organizeTournamentGroupsBodyData(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	sport, err := organizeTournamentGroupsValidateQueries(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if err := h.tournament.ModifyTournamentGroups(ctx, tournamentID, roundID, competitorDTOs, *sport); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Matches succsessfully updated!"})

}
