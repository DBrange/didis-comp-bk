package handlers

import (
	"context"
	"net/http"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCategoryInfo(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	categoryID := c.Param("categoryID")

	categoryInfoDTO, err := h.category.GetCategoryInfo(ctx, categoryID)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": categoryInfoDTO,"status": http.StatusCreated, "message": "Category information was successfully found!"})
}
