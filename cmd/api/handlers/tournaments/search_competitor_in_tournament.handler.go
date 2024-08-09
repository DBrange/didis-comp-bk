package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) SearchCompetitorInTournament(c *gin.Context) {
	// ctx, cancel := context.WithCancel(c.Request.Context())
	// defer cancel()

	// tournamentID := c.Param("tournamentID")
	// name := c.Query("name")

	// if err := h.tournament.SearchCompetitorInTournament(ctx, tournamentID, name); err != nil {
	// 	customerrors.ErrorResponse(err, c)
	// 	return
	// }

	// c.JSON(http.StatusCreated, gin.H{"status": http.StatusOK, "message": "Competitor finded!"})

}
