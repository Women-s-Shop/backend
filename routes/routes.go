package routes

import (
	"PracticalProject/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	InitAuthRoutes(r)

	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())

	InitProductRoutes(authorized)
	InitUserRoutes(authorized)
	InitCartRoutes(authorized)
	InitCategoryRoutes(authorized)
	InitInventoryRoutes(authorized)
	InitOrderRoutes(authorized)
	InitOrderItemRoutes(authorized)
	InitPaymentRoutes(authorized)
	InitPromocodeRoutes(authorized)

	return r
}
