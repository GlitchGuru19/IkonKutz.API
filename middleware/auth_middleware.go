package middleware

import (
	"net/http"
	"strings"

	"IkonKutz.API/utils"
	"github.com/gin-gonic/gin"
)

// function that returns a Gin middleware function to require authentication using JWT tokens
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.Error(c, http.StatusUnauthorized, "Authorization header is required")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" || strings.TrimSpace(parts[1]) == "" {
			utils.Error(c, http.StatusUnauthorized, "Authorization header must be in the format: Bearer <token>")
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			utils.Error(c, http.StatusUnauthorized, "Invalid or expired token")
			c.Abort()
			return
		}

		// Store important token data in Gin context so controllers can use it.
		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)
		c.Set("email", claims.Email)

		c.Next()
	}
}
