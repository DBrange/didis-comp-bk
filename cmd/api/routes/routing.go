package routes

import (
	"time"

	category_handlers "github.com/DBrange/didis-comp-bk/cmd/api/handlers/categories"
	profile_handlers "github.com/DBrange/didis-comp-bk/cmd/api/handlers/profiles"
	tournament_handlers "github.com/DBrange/didis-comp-bk/cmd/api/handlers/tournaments"
	"github.com/DBrange/didis-comp-bk/compose"
	"github.com/DBrange/didis-comp-bk/compose/dashboard"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.MaxAge = 12 * time.Hour

	router.Use(cors.New(config))
	dashboard, _ := compose.Compose()
	RoutesHandler(router, dashboard)
	return router
}

func RoutesHandler(router *gin.Engine, dashboard dashboard.Dashboard) {
	profileRoutes(router, dashboard.ControlPlane(), profile_handlers.NewHandlerProfile(dashboard.Profile()))
	tournamentRoutes(router, tournament_handlers.NewHandlerTournament(dashboard.Tournament()))
	categoryRoutes(router, category_handlers.NewHandlerCategory(dashboard.Category()))
}
