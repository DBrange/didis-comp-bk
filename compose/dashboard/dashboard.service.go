package dashboard

import (
	user_adap_drivers "github.com/DBrange/didis-comp-bk/internal/user/adapters/drivers"
	user_ports_drivers "github.com/DBrange/didis-comp-bk/internal/user/ports/drivers"
)

type DashboardService struct {
	ForUser user_ports_drivers.ForUser
}

func NewDashboardService(UserAdater *user_adap_drivers.UserAdapter) *DashboardService {
	return &DashboardService{
		ForUser: UserAdater,
	}
}

func (d *DashboardService) User() user_ports_drivers.ForUser {
	return d.ForUser
}
