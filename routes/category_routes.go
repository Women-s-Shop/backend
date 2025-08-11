package routes

import (
	"PracticalProject/handlers"
	"github.com/gin-gonic/gin"
)

func InitCategoryRoutes(r *gin.RouterGroup) {
	categories := r.Group("/categories")
	{
		categories.GET("/categories", handlers.GetCategories)
		categories.GET("/categories/:id", handlers.GetCategoryByID)
		categories.POST("/categories", handlers.CreateCategory)
		categories.PUT("/categories/:id", handlers.UpdateCategory)
		categories.DELETE("/categories/:id", handlers.DeleteCategory)
	}
}
