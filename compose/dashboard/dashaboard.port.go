package dashboard

import (
	user_ports_drivers "github.com/DBrange/didis-comp-bk/domains/user/ports/drivers"
	location_ports_drivers "github.com/DBrange/didis-comp-bk/domains/location/ports/drivers"
)

type Dashboard interface {
	User() user_ports_drivers.ForUser
	Location() location_ports_drivers.ForLocation
}
