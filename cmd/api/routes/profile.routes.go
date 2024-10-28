package routes

import (
	handlers "github.com/DBrange/didis-comp-bk/cmd/api/handlers/profiles"
	ports "github.com/DBrange/didis-comp-bk/domains/control_plane/ports/drivers"

	"github.com/gin-gonic/gin"
)

func profileRoutes(router *gin.Engine, controlPlane ports.ForControlPlane, handler *handlers.Handler) {
	profilesRouter := router.Group("profiles")

	profilesRouter.POST("/register", handler.RegisterUser)

	// profilesRouter.POST("/login", controlPlane.AuthenticationMiddleware(), controlPlane.AuthorizationMiddleware(models.ROLE_FREE), handler.Login)
	profilesRouter.POST("/login", handler.Login)

	profilesRouter.POST("/refresh-token", handler.RefreshToken)

	profilesRouter.POST("/register-competitor", handler.RegisterCompetitor)

	profilesRouter.POST("/follow-profile/:fromUserID/:toUserID", handler.FollowProfile)

	profilesRouter.GET("/number-followers/:userID", handler.GetNumberFollowers)

	profilesRouter.GET("/user-followers/:userID", handler.GetUserFollowers)

	profilesRouter.GET("/personal-info/:userID", handler.GetPersonalInfo)

	profilesRouter.GET("/primary-info/:fromUserID/:toUserID", handler.GetUserPrimaryInfo)

	profilesRouter.GET("/daily-availability-info/:availabilityID", handler.GetProfileDailyAvailability)

	profilesRouter.GET("/info-category/:categoryID/:competitorID", handler.GetProfileInfoInCategory)

	profilesRouter.GET("/availability-category/:competitorID", handler.GetProfileAvailabilityInCategory)

	profilesRouter.GET("/tournaments-category/:categoryID/:competitorID", handler.GetProfileTournamentsInCategory)

	profilesRouter.GET("/categories/:userID", handler.GetProfileCategories)

	profilesRouter.GET("/tournaments/:userID", handler.GetProfileTournaments)

	profilesRouter.GET("/organizer-data/:userID", handler.GetOrganizerData)

	profilesRouter.PUT("/availability/:availabilityID", handler.ModifyProfileAvailability)

	profilesRouter.PUT("/personal-info/:userID", handler.ModifyPersonalInfo)

	profilesRouter.PUT("/new-password/:userID", handler.ModifyPassword)

	profilesRouter.DELETE("/close-profile/:userID", handler.CloseProfile)
}
