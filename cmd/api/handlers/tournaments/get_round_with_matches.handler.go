package handlers

import (
	"context"
	"net/http"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetRoundWithMatches(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	roundID := c.Param("roundID")
	categoryID := c.Query("category_id")

	round, err := h.tournament.GetRoundWithMatches(ctx, roundID,categoryID)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": round, "status": http.StatusOK, "message": "Round successfully found!"})

}
