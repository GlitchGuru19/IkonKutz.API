package routes

import (
	"IkonKutz.API/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterProtectedAppointmentRoutes(api *gin.RouterGroup) {
	appointments := api.Group("/appointments")
	{
		appointments.GET("/me", controllers.GetMyAppointments)
		appointments.GET("/:id", controllers.GetAppointment)
		appointments.POST("", controllers.CreateAppointment)
		appointments.PUT("/:id", controllers.UpdateAppointment)
		appointments.PATCH("/:id/cancel", controllers.CancelAppointment)
		appointments.DELETE("/:id", controllers.DeleteAppointment)
	}
}

func RegisterAdminAppointmentRoutes(api *gin.RouterGroup) {
	appointments := api.Group("/appointments")
	{
		appointments.GET("", controllers.GetAppointments)
	}
}
