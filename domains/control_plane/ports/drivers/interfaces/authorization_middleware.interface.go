package interfaces

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/gin-gonic/gin"
)

type AuthorizationMiddleware interface {
	AuthorizationMiddleware(requiredRole ...models.ROLE) gin.HandlerFunc
}
