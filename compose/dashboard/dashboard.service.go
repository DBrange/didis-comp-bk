package dashboard

import (
	category_ports "github.com/DBrange/didis-comp-bk/domains/category/ports/drivers"
	chat_ports "github.com/DBrange/didis-comp-bk/domains/chat/ports/drivers"
	control_plane_ports "github.com/DBrange/didis-comp-bk/domains/control_plane/ports/drivers"
	location_ports "github.com/DBrange/didis-comp-bk/domains/location/ports/drivers"
	profile_ports "github.com/DBrange/didis-comp-bk/domains/profile/ports/drivers"
	tournament_ports "github.com/DBrange/didis-comp-bk/domains/tournament/ports/drivers"
)

type DashboardService struct {
	forControlPlane control_plane_ports.ForControlPlane
	forProfile      profile_ports.ForProfile
	forLocation     location_ports.ForLocation
	forTournament   tournament_ports.ForTournament
	forCategory     category_ports.ForCategory
	forChat         chat_ports.ForChat
}

func NewDashboardService(
	controlPlaneAdapter control_plane_ports.ForControlPlane,
	profileAdapter profile_ports.ForProfile,
	locationAdapter location_ports.ForLocation,
	tournamentAdapter tournament_ports.ForTournament,
	categoryAdapter category_ports.ForCategory,
	chatAdapter chat_ports.ForChat,
) *DashboardService {
	return &DashboardService{
		forControlPlane: controlPlaneAdapter,
		forProfile:      profileAdapter,
		forLocation:     locationAdapter,
		forTournament:   tournamentAdapter,
		forCategory:     categoryAdapter,
		forChat:         chatAdapter,
	}
}

func (d *DashboardService) ControlPlane() control_plane_ports.ForControlPlane {
	return d.forControlPlane
}

func (d *DashboardService) Profile() profile_ports.ForProfile {
	return d.forProfile
}

func (d *DashboardService) Location() location_ports.ForLocation {
	return d.forLocation
}

func (d *DashboardService) Tournament() tournament_ports.ForTournament {
	return d.forTournament
}

func (d *DashboardService) Category() category_ports.ForCategory {
	return d.forCategory
}

func (d *DashboardService) Chat() chat_ports.ForChat {
	return d.forChat
}
