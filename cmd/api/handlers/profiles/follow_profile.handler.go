package handlers

import (
	"context"
	"net/http"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) FollowProfile(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	fromUserID := c.Param("fromUserID")
	toUserID := c.Param("toUserID")

	if err := h.profile.FollowProfile(ctx, fromUserID, toUserID); err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "You hace successfully followed the user!"})
}
