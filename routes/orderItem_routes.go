package routes

import (
	"PracticalProject/config"
	"PracticalProject/handlers"
	"PracticalProject/repositories"
	"PracticalProject/services"
	"github.com/gin-gonic/gin"
)

func InitOrderItemRoutes(r *gin.RouterGroup) {

	OrderItemRepo := repositories.NewOrderItemRepository(config.DB)
	OrderItemService := services.NewOrderItemService(OrderItemRepo)
	OrderItemHandler := handlers.NewOrderItemHandler(OrderItemService)

	orderItems := r.Group("/ordersItems")
	{
		orderItems.GET("/", OrderItemHandler.GetOrderItems)
		orderItems.GET("/:id", OrderItemHandler.GetOrderItemByID)
		orderItems.POST("/", OrderItemHandler.CreateOrderItem)
		orderItems.PUT("/:id", OrderItemHandler.UpdateOrderItem)
		orderItems.DELETE("/:id", OrderItemHandler.DeleteOrderItem)
	}
}
