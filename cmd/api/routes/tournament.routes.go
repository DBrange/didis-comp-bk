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

	tournamnetsRouter.PUT("/modify-bracket-match/:tournamentID/:userID", handler.ModifyBracketMatch)

	tournamnetsRouter.POST("/organize-bracket/:tournamentID", handler.OrganizeBracket)

	tournamnetsRouter.POST("/organize-tournament-group/:tournamentID/:roundID", handler.OrganizeTournamentGroups)

	tournamnetsRouter.PUT("/modify-tournament-group/:tournamentID/:roundID", handler.ModifyTournamentGroups)

	tournamnetsRouter.POST("/organize-pots/:tournamentID", handler.OrganizePots)

	tournamnetsRouter.PUT("/modify-pots/:tournamentID/:potID", handler.ModifyPots)

	tournamnetsRouter.PUT("/quantity-pots/:tournamentID", handler.UpdateQuantityPotsInTournament)

	tournamnetsRouter.PUT("/quantity-groups/:tournamentID", handler.UpdateQuantityGroupsInTournament)

	tournamnetsRouter.POST("/end-match", handler.EndMatch)

	tournamnetsRouter.POST("/end-tournament/:tournamentID", handler.EndTournament)

	tournamnetsRouter.GET("/list-competitors/:tournamentID", handler.ListCompetitorsInTournament)

	tournamnetsRouter.PUT("/round-total-prize/:roundID", handler.ModifyRoundTotalPrize)

	tournamnetsRouter.PUT("/round-points/:roundID", handler.ModifyRoundPoints)

	tournamnetsRouter.GET("/round-matches/:roundID", handler.GetRoundWithMatches)
}
