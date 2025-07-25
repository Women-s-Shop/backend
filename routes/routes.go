package routes

import (
	"PracticalProject/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	users := r.Group("/users")
	{
		users.GET("/", handlers.GetUsers)
		users.GET("/:id", handlers.GetByIdUser)
		users.POST("/", handlers.CreateUser)
		users.PUT("/:id", handlers.UpdateUser)
		users.DELETE("/:id", handlers.DeleteUser)
	}
	products := r.Group("/products")
	{
		products.GET("/", handlers.GetProduct)
		products.GET("/:id", handlers.GetById)
		products.POST("/", handlers.CreateProduct)
		products.PUT("/:id", handlers.UpdateProduct)
		products.DELETE("/:id", handlers.DeleteProduct)
	}

	carts := r.Group("/carts")
	{
		carts.GET("/", handlers.GetCarts)
		carts.GET("/:id", handlers.GetCartByID)
		carts.POST("/", handlers.CreateCart)
		carts.PUT("/:id", handlers.UpdateCart)
		carts.DELETE("/:id", handlers.DeleteCart)
	}

	categories := r.Group("/categories")
	{
		categories.GET("/categories", handlers.GetCategories)
		categories.GET("/categories/:id", handlers.GetCategoryByID)
		categories.POST("/categories", handlers.CreateCategory)
		categories.PUT("/categories/:id", handlers.UpdateCategory)
		categories.DELETE("/categories/:id", handlers.DeleteCategory)
	}
	inventories := r.Group("/inventories")
	{
		inventories.GET("/", handlers.GetInventories)
		inventories.GET("/:id", handlers.GetInventoryByID)
		inventories.POST("/", handlers.CreateInventory)
		inventories.PUT("/:id", handlers.UpdateInventory)
		inventories.DELETE("/:id", handlers.DeleteInventory)
	}

	orders := r.Group("/orders")
	{
		orders.GET("/", handlers.GetOrders)
		orders.GET("/:id", handlers.GetOrderByID)
		orders.POST("/", handlers.CreateOrder)
		orders.PUT("/:id", handlers.UpdateOrder)
		orders.DELETE("/:id", handlers.DeleteOrder)
	}

	orderItems := r.Group("/ordersItems")
	{
		orderItems.GET("/", handlers.GetOrderItems)
		orderItems.GET("/:id", handlers.GetOrderItemByID)
		orderItems.POST("/", handlers.CreateOrderItem)
		orderItems.PUT("/:id", handlers.UpdateOrderItem)
		orderItems.DELETE("/:id", handlers.DeleteOrderItem)
	}

	payments := r.Group("/payments")
	{
		payments.GET("/", handlers.GetPayments)
		products.GET("/:id", handlers.GetPaymentByID)
		payments.POST("/", handlers.CreatePayment)
		payments.PUT("/:id", handlers.UpdatePayment)
		payments.DELETE("/:id", handlers.DeletePayment)
	}

	promocodes := r.Group("/promocodes")
	{
		promocodes.GET("/", handlers.GetPromocodes)
		promocodes.GET("/:id", handlers.GetPromocodeByID)
		promocodes.POST("/", handlers.CreatePromocode)
		promocodes.PUT("/:id", handlers.UpdatePromocode)
		promocodes.DELETE("/:id", handlers.DeletePromocode)
	}

	return r
}
