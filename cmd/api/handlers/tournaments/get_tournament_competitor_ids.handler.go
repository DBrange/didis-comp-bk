package handlers

import (
	"context"
	"net/http"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTournamentCompetitorIDs(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	tournamentID := c.Param("tournamentID")

	competitorIds, err := h.tournament.GetTournamentCompetitorIDs(ctx, tournamentID)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if competitorIds == nil{
		competitorIds = []string{}
	}

	c.JSON(http.StatusOK, gin.H{"data": competitorIds, "status": http.StatusOK, "message": "Competitor IDs successfully found!"})

}
