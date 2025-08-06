package routes

import (
	"PracticalProject/config"
	"PracticalProject/handlers"
	"PracticalProject/repositories"
	"PracticalProject/services"
	"github.com/gin-gonic/gin"
)

func InitAuthRoutes(r *gin.Engine) {
	authRepository := repositories.NewAuthRepository(config.DB)
	authService := services.NewAuthService(authRepository)
	authHandler := handlers.NewAuthHandler(authService)

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

}
