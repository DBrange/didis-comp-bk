package routes

import (
	handlers "github.com/DBrange/didis-comp-bk/cmd/api/handlers/profiles"

	"github.com/gin-gonic/gin"
)

func profileRoutes(router *gin.Engine, handler *handlers.Handler) {
	profilesRouter := router.Group("profiles")

	profilesRouter.POST("/register", handler.RegisterUser)

	profilesRouter.PUT("/availability/:availabilityID", handler.ModifyProfileAvailability)

	profilesRouter.PUT("/personal-info/:userID", handler.ModifyPersonalInfo)

	profilesRouter.GET("/personal-info/:userID", handler.GetPersonalInfoByID)

	profilesRouter.GET("/availability-info/:availabilityID", handler.GetProfileAvailabilityInfo)

	profilesRouter.DELETE("/close-profile/:userID", handler.CloseProfile)

	profilesRouter.PUT("/new-password/:userID", handler.ModifyPassword)

	profilesRouter.POST("/add-competitor/:userID", handler.RegisterCompetitor)

}
