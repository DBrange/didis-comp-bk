package routes

import (
	"didis-comp-bk/cmd/api/compose"
	"didis-comp-bk/cmd/api/compose/dashboard"
	handlers "didis-comp-bk/cmd/api/handlers/users"
	"time"

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

	userRoutes(router, handlers.NewHandlerUser(dashboard.User()))

}