package routes

import (
	"PracticalProject/config"
	"PracticalProject/handlers"
	"PracticalProject/repositories"
	"PracticalProject/services"
	"github.com/gin-gonic/gin"
)

func InitProductRoutes(r *gin.RouterGroup) {
	productRepository := repositories.NewProductRepository(config.DB)
	productService := services.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService)

	products := r.Group("/products")
	{
		products.GET("/", productHandler.GetProduct)
		products.GET("/:id", productHandler.GetById)
		products.POST("/", productHandler.CreateProduct)
		products.PUT("/:id", productHandler.UpdateProduct)
		products.DELETE("/:id", productHandler.DeleteProduct)
	}

}
