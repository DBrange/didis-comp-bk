package handlers

import (
	"context"
	"net/http"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProfileAvailabilityInCategory(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	competitorID := c.Param("competitorID")

	day := c.Query("day")

	availabilityInfo, availabilityID, err := h.profile.GetProfileAvailabilityInCategory(ctx, competitorID, day)
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

	c.JSON(http.StatusOK, gin.H{"data": data, "status": http.StatusOK, "message": "The search for competitor profile availability has been a success"})

}
