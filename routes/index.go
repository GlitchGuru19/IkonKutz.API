package routes

import (
	"IkonKutz.API/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	RegisterHealthRoutes(r)

	api := r.Group("/api")

	// Public routes
	RegisterAuthRoutes(api)
	RegisterPublicServiceRoutes(api)
	RegisterPublicSlotRoutes(api)

	// Protected routes for any authenticated user
	protected := api.Group("")
	protected.Use(middleware.RequireAuth())
	{
		RegisterProtectedAppointmentRoutes(protected)
	}

	// Admin-only routes
	admin := api.Group("")
	admin.Use(middleware.RequireAuth(), middleware.RequireAdmin())
	{
		RegisterAdminServiceRoutes(admin)
		RegisterAdminSlotRoutes(admin)
		RegisterAdminAppointmentRoutes(admin)
	}
}
