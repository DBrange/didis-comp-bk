package routes

import (
	handlers "github.com/DBrange/didis-comp-bk/cmd/api/handlers/leagues"
	"github.com/gin-gonic/gin"
)

func leagueRoutes(router *gin.Engine, handler *handlers.Handler) {
	leaguesRouter := router.Group("leagues")

	leaguesRouter.POST("/organize/:organizerID", handler.OrganizeLeague)

	leaguesRouter.POST("/add-tournament/:leagueID/:tournamentID", handler.AddTournamentInLeague)

	leaguesRouter.POST("/add-competitor/:leagueID/:competitorID", handler.AddCompetitorInLeague)
}
