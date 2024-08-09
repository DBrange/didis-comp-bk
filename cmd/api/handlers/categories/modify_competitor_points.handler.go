package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ModifyCompetitorPoints(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	categoryID := c.Param("categoryID")
	competitorID := c.Param("competitorID")

	pointsNum, err := utils.ParseToInt(c, "points")
	if err != nil {
		fmt.Print("aaaa")
		customerrors.ErrorResponse(err, c)
		return
	}

	if err := h.category.ModifyCompetitorPoints(ctx, categoryID, competitorID, pointsNum); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusOK, "message": "Competitor has successfully updated!"})
}
