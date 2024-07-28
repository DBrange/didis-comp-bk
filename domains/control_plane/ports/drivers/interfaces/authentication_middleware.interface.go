package interfaces

import "github.com/gin-gonic/gin"

type AuthenticationMiddleware interface {
	AuthenticationMiddleware() gin.HandlerFunc
}
