package routes

import (
	"PracticalProject/config"
	"PracticalProject/handlers"
	"PracticalProject/repositories"
	"PracticalProject/services"
	"github.com/gin-gonic/gin"
)

func InitOrderRoutes(r *gin.RouterGroup) {
	orderRepo := repositories.NewOrderRepository(config.DB)
	orderServ := services.NewOrderService(orderRepo)
	orderHand := handlers.NewOrderHandler(orderServ)

	orders := r.Group("/orders")
	{
		orders.GET("/", orderHand.GetOrders)
		orders.GET("/:id", orderHand.GetOrderByID)
		orders.POST("/", orderHand.CreateOrder)
		orders.PUT("/:id", orderHand.UpdateOrder)
		orders.DELETE("/:id", orderHand.DeleteOrder)
	}
}
