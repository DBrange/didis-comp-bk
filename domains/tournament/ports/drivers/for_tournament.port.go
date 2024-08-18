package ports

import "github.com/DBrange/didis-comp-bk/domains/tournament/ports/drivers/interfaces"

type ForTournament interface {
	interfaces.OrganizeTournament
	interfaces.AddCompetitorInTournament
	interfaces.AddGuestUserInTournament
	interfaces.ListCompetitorsInTournament
	interfaces.ModifyBracketMatch
	interfaces.ModifyRoundTotalPrize
	interfaces.GetRoundWithMatches
	interfaces.OrganizeBracket
	interfaces.EndMatch
	interfaces.EndTournament
	interfaces.ModifyRoundPoints
	interfaces.AddCompetitorInTournamentGroup
	interfaces.OrganizeTournamentGroups
	interfaces.ModifyTournamentGroups
	interfaces.OrganizePots
	interfaces.ModifyPots
	interfaces.UpdateQuantityPotsInTournament
	interfaces.UpdateQuantityGroupsInTournament
}
