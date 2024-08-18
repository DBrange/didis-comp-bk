package handlers

import (
	"context"
	"net/http"

	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ModifyRoundTotalPrize(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	roundID := c.Param("roundID")

	totalPrize, err := utils.ParseToFloat(c, "prize_money")
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if err := h.tournament.ModifyRoundTotalPrize(ctx, roundID, totalPrize); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Round succsessfully updated!"})

}
