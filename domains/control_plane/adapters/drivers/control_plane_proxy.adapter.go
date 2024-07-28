package adapters

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/control_plane/services"
	"github.com/gin-gonic/gin"
)

type ControlPlaneProxyAdapter struct {
	controlPlaneService *services.ControlPlaneService
}

func NewControlPlaneProxyAdapter(controlPlaneService *services.ControlPlaneService) *ControlPlaneProxyAdapter {
	return &ControlPlaneProxyAdapter{
		controlPlaneService: controlPlaneService,
	}
}

func (a *ControlPlaneProxyAdapter) AuthenticationMiddleware() gin.HandlerFunc {
	return a.controlPlaneService.AuthenticationMiddleware()
}

func (a *ControlPlaneProxyAdapter) AuthorizationMiddleware(requiredRole ...models.ROLE) gin.HandlerFunc {
	return a.controlPlaneService.AuthorizationMiddleware(requiredRole)
}
