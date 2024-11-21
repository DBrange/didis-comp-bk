package handlers

import (
	"context"
	"net/http"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddTournamentInCategory(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	tournamentID := c.Param("tournamentID")
	categoryID := c.Param("categoryID")

	if err := h.category.AddTournamentInCategory(ctx, categoryID, tournamentID); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Tournament successfully added!"})
}
