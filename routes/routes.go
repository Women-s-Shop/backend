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

	return r
}
