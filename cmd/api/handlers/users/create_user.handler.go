package handlers

import (
	"context"
	"net/http"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	user, location, err := saveBodyData(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if location != nil {
		locationID, err := h.location.CreateLocation(ctx, location)
		if err != nil {
			customerrors.ErrorResponse(err, c)
			return
		}
		user.LocationID = &locationID
	}

	if err := h.user.CreateUser(ctx, user); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "User created successfully!"})
}
