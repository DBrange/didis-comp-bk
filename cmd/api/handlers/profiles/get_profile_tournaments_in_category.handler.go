package handlers

import (
	"context"
	"net/http"

	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProfileTournamentsInCategory(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	categoryID := c.Param("categoryID")
	competitorID := c.Param("competitorID")
	lastID := c.Query("lastID")

	limit, err := utils.ParseToInt(c, "limit")
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	tournaments, err := h.profile.GetProfileTournamentsInCategory(ctx, categoryID, competitorID, lastID, limit)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tournaments, "status": http.StatusOK, "message": "The search for competitor tournaments has been a success"})

}
