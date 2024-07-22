package handlers

import (
	"context"
	"net/http"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPersonalInfoByID(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	userID := c.Param("userID")

	userInfo, err := h.profile.GetPersonalInfoByID(ctx, userID)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": userInfo, "status": http.StatusOK, "message": "The profile search has been a success"})

}
