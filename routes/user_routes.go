package routes

import (
	"PracticalProject/handlers"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.GET("/", handlers.GetUsers)
		users.GET("/:id", handlers.GetByIdUser)
		users.POST("/", handlers.CreateUser)
		users.PUT("/:id", handlers.UpdateUser)
		users.DELETE("/:id", handlers.DeleteUser)
	}

}
