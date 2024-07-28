package dashboard

import (
	control_plane_ports "github.com/DBrange/didis-comp-bk/domains/control_plane/ports/drivers"
	league_ports "github.com/DBrange/didis-comp-bk/domains/league/ports/drivers"
	location_ports "github.com/DBrange/didis-comp-bk/domains/location/ports/drivers"
	profile_ports "github.com/DBrange/didis-comp-bk/domains/profile/ports/drivers"
	tournament_ports "github.com/DBrange/didis-comp-bk/domains/tournament/ports/drivers"
)

type DashboardService struct {
	forControlPlane control_plane_ports.ForControlPlane
	forProfile      profile_ports.ForProfile
	forLocation     location_ports.ForLocation
	forTournament   tournament_ports.ForTournament
	forLeague       league_ports.ForLeague
}

func NewDashboardService(
	controlPlaneAdapter control_plane_ports.ForControlPlane,
	profileAdapter profile_ports.ForProfile,
	locationAdapter location_ports.ForLocation,
	tournamentAdapter tournament_ports.ForTournament,
	leagueAdapter league_ports.ForLeague,
) *DashboardService {
	return &DashboardService{
		forControlPlane: controlPlaneAdapter,
		forProfile:      profileAdapter,
		forLocation:     locationAdapter,
		forTournament:   tournamentAdapter,
		forLeague:       leagueAdapter,
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

func (d *DashboardService) League() league_ports.ForLeague {
	return d.forLeague
}
