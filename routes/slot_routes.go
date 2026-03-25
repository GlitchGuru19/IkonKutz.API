package routes

import (
	"IkonKutz.API/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterPublicSlotRoutes(api *gin.RouterGroup) {
	slots := api.Group("/slots")
	{
		slots.GET("", controllers.GetSlots)
		slots.GET("/:id", controllers.GetSlot)
	}
}

func RegisterAdminSlotRoutes(api *gin.RouterGroup) {
	slots := api.Group("/slots")
	{
		slots.POST("", controllers.CreateSlot)
		slots.PUT("/:id", controllers.UpdateSlot)
		slots.PATCH("/:id/lock", controllers.LockSlot)
		slots.PATCH("/:id/unlock", controllers.UnlockSlot)
		slots.DELETE("/:id", controllers.DeleteSlot)
	}
}
