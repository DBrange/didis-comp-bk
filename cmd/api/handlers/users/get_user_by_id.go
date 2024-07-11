package handlers

import (
	"context"
	"net/http"

	"github.com/DBrange/didis-comp-bk/cmd/api/handlers/users/mappers"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUserByID(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()
	
	id := c.Param("id")
	user, err := h.user.GetUserByID(ctx,id)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	location, err := h.location.GetLocationByID(ctx, *user.LocationID)

	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}
	completeUser := mappers.UserAndLocation(user, location)

	c.JSON(http.StatusOK, gin.H{"data": completeUser, "status": http.StatusOK, "message": "user found"})
}
