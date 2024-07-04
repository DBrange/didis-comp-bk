package handlers

import (
	customerrors "didis-comp-bk/pkg/custom_errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(c *gin.Context) {
	user, err := saveBodyData(c)

	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if err := h.user.CreateUser(user); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "User created successfully!"})
}
