package chats

import (
	"context"
	"net/http"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) EnterChat(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	chatID := c.Param("chatID")

	chat, err := h.chat.EnterChat(ctx, chatID)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": chat, "status": http.StatusCreated, "message": "Chat successfully found!"})

}
