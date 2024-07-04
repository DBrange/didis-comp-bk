package routes

import (
	handlers "didis-comp-bk/cmd/api/handlers/users"

	"github.com/gin-gonic/gin"
)

func userRoutes(router *gin.Engine, handler *handlers.Handler) {
	userRouter := router.Group("users")

	userRouter.POST("/", handler.CreateUser)

	router.GET("/:id", handler.GetUserByID)
}
