package routes

import (
	"PracticalProject/config"
	"PracticalProject/handlers"
	"PracticalProject/middleware"
	"PracticalProject/repasitories"
	"PracticalProject/services"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	productRepository := repasitories.NewProductRepository(config.DB)
	productService := services.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService)

	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())

	users := authorized.Group("/users")
	{
		users.GET("/", handlers.GetUsers)
		users.GET("/:id", handlers.GetByIdUser)
		users.POST("/", handlers.CreateUser)
		users.PUT("/:id", handlers.UpdateUser)
		users.DELETE("/:id", handlers.DeleteUser)
	}
	products := authorized.Group("/products")
	{
		products.GET("/", productHandler.GetProduct)
		products.GET("/:id", productHandler.GetById)
		products.POST("/", productHandler.CreateProduct)
		products.PUT("/:id", productHandler.UpdateProduct)
		products.DELETE("/:id", productHandler.DeleteProduct)
	}

	carts := authorized.Group("/carts")
	{
		carts.GET("/", handlers.GetCarts)
		carts.GET("/:id", handlers.GetCartByID)
		carts.POST("/", handlers.CreateCart)
		carts.PUT("/:id", handlers.UpdateCart)
		carts.DELETE("/:id", handlers.DeleteCart)
	}

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
