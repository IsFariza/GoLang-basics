package http

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("userID")
		userRole := session.Get("role")
		if userID == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. Please login."})
			c.Abort()
			return
		}

		if requiredRole == "admin" && userRole != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Admin access only"})
			c.Abort()
			return
		}

		c.Set("currentUserID", userID)
		c.Set("currentUserRole", userRole)
		c.Next()
	}
}
