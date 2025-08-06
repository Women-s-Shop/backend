package routes

import (
	"PracticalProject/config"
	"PracticalProject/handlers"
	"PracticalProject/repositories"
	"PracticalProject/services"
	"github.com/gin-gonic/gin"
)

func InitCartRoutes(r *gin.RouterGroup) {

	cartRepository := repositories.NewCartRepository(config.DB)
	cartService := services.NewCartService(cartRepository)
	cartHandler := handlers.NewCartHandler(cartService)
	carts := r.Group("/carts")
	{
		carts.GET("/", cartHandler.GetCarts)
		carts.GET("/:id", cartHandler.GetCartByID)
		carts.POST("/", cartHandler.CreateCart)
		carts.PUT("/:id", cartHandler.UpdateCart)
		carts.DELETE("/:id", cartHandler.DeleteCart)
	}
}
