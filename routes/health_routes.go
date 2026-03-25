package routes

import (
	"IkonKutz.API/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterHealthRoutes(r *gin.Engine) {
	r.GET("/api/health", controllers.Health)
}
