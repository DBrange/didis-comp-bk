package dashboard

import (
	category_ports "github.com/DBrange/didis-comp-bk/domains/category/ports/drivers"
	control_plane_ports "github.com/DBrange/didis-comp-bk/domains/control_plane/ports/drivers"
	location_ports "github.com/DBrange/didis-comp-bk/domains/location/ports/drivers"
	profile_ports "github.com/DBrange/didis-comp-bk/domains/profile/ports/drivers"
	tournament_ports "github.com/DBrange/didis-comp-bk/domains/tournament/ports/drivers"
)

type Dashboard interface {
	ControlPlane() control_plane_ports.ForControlPlane
	Profile() profile_ports.ForProfile
	Location() location_ports.ForLocation
	Tournament() tournament_ports.ForTournament
	Category() category_ports.ForCategory
}
