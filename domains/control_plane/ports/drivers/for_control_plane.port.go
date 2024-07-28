package ports

import "github.com/DBrange/didis-comp-bk/domains/control_plane/ports/drivers/interfaces"

type ForControlPlane interface {
	interfaces.AuthenticationMiddleware
	interfaces.AuthorizationMiddleware
}