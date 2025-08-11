package routes

import (
	"PracticalProject/handlers"
	"github.com/gin-gonic/gin"
)

func InitPaymentRoutes(r *gin.RouterGroup) {
	payments := r.Group("/payments")
	{
		payments.GET("/", handlers.GetPayments)
		payments.GET("/:id", handlers.GetPaymentByID)
		payments.POST("/", handlers.CreatePayment)
		payments.PUT("/:id", handlers.UpdatePayment)
		payments.DELETE("/:id", handlers.DeletePayment)
	}
}
