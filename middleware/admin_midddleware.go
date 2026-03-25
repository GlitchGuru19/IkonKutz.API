package middleware

import (
	"net/http"

	"IkonKutz.API/utils"
	"github.com/gin-gonic/gin"
)

// function that returns a Gin middleware function to require admin role
func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleValue, exists := c.Get("role")
		if !exists {
			utils.Error(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		role, ok := roleValue.(string)
		if !ok || role != "admin" {
			utils.Error(c, http.StatusForbidden, "Admin access required")
			c.Abort()
			return
		}

		c.Next()
	}
}
