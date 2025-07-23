package handlers

import (
	"PracticalProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var users = []models.User{
	{Id: 1, Name: "Beibit", Email: "beibit.muzappar@narxoz.kz", Password: "bbb", Phone: "87072960605", Address: "T60"},
	{Id: 2, Name: "Yersin", Email: "yersin.bizak@narxoz.kz", Password: "yyy", Phone: "87054562135", Address: "R41"},
}

func GetUsers(c *gin.Context) {
	c.JSON(200, users)
}
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	user.Id = len(users) + 1
	users = append(users, user)
	c.JSON(http.StatusCreated, user)
}

func GetByIdUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}

	for _, u := range users {
		if u.Id == id {
			c.JSON(200, u)
			return
		}
	}

	c.JSON(404, gin.H{"Error": "Not found"})
}

func UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "There is no such id"})
		return
	}

	var updateUser models.User
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(400, gin.H{"Error": "JSON error"})
		return
	}

	for i, u := range users {
		if u.Id == id {
			users[i] = updateUser
			c.JSON(200, updateUser)
			return
		}
	}

	c.JSON(404, gin.H{"Error": "Not found"})
}

func DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(400, gin.H{"Error": "There is no such id"})
		return
	}
	for i, u := range users {
		if u.Id == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(200, gin.H{"message": "Successfully deleted"})
			return
		}
	}

	c.JSON(404, gin.H{"Error": "Not found"})
}
