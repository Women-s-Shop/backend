package routes

import (
	"PracticalProject/handlers"
	"PracticalProject/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Публичные роуты
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// Защищённые роуты
	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())

	// 🔹 Текущий пользователь по токену
	authorized.GET("/me", handlers.GetMe)

	// Users
	users := authorized.Group("/users")
	{
		users.GET("", handlers.GetUsers)          // GET /users
		users.GET("/:id", handlers.GetByIdUser)   // GET /users/:id
		users.POST("", handlers.CreateUser)       // POST /users
		users.PUT("/:id", handlers.UpdateUser)    // PUT /users/:id
		users.DELETE("/:id", handlers.DeleteUser) // DELETE /users/:id
	}

	// Products
	products := authorized.Group("/products")
	{
		products.GET("", handlers.GetProduct)
		products.GET("/:id", handlers.GetById)
		products.POST("", handlers.CreateProduct)
		products.PUT("/:id", handlers.UpdateProduct)
		products.DELETE("/:id", handlers.DeleteProduct)
	}

	// Carts
	carts := authorized.Group("/carts")
	{
		carts.GET("", handlers.GetCarts)
		carts.GET("/:id", handlers.GetCartByID)
		carts.POST("", handlers.CreateCart)
		carts.PUT("/:id", handlers.UpdateCart)
		carts.DELETE("/:id", handlers.DeleteCart)
	}

	// ✅ Categories (убрал двойные префиксы)
	categories := authorized.Group("/categories")
	{
		categories.GET("", handlers.GetCategories)
		categories.GET("/:id", handlers.GetCategoryByID)
		categories.POST("", handlers.CreateCategory)
		categories.PUT("/:id", handlers.UpdateCategory)
		categories.DELETE("/:id", handlers.DeleteCategory)
	}

	// Inventories
	inventories := authorized.Group("/inventories")
	{
		inventories.GET("", handlers.GetInventories)
		inventories.GET("/:id", handlers.GetInventoryByID)
		inventories.POST("", handlers.CreateInventory)
		inventories.PUT("/:id", handlers.UpdateInventory)
		inventories.DELETE("/:id", handlers.DeleteInventory)
	}

	// Orders
	orders := authorized.Group("/orders")
	{
		orders.GET("", handlers.GetOrders)
		orders.GET("/:id", handlers.GetOrderByID)
		orders.POST("", handlers.CreateOrder)
		orders.PUT("/:id", handlers.UpdateOrder)
		orders.DELETE("/:id", handlers.DeleteOrder)
	}

	// Order Items
	orderItems := authorized.Group("/ordersItems")
	{
		orderItems.GET("", handlers.GetOrderItems)
		orderItems.GET("/:id", handlers.GetOrderItemByID)
		orderItems.POST("", handlers.CreateOrderItem)
		orderItems.PUT("/:id", handlers.UpdateOrderItem)
		orderItems.DELETE("/:id", handlers.DeleteOrderItem)
	}

	// Payments
	payments := authorized.Group("/payments")
	{
		payments.GET("", handlers.GetPayments)
		payments.GET("/:id", handlers.GetPaymentByID)
		payments.POST("", handlers.CreatePayment)
		payments.PUT("/:id", handlers.UpdatePayment)
		payments.DELETE("/:id", handlers.DeletePayment)
	}

	// Promocodes
	promocodes := authorized.Group("/promocodes")
	{
		promocodes.GET("", handlers.GetPromocodes)
		promocodes.GET("/:id", handlers.GetPromocodeByID)
		promocodes.POST("", handlers.CreatePromocode)
		promocodes.PUT("/:id", handlers.UpdatePromocode)
		promocodes.DELETE("/:id", handlers.DeletePromocode)
	}

	return r
}
