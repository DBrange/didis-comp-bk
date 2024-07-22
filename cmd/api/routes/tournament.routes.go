package routes

import (
	handlers "github.com/DBrange/didis-comp-bk/cmd/api/handlers/tournaments"
	"github.com/gin-gonic/gin"
)

func tournamentRoutes(router *gin.Engine, handler *handlers.Handler) {
	tournamnetsRouter := router.Group("tournaments")

	tournamnetsRouter.POST("/organize", handler.OrganizeTournament)
}
