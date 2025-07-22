package main

import (
	"PracticalProject/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	fmt.Println("Hello World")
	r := gin.Default()

	r.GET("/users", GetUsers)
	r.POST("/users", OnOrOff)

	r.Run(":8888")
}

var users = []models.User{
	{Id: 1, Name: "Beibit", Email: "beibit.muzappar@narxoz.kz", Password: "bbb", Phone: "87072960605", Address: "T60"},
	{Id: 2, Name: "Yersin", Email: "yersin.bizak@narxoz.kz", Password: "yyy", Phone: "87054562135", Address: "R41"},
}

func GetUsers(c *gin.Context) {
	c.JSON(500, users)
}
func OnOrOff(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	user.Id = len(users) + 1
	users = append(users, user)
	c.JSON(http.StatusCreated, user)
}
