package handlers

import (
	"net/http"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.user.GetUserByID(id)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user, "status": http.StatusOK, "message": "user found"})
}
