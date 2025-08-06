package routes

import (
	"PracticalProject/handlers"
	"PracticalProject/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	InitAuthRoutes(r)

	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())

	InitProductRoutes(authorized)

	users := authorized.Group("/users")
	{
		users.GET("/", handlers.GetUsers)
		users.GET("/:id", handlers.GetByIdUser)
		users.POST("/", handlers.CreateUser)
		users.PUT("/:id", handlers.UpdateUser)
		users.DELETE("/:id", handlers.DeleteUser)
	}

	InitCartRoutes(authorized)

	categories := authorized.Group("/categories")
	{
		categories.GET("/categories", handlers.GetCategories)
		categories.GET("/categories/:id", handlers.GetCategoryByID)
		categories.POST("/categories", handlers.CreateCategory)
		categories.PUT("/categories/:id", handlers.UpdateCategory)
		categories.DELETE("/categories/:id", handlers.DeleteCategory)
	}
	inventories := authorized.Group("/inventories")
	{
		inventories.GET("/", handlers.GetInventories)
		inventories.GET("/:id", handlers.GetInventoryByID)
		inventories.POST("/", handlers.CreateInventory)
		inventories.PUT("/:id", handlers.UpdateInventory)
		inventories.DELETE("/:id", handlers.DeleteInventory)
	}

	orders := authorized.Group("/orders")
	{
		orders.GET("/", handlers.GetOrders)
		orders.GET("/:id", handlers.GetOrderByID)
		orders.POST("/", handlers.CreateOrder)
		orders.PUT("/:id", handlers.UpdateOrder)
		orders.DELETE("/:id", handlers.DeleteOrder)
	}

	orderItems := authorized.Group("/ordersItems")
	{
		orderItems.GET("/", handlers.GetOrderItems)
		orderItems.GET("/:id", handlers.GetOrderItemByID)
		orderItems.POST("/", handlers.CreateOrderItem)
		orderItems.PUT("/:id", handlers.UpdateOrderItem)
		orderItems.DELETE("/:id", handlers.DeleteOrderItem)
	}

	payments := authorized.Group("/payments")
	{
		payments.GET("/", handlers.GetPayments)
		payments.GET("/:id", handlers.GetPaymentByID)
		payments.POST("/", handlers.CreatePayment)
		payments.PUT("/:id", handlers.UpdatePayment)
		payments.DELETE("/:id", handlers.DeletePayment)
	}

	promocodes := authorized.Group("/promocodes")
	{
		promocodes.GET("/", handlers.GetPromocodes)
		promocodes.GET("/:id", handlers.GetPromocodeByID)
		promocodes.POST("/", handlers.CreatePromocode)
		promocodes.PUT("/:id", handlers.UpdatePromocode)
		promocodes.DELETE("/:id", handlers.DeletePromocode)
	}

	return r
}
