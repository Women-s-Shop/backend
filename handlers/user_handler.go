package handlers

import (
	"PracticalProject/config"
	"PracticalProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(200, users)
}
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(200, user)
}

func GetByIdUser(c *gin.Context) {
	var users models.User
	idParam := c.Param("id")
	if err := config.DB.First(&users, idParam).Error; err != nil {
		c.JSON(404, gin.H{"Error": "There is no such id"})
		return
	}

	c.JSON(200, gin.H{"This user": users})
}

func UpdateUser(c *gin.Context) {
	var users models.User
	idParam := c.Param("id")

	if err := config.DB.First(&users, idParam).Error; err != nil {
		c.JSON(404, gin.H{"Error": "User not found"})
		return
	}

	var updateUser models.User
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(400, gin.H{"Error": "JSON error"})
		return
	}

	config.DB.Model(&users).Updates(updateUser)
	c.JSON(200, users)
}

func DeleteUser(c *gin.Context) {
	var users models.User
	idParam := c.Param("id")

	if err := config.DB.First(&users, idParam).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "User not found"})
		return
	}
	config.DB.Delete(&users)

	c.JSON(200, gin.H{"message": "User deleted successfully!!!"})
}
