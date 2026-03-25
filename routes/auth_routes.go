package routes

import (
	"IkonKutz.API/controllers"
	"IkonKutz.API/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)

		authenticated := auth.Group("")
		authenticated.Use(middleware.RequireAuth())
		{
			authenticated.GET("/me", controllers.Me)
			authenticated.POST("/logout", controllers.Logout)
		}
	}
}
