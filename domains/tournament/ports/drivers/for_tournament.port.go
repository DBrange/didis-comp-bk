package ports

import "github.com/DBrange/didis-comp-bk/domains/tournament/ports/drivers/interfaces"

type ForTournament interface {
	interfaces.OrganizeTournament
	interfaces.AddCompetitorInTournament
	interfaces.AddGuestUserInTournament
	interfaces.ListCompetitorsInTournament
}
