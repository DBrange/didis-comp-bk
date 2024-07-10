package dashboard

import (
	location_ports "github.com/DBrange/didis-comp-bk/domains/location/ports/drivers"
	user_adap_drivers "github.com/DBrange/didis-comp-bk/domains/user/adapters/drivers"
	user_ports_drivers "github.com/DBrange/didis-comp-bk/domains/user/ports/drivers"
)

type DashboardService struct {
	forUser     user_ports_drivers.ForUser
	forLocation location_ports.ForLocation
}

func NewDashboardService(userAdapter *user_adap_drivers.UserProxyAdapter, locationAdapter location_ports.ForLocation) *DashboardService {
	return &DashboardService{
		forUser:     userAdapter,
		forLocation: locationAdapter,
	}
}

func (d *DashboardService) User() user_ports_drivers.ForUser {
	return d.forUser
}

func (d *DashboardService) Location() location_ports.ForLocation {
	return d.forLocation
}
