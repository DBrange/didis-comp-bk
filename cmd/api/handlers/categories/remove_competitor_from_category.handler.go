package handlers

import (
	"context"
	"net/http"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RemoveCompetitorFromCategory(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	categoryID := c.Param("categoryID")
	competitorID := c.Param("competitorID")

	if err := h.category.RemoveCompetitorFromCategory(ctx, categoryID, competitorID); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": "CategoryRegistration was successfully deleted!"})
}
