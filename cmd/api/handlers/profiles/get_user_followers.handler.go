package handlers

import (
	"context"
	"net/http"

	"github.com/DBrange/didis-comp-bk/cmd/api/utils"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUserFollowers(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	userID := c.Param("userID")
	name := c.Query("name")

	limitParsed, err := utils.ParseToInt(c, "limit")
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	lastCreatedAtParsed, err := utils.ParseToTime(c, "last_created_at")
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	followers, err := h.profile.GetUserFollowers(ctx, userID,name, limitParsed, lastCreatedAtParsed)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": followers, "status": http.StatusOK, "message": "The search for user followers has been a success!"})
}
