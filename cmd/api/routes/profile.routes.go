package routes

import (
	handlers "github.com/DBrange/didis-comp-bk/cmd/api/handlers/profiles"
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	ports "github.com/DBrange/didis-comp-bk/domains/control_plane/ports/drivers"

	"github.com/gin-gonic/gin"
)

func profileRoutes(router *gin.Engine, controlPlane ports.ForControlPlane, handler *handlers.Handler) {
	profilesRouter := router.Group("profiles")

	profilesRouter.POST("/register", handler.RegisterUser)

	profilesRouter.POST("/login", controlPlane.AuthenticationMiddleware(), controlPlane.AuthorizationMiddleware(models.ROLE_FREE), handler.Login)

	profilesRouter.POST("/register-competitor/:userID", handler.RegisterCompetitor)

	profilesRouter.GET("/personal-info/:userID", handler.GetPersonalInfo)

	profilesRouter.GET("/daily-availability-info/:availabilityID", handler.GetProfileDailyAvailability)

	profilesRouter.PUT("/availability/:availabilityID", handler.ModifyProfileAvailability)

	profilesRouter.PUT("/personal-info/:userID", handler.ModifyPersonalInfo)

	profilesRouter.PUT("/new-password/:userID", handler.ModifyPassword)

	profilesRouter.DELETE("/close-profile/:userID", handler.CloseProfile)
}
