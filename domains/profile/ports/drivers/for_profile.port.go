//go:generate mockgen -destination=tests/mocks/for_profile.mock.go -package=mocks github.com/DBrange/didis-comp-bk/domains/profile/ports/drivers ForProfile

package ports

import "github.com/DBrange/didis-comp-bk/domains/profile/ports/drivers/interfaces"

type ForProfile interface {
	interfaces.RegisterUser
	interfaces.ModifyProfileAvailability
	interfaces.ModifyPersonalInfo
	interfaces.GetPersonalInfoByID
	interfaces.GetProfileDailyAvailabilityByID
	interfaces.CloseProfile
	interfaces.ModifyPassword
	interfaces.RegisterCompetitor
	interfaces.Login
	interfaces.RefreshToken
	interfaces.FollowProfile
	interfaces.GetProfileInfoInCategory
	interfaces.GetProfileAvailabilityCompetitor
	interfaces.GetProfileTournamentsInCategory
	interfaces.GetProfileCategories
	interfaces.GetNumberFollowers
	interfaces.GetUserFollowers
	interfaces.GetUserPrimaryInfo
	interfaces.GetOrganizerData
	interfaces.GetProfileUserTournaments
	interfaces.GetProfileCompetitorTournaments
}
