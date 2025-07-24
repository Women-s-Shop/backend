package handlers

import (
	"PracticalProject/config"
	"PracticalProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOrderItems(c *gin.Context) {
	var orderItems []models.OrderItem
	if err := config.DB.Find(&orderItems).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orderItems)
}

func CreateOrderItem(c *gin.Context) {
	var orderItem models.OrderItem
	if err := c.ShouldBindJSON(&orderItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	if err := config.DB.Create(&orderItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orderItem)
}

func GetOrderItemByID(c *gin.Context) {
	var orderItem models.OrderItem
	id := c.Param("id")
	if err := config.DB.First(&orderItem, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "OrderItem not found"})
		return
	}
	c.JSON(http.StatusOK, orderItem)
}

func UpdateOrderItem(c *gin.Context) {
	var orderItem models.OrderItem
	id := c.Param("id")

	if err := config.DB.First(&orderItem, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "OrderItem not found"})
		return
	}

	var updatedOrderItem models.OrderItem
	if err := c.ShouldBindJSON(&updatedOrderItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := config.DB.Model(&orderItem).Updates(updatedOrderItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orderItem)
}

func DeleteOrderItem(c *gin.Context) {
	var orderItem models.OrderItem
	id := c.Param("id")

	if err := config.DB.First(&orderItem, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "OrderItem not found"})
		return
	}

	if err := config.DB.Delete(&orderItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OrderItem deleted successfully"})
}
