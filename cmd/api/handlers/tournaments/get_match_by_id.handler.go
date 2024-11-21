package handlers

import (
	"context"
	"net/http"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetMatchByID(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	matchID := c.Param("matchID")

	match, err := h.tournament.GetMatchByID(ctx, matchID)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": match, "status": http.StatusOK, "message": "Match founded!"})

}
