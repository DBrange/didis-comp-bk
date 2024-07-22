package dashboard

import (
	location_ports "github.com/DBrange/didis-comp-bk/domains/location/ports/drivers"
	profile_ports "github.com/DBrange/didis-comp-bk/domains/profile/ports/drivers"
	tournament_ports "github.com/DBrange/didis-comp-bk/domains/tournament/ports/drivers"
)

type Dashboard interface {
	Profile() profile_ports.ForProfile
	Location() location_ports.ForLocation
	Tournament() tournament_ports.ForTournament
}
