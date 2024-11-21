package handlers

import (
	"context"
	"net/http"

	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateMatchDate(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	matchID := c.Param("matchID")

	dateParsed, err := utils.ParseToTime(c, "date")
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if err := h.tournament.UpdateMatchDate(ctx, matchID, dateParsed); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Match date was successfully updated!"})
}
