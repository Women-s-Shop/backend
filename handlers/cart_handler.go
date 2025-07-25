package handlers

import (
	"PracticalProject/config"
	"PracticalProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCarts(c *gin.Context) {
	var carts []models.Cart
	if err := config.DB.Find(&carts).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(200, carts)
}

func CreateCart(c *gin.Context) {
	var cart models.Cart
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	if err := config.DB.Create(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cart)
}

func GetCartByID(c *gin.Context) {
	var cart models.Cart
	id := c.Param("id")
	if err := config.DB.First(&cart, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Cart not found"})
		return
	}
	c.JSON(http.StatusOK, cart)
}

func UpdateCart(c *gin.Context) {
	var cart models.Cart
	id := c.Param("id")

	if err := config.DB.First(&cart, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Cart not found"})
		return
	}

	var updatedCart models.Cart
	if err := c.ShouldBindJSON(&updatedCart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := config.DB.Model(&cart).Updates(updatedCart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cart)
}

func DeleteCart(c *gin.Context) {
	var cart models.Cart
	id := c.Param("id")

	if err := config.DB.First(&cart, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Cart not found"})
		return
	}

	if err := config.DB.Delete(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart deleted successfully"})
}
