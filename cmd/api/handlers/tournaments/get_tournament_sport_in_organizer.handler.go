package handlers

import (
	"context"
	"net/http"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTournamentSportsInOrganizer(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	organizerID := c.Param("organizerID")

	tournament, err := h.tournament.GetTournamentSportsInOrganizer(ctx, organizerID)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tournament, "status": http.StatusOK, "message": "Tournament founded!"})

}
