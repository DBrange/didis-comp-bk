package dashboard

import user_ports_drivers "didis-comp-bk/internal/user/ports/drivers"

type Dashboard interface {
	User() user_ports_drivers.ForUser
}
