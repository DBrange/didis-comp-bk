package handlers

import (
	"context"
	"net/http"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)


func (h *Handler) GetTournamentPrimaryInfo(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	tournamentID := c.Param("tournamentID")

	tournament, err := h.tournament.GetTournamentPrimaryInfo(ctx, tournamentID)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tournament, "status": http.StatusOK, "message": "Tournament founded!"})

}