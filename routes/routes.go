package routes

import (
	"PracticalProject/handlers"
	"PracticalProject/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// CORS және preflight
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.OPTIONS("/*path", func(c *gin.Context) { c.Status(204) })

	// Пайдалы healthcheck обшем бұл тексерп көретнғо
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })

	// Public
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// Protected
	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())

	users := authorized.Group("/users")
	{
		users.GET("", handlers.GetUsers)
		users.GET("/:id", handlers.GetByIdUser)
		users.POST("", handlers.CreateUser)
		users.PUT("/:id", handlers.UpdateUser)
		users.DELETE("/:id", handlers.DeleteUser)
	}

	products := authorized.Group("/products")
	{
		products.GET("", handlers.GetProduct)
		products.GET("/:id", handlers.GetById)
		products.POST("", handlers.CreateProduct)
		products.PUT("/:id", handlers.UpdateProduct)
		products.DELETE("/:id", handlers.DeleteProduct)
	}

	carts := authorized.Group("/carts")
	{
		carts.GET("", handlers.GetCarts)
		carts.GET("/:id", handlers.GetCartByID)
		carts.POST("", handlers.CreateCart)
		carts.PUT("/:id", handlers.UpdateCart)
		carts.DELETE("/:id", handlers.DeleteCart)
	}

	// БҰРЫН: /categories/categories -> ЕНДІ: /categories
	categories := authorized.Group("/categories")
	{
		categories.GET("", handlers.GetCategories)
		categories.GET("/:id", handlers.GetCategoryByID)
		categories.POST("", handlers.CreateCategory)
		categories.PUT("/:id", handlers.UpdateCategory)
		categories.DELETE("/:id", handlers.DeleteCategory)
	}

	inventories := authorized.Group("/inventories")
	{
		inventories.GET("", handlers.GetInventories)
		inventories.GET("/:id", handlers.GetInventoryByID)
		inventories.POST("", handlers.CreateInventory)
		inventories.PUT("/:id", handlers.UpdateInventory)
		inventories.DELETE("/:id", handlers.DeleteInventory)
	}

	orders := authorized.Group("/orders")
	{
		orders.GET("", handlers.GetOrders)
		orders.GET("/:id", handlers.GetOrderByID)
		orders.POST("", handlers.CreateOrder)
		orders.PUT("/:id", handlers.UpdateOrder)
		orders.DELETE("/:id", handlers.DeleteOrder)
	}

	// Қалауыңша атауды /order-items қыла аламыз; әзірге сенікін қалдырдым
	orderItems := authorized.Group("/ordersItems")
	{
		orderItems.GET("", handlers.GetOrderItems)
		orderItems.GET("/:id", handlers.GetOrderItemByID)
		orderItems.POST("", handlers.CreateOrderItem)
		orderItems.PUT("/:id", handlers.UpdateOrderItem)
		orderItems.DELETE("/:id", handlers.DeleteOrderItem)
	}

	payments := authorized.Group("/payments")
	{
		payments.GET("", handlers.GetPayments)
		payments.GET("/:id", handlers.GetPaymentByID)
		payments.POST("", handlers.CreatePayment)
		payments.PUT("/:id", handlers.UpdatePayment)
		payments.DELETE("/:id", handlers.DeletePayment)
	}

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
