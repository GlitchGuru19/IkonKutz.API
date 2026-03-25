package routes

import (
	"IkonKutz.API/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterPublicServiceRoutes(api *gin.RouterGroup) {
	services := api.Group("/services")
	{
		services.GET("", controllers.GetServices)
		services.GET("/:id", controllers.GetService)
	}
}

func RegisterAdminServiceRoutes(api *gin.RouterGroup) {
	services := api.Group("/services")
	{
		services.POST("", controllers.CreateService)
		services.PUT("/:id", controllers.UpdateService)
		services.DELETE("/:id", controllers.DeleteService)
	}
}
