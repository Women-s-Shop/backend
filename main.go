package main

import (
	"PracticalProject/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	fmt.Println("Hello World")
	r := gin.Default()

	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetByIdUser)
	r.POST("/users", handlers.CreateUser)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	r.GET("/products", handlers.GetProduct)
	r.GET("/products/:id", handlers.GetById)
	r.POST("/products", handlers.CreateProduct)
	r.PUT("/products/:id", handlers.UpdateProduct)
	r.DELETE("/products/:id", handlers.DeleteProduct)

	r.Run(":8888")
}
