package main

import (
	"PracticalProject/config"
	"PracticalProject/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	config.InitDB()

	fmt.Println("Hello World")
	r := gin.Default()

	r.GET("/carts", handlers.GetCarts)
	r.GET("/carts/:id", handlers.GetCartByID)
	r.POST("/carts", handlers.CreateCart)
	r.PUT("/carts/:id", handlers.UpdateCart)
	r.DELETE("/carts/:id", handlers.DeleteCart)

	r.GET("/categories", handlers.GetCategories)
	r.GET("/categories/:id", handlers.GetCategoryByID)
	r.POST("/categories", handlers.CreateCategory)
	r.PUT("/categories/:id", handlers.UpdateCategory)
	r.DELETE("/categories/:id", handlers.DeleteCategory)

	r.GET("/inventories", handlers.GetInventories)
	r.GET("/inventories/:id", handlers.GetInventoryByID)
	r.POST("/inventories", handlers.CreateInventory)
	r.PUT("/inventories/:id", handlers.UpdateInventory)
	r.DELETE("/inventories/:id", handlers.DeleteInventory)

	r.GET("/orders", handlers.GetOrders)
	r.GET("/orders/:id", handlers.GetOrderByID)
	r.POST("/orders", handlers.CreateOrder)
	r.PUT("/orders/:id", handlers.UpdateOrder)
	r.DELETE("/orders/:id", handlers.DeleteOrder)

	r.GET("/order-items", handlers.GetOrderItems)
	r.GET("/order-items/:id", handlers.GetOrderItemByID)
	r.POST("/order-items", handlers.CreateOrderItem)
	r.PUT("/order-items/:id", handlers.UpdateOrderItem)
	r.DELETE("/order-items/:id", handlers.DeleteOrderItem)

	r.GET("/payments", handlers.GetPayments)
	r.GET("/payments/:id", handlers.GetPaymentByID)
	r.POST("/payments", handlers.CreatePayment)
	r.PUT("/payments/:id", handlers.UpdatePayment)
	r.DELETE("/payments/:id", handlers.DeletePayment)

	r.GET("/promocodes", handlers.GetPromocodes)
	r.GET("/promocodes/:id", handlers.GetPromocodeByID)
	r.POST("/promocodes", handlers.CreatePromocode)
	r.PUT("/promocodes/:id", handlers.UpdatePromocode)
	r.DELETE("/promocodes/:id", handlers.DeletePromocode)

	r.Run(":8888")
}
