package routes

import (
	handlers "github.com/DBrange/didis-comp-bk/cmd/api/handlers/users"

	"github.com/gin-gonic/gin"
)

func userRoutes(router *gin.Engine, handler *handlers.Handler) {
	userRouter := router.Group("users")

	userRouter.POST("/", handler.CreateUser)

	userRouter.GET("/:id", handler.GetUserByID)
}
