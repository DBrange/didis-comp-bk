package handlers

import (
	"context"
	"net/http"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProfileAvailabilityInCategory(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	competitorID := c.Param("competitorID")

	day := c.Query("day")

	availabilityInfo, err := h.profile.GetProfileAvailabilityInCategory(ctx, competitorID, day)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": availabilityInfo, "status": http.StatusOK, "message": "The search for competitor profile availability has been a success"})

}
