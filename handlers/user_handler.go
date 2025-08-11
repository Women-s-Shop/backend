package handlers

import (
	"PracticalProject/models"
	"PracticalProject/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	var userService services.UserService
	users, err := userService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	var userService services.UserService
	createdUser, err := userService.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, createdUser)
}

func GetByIdUser(c *gin.Context) {
	idParam := c.Param("id")

	var userService services.UserService
	user, err := userService.GetByID(idParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
func UpdateUser(c *gin.Context) {
	idParam := c.Param("id")

	var userService services.UserService
	user, err := userService.GetByID(idParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "User not found"})
		return
	}

	var updateUser models.User
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "JSON error"})
		return
	}
	updateUser.ID = user.ID
	updatedUser, err := userService.Update(updateUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func DeleteUser(c *gin.Context) {
	idParam := c.Param("id")

	var userService services.UserService
	err := userService.Delete(idParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
