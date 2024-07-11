package handlers

import (
	"context"
	"fmt"
	"net/http"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	user, location, err := UpdateUserSaveBody(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if !user.AreAllFieldsNil() {
		fmt.Println("no tendria que estar aca")
		if err := h.user.UpdateUser(ctx, userID, user); err != nil {
			customerrors.ErrorResponse(err, c)
			return
		}
	}

	if location != nil {
		err := h.location.UpdateLocation(ctx, location.ID, location)
		if err != nil {
			customerrors.ErrorResponse(err, c)
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusOK, "message": "User updated successfully!"})
}
