package handlers

import (
	"context"
	"net/http"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddCompetitorInTournamentGroup(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	groupID := c.Param("groupID")
	tournamentID := c.Param("tournamentID")
	competitorID := c.Param("competitorID")

	if err := h.tournament.AddCompetitorInTournamentGroup(ctx, groupID, tournamentID, competitorID); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Competitor successfully added!"})

}
