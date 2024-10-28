package handlers

import (
	"context"
	"net/http"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTournamentAvailability(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	tournamentID := c.Param("tournamentID")

	day := c.Query("day")

	availabilityInfo, availabilityID, err := h.tournament.GetTournamentAvailability(ctx, tournamentID, day)
	if err != nil {
		customerrors.ErrorResponse(err, c)
		return
	}

	type Data struct {
		ID                                     string `json:"id"`
		*models.GetDailyAvailabilityByIDDTORes `json:",inline"`
	}

	data := Data{
		ID:                             availabilityID,
		GetDailyAvailabilityByIDDTORes: availabilityInfo,
	}

	c.JSON(http.StatusOK, gin.H{"data": data, "status": http.StatusOK, "message": "The search for tournament availability has been a success"})

}
