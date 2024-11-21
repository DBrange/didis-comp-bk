package routes

import (
	handlers "github.com/DBrange/didis-comp-bk/cmd/api/handlers/categories"
	"github.com/gin-gonic/gin"
)

func categoryRoutes(router *gin.Engine, handler *handlers.Handler) {
	categoriesRouter := router.Group("categories")

	categoriesRouter.POST("/organize/:organizerID", handler.OrganizeCategory)

	categoriesRouter.POST("/register-guest-competitor/:categoryID", handler.AddGuestUserInCategory)

	categoriesRouter.GET("/search-competitor-in-category/:categoryID", handler.SearchCompetitorInCategory)

	categoriesRouter.GET("/search-competitor-for-category/:userID", handler.SearchCompetitorForCategory)

	categoriesRouter.GET("/category-info/:categoryID", handler.GetCategoryInfo)

	categoriesRouter.GET("/competitors-category/:categoryID", handler.GetCompetitorsOfCategory)

	categoriesRouter.GET("/list-categories/:organizerID", handler.ListCategories)

	categoriesRouter.GET("/list-tournaments/:categoryID", handler.GetTournamentsFromCategory)

	categoriesRouter.GET("/list-tournaments-by-name/:categoryID", handler.GetTournamentsByNameFromCategory)

	categoriesRouter.PUT("/modify-info/:categoryID", handler.ModifyCategoryInfo)

	categoriesRouter.PUT("/modify-competitor-points/:categoryID/:competitorID", handler.ModifyCompetitorPoints)

	categoriesRouter.PUT("/add-tournament/:categoryID/:tournamentID", handler.AddTournamentInCategory)

	categoriesRouter.PUT("/add-competitor/:categoryID/:competitorID", handler.AddCompetitorInCategory)

	categoriesRouter.DELETE("/remove-competitor/:categoryID/:competitorID", handler.RemoveCompetitorFromCategory)
}
