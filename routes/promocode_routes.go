package routes

import (
	"PracticalProject/handlers"
	"github.com/gin-gonic/gin"
)

func InitPromocodeRoutes(r *gin.RouterGroup) {
	promocodes := r.Group("/promocodes")
	{
		promocodes.GET("/", handlers.GetPromocodes)
		promocodes.GET("/:id", handlers.GetPromocodeByID)
		promocodes.POST("/", handlers.CreatePromocode)
		promocodes.PUT("/:id", handlers.UpdatePromocode)
		promocodes.DELETE("/:id", handlers.DeletePromocode)
	}

}
