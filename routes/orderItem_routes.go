package routes

import (
	"PracticalProject/handlers"
	"github.com/gin-gonic/gin"
)

func InitOrderItemRoutes(r *gin.RouterGroup) {
	orderItems := r.Group("/ordersItems")
	{
		orderItems.GET("/", handlers.GetOrderItems)
		orderItems.GET("/:id", handlers.GetOrderItemByID)
		orderItems.POST("/", handlers.CreateOrderItem)
		orderItems.PUT("/:id", handlers.UpdateOrderItem)
		orderItems.DELETE("/:id", handlers.DeleteOrderItem)
	}
}
