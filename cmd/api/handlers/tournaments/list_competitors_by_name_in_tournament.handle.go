package handlers

import (
	"context"
	"net/http"

	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ListCompetitorsByNameInTournament(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	tournamentID := c.Param("tournamentID")
	name := c.Query("name")

	limit, err := utils.ParseToInt(c, "limit")
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	competitors, err := h.tournament.ListCompetitorsByNameInTournament(ctx, tournamentID,  name, limit)

	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": competitors, "status": http.StatusOK, "message": "Competitors founded!"})

}
