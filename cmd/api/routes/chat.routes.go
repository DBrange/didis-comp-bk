package routes

import (
	handlers "github.com/DBrange/didis-comp-bk/cmd/api/handlers/chats"
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

func chatRoutes(router *gin.Engine, socketServer *socketio.Server, handler *handlers.Handler) {
	chatRouter := router.Group("chats")

	chatRouter.GET("/enter/:chatID", handler.EnterChat)

	socketServer.OnEvent("/", string(models.SOCKET_EVENT_SEND_MESSAGE), handler.SendMessage)
}
