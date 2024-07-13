package handlers

import (
	"context"
	"net/http"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterUser(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	userInfoDTO, err := saveBodyData(c)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	if err := h.user.RegisterUser(ctx, userInfoDTO); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "User created successfully!"})
}
