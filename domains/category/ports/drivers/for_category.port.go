package ports

import "github.com/DBrange/didis-comp-bk/domains/category/ports/drivers/interfaces"

type ForCategory interface {
	interfaces.OrganizeCategory
	interfaces.AddTournamentInCategory
	interfaces.AddCompetitorInCategory
	interfaces.SearchCompetitorInCategory
	interfaces.SearchCompetitorForCategory
	interfaces.ModifyCategoryInfo
	interfaces.GetCategoryInfo
	interfaces.GetParticipantsOfCategory
	interfaces.RemoveCompetitorFromCategory
	interfaces.ListCategories
	interfaces.GetTournamentsFromCategory
	interfaces.ModifyCompetitorPoints
	interfaces.AddGuestUserInCategory
}
