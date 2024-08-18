package handlers

import (
	"context"
	"net/http"

	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateQuantityPotsInTournament(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	tournamentID := c.Param("tournamentID")

	position, err := utils.ParseToInt(c, "position")
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	add, err := utils.ParseToBool(c, "add")
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if err := h.tournament.UpdateQuantityPotsInTournament(ctx, tournamentID, position, add); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Tournament succsessfully updated!"})

}
