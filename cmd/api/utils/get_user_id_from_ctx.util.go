package utils

import "github.com/gin-gonic/gin"

func GetUserIDFromContext(c *gin.Context) string {
	userID, ok := c.Get("userID")
	if !ok {
		return ""
	}

	if id, ok := userID.(string); ok {
		return id
	}

	return ""
}