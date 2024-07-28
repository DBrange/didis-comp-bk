package ports

type ForQueryingControlPlane interface {
	// AuthenticationMiddleware() gin.HandlerFunc
	// AuthorizationMiddleware(requiredRole ...models.ROLE) gin.HandlerFunc
}
