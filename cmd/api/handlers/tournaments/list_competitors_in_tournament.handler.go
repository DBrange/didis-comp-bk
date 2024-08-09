package handlers

import (
	"context"
	"net/http"

	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ListCompetitorsInTournament(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	tournamentID := c.Param("tournamentID")
	categoryID := c.Query("category_id")

	lastID := c.Query("lastID")

	limit, err := utils.ParseToInt(c, "limit")
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	competitors, err := h.tournament.ListCompetitorsInTournament(ctx, tournamentID, categoryID, lastID, limit)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": competitors, "status": http.StatusCreated, "message": "New tournament created successfully!"})

}
