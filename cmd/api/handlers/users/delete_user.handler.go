package handlers

import (
	"context"
	"net/http"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteUser(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	userID := c.Param("id")

	userRelationsToDelete, err := h.user.DeleteUser(ctx, userID)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	err = h.location.DeleteLocation(ctx, userRelationsToDelete.LocationID)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "the user has been successfully removed"})
}
