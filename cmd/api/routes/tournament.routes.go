package routes

import (
	handlers "github.com/DBrange/didis-comp-bk/cmd/api/handlers/tournaments"
	"github.com/gin-gonic/gin"
)

func tournamentRoutes(router *gin.Engine, handler *handlers.Handler) {
	tournamnetsRouter := router.Group("tournaments")

	tournamnetsRouter.POST("/organize", handler.OrganizeTournament)

	tournamnetsRouter.PUT("/add-competitor", handler.AddCompetitorInTournament)

	tournamnetsRouter.PUT("/add-competitor-group/:groupID/:tournamentID/:competitorID", handler.AddCompetitorInTournamentGroup)

	tournamnetsRouter.PUT("/register-guest-competitor/:tournamentID", handler.AddGuestUserInTournament)

	tournamnetsRouter.PUT("/register-double-competitor/:tournamentID", handler.RegisterDoubleCompetitorInTournament)

	tournamnetsRouter.PUT("/modify-bracket-match/:tournamentID/:userID", handler.ModifyBracketMatch)

	tournamnetsRouter.PUT("/organize-bracket/:tournamentID", handler.OrganizeBracket)

	tournamnetsRouter.PUT("/organize-tournament-groups/:tournamentID/:roundID", handler.OrganizeTournamentGroups)

	tournamnetsRouter.PUT("/modify-tournament-group/:tournamentID/:roundID", handler.ModifyTournamentGroups)

	tournamnetsRouter.POST("/organize-pots/:tournamentID", handler.OrganizePots)

	tournamnetsRouter.PUT("/modify-pots/:tournamentID/:potID", handler.ModifyPots)

	tournamnetsRouter.PUT("/quantity-pots/:tournamentID", handler.UpdateQuantityPotsInTournament)

	tournamnetsRouter.PUT("/quantity-groups/:tournamentID", handler.UpdateQuantityGroupsInTournament)

	tournamnetsRouter.PUT("/end-match", handler.EndMatch)

	tournamnetsRouter.POST("/end-tournament/:tournamentID", handler.EndTournament)

	tournamnetsRouter.GET("/list-competitors/:tournamentID", handler.ListCompetitorsInTournament)

	tournamnetsRouter.GET("/list-competitors-by-name/:tournamentID", handler.ListCompetitorsByNameInTournament)

	tournamnetsRouter.GET("/search-follower-for-tournament/:userID", handler.SearchCompetitorForTournament)

	tournamnetsRouter.GET("/user-tournaments/:userID", handler.GetUserTournaments)

	tournamnetsRouter.GET("/organizer-tournaments/:organizerID", handler.GetTournamentsInOrganizer)

	tournamnetsRouter.GET("/primary-info/:tournamentID", handler.GetTournamentPrimaryInfo)

	tournamnetsRouter.GET("/round-matches/:roundID", handler.GetRoundWithMatches)

	tournamnetsRouter.GET("/round-groups/:roundID", handler.GetRoundGroups)

	tournamnetsRouter.GET("/filters/:tournamentID", handler.GetTournamentFilters)

	tournamnetsRouter.GET("/competitor-ids/:tournamentID", handler.GetTournamentCompetitorIDs)

	tournamnetsRouter.GET("/availability-tournament/:tournamentID", handler.GetTournamentAvailability)

	tournamnetsRouter.PUT("/round-total-prize/:roundID", handler.ModifyRoundTotalPrize)

	tournamnetsRouter.PUT("/round-points/:roundID", handler.ModifyRoundPoints)

	tournamnetsRouter.PUT("/remove-competitor/:tournamentID/:competitorID", handler.RemoveCompetitorFromTournament)
}
