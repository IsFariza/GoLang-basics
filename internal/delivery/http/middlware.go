package http

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		role := session.Get("role")

		if session.Get("userID") == nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		userRole := role.(string)

		if userRole == "admin" {
			c.Next()
			return
		}

		if requiredRole == "moderator" && userRole != "moderator" {
			c.AbortWithStatusJSON(403, gin.H{"error": "Moderator access only"})
			return
		}

		c.Next()
	}
}
