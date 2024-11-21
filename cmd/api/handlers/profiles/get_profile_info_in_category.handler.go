package handlers

import (
	"context"
	"net/http"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProfileInfoInCategory(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	categoryID := c.Param("categoryID")
	competitorID := c.Param("competitorID")

	profile, err := h.profile.GetProfileInfoInCategory(ctx, categoryID, competitorID)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": profile, "status": http.StatusOK, "message": "Profile has been successfully found!"})

}
