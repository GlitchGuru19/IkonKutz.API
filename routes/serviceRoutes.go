package routes

import (
	"IkonKutz.API/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterServiceRoutes(api *gin.RouterGroup) {
	services := api.Group("/services")
	{
		services.GET("", controllers.GetServices)
		services.GET("/:id", controllers.GetService)
		services.POST("", controllers.CreateService)
		services.PUT("/:id", controllers.UpdateService)
		services.DELETE("/:id", controllers.DeleteService)
	}
}
