package routes

import (
	"PracticalProject/handlers"
	"github.com/gin-gonic/gin"
)

func InitInventoryRoutes(r *gin.RouterGroup) {
	inventories := r.Group("/inventories")
	{
		inventories.GET("/", handlers.GetInventories)
		inventories.GET("/:id", handlers.GetInventoryByID)
		inventories.POST("/", handlers.CreateInventory)
		inventories.PUT("/:id", handlers.UpdateInventory)
		inventories.DELETE("/:id", handlers.DeleteInventory)
	}
}
