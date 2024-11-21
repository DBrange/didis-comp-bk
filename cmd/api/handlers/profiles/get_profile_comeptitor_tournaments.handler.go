package handlers

import (
	"context"
	"net/http"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProfileCompetitorTournaments(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
		defer cancel()

		lastID := c.Query("last_id")
		competitorID := c.Param("competitorID")
		categoryID := c.Query("category_id")
		sport,  limit, err := getProfileTournamentsValidateQueries(c)
		if err != nil {
			customerrors.ErrorResponse(err, c)
			return
		}

		tournaments, err := h.profile.GetProfileCompetitorTournaments(ctx, competitorID,categoryID, *sport, limit, lastID)
		if err != nil {
			customerrors.ErrorResponse(err, c)
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": tournaments, "status": http.StatusOK, "message": "Competitor finded!"})
	}

	