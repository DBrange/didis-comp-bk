package routes

import (
	"time"

	category_handlers "github.com/DBrange/didis-comp-bk/cmd/api/handlers/categories"
	chat_handlers "github.com/DBrange/didis-comp-bk/cmd/api/handlers/chats"
	profile_handlers "github.com/DBrange/didis-comp-bk/cmd/api/handlers/profiles"
	tournament_handlers "github.com/DBrange/didis-comp-bk/cmd/api/handlers/tournaments"
	"github.com/DBrange/didis-comp-bk/compose"
	"github.com/DBrange/didis-comp-bk/compose/dashboard"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.MaxAge = 12 * time.Hour
	// Configurar Socket.IO
	socketServer := socketio.NewServer(nil)
	go socketServer.Serve()
	defer socketServer.Close()

	router.Use(cors.New(config))
	dashboard, _ := compose.Compose()
	RoutesHandler(router, socketServer, dashboard)
	return router
}

func RoutesHandler(router *gin.Engine, socketServer *socketio.Server, dashboard dashboard.Dashboard) {
	WSHandler(socketServer)

	profileRoutes(router, dashboard.ControlPlane(), profile_handlers.NewHandlerProfile(dashboard.Profile()))
	tournamentRoutes(router, tournament_handlers.NewHandlerTournament(dashboard.Tournament()))
	categoryRoutes(router, category_handlers.NewHandlerCategory(dashboard.Category()))
	chatRoutes(router, socketServer, chat_handlers.NewHandlerChat(dashboard.Chat(), socketServer))

}

func WSHandler(socketServer *socketio.Server) {
	socketServer.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		return nil
	})

	socketServer.OnDisconnect("/", func(s socketio.Conn, reason string) {
		// Aquí puedes manejar la desconexión
	})

}
