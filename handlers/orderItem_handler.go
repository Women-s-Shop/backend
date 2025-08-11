package handlers

import (
	"PracticalProject/models"
	"PracticalProject/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOrderItems(c *gin.Context) {
	var orderItemService services.OrderService
	orderItems, err := orderItemService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON",
			"Error":   err.Error(),
		})
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
	var orderItemService services.OrderItemService
	createOrderItem, err := orderItemService.Create(orderItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Order Item creation filed",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createOrderItem)
}

func GetOrderItemByID(c *gin.Context) {
	var orderItemService services.OrderItemService
	id := c.Param("id")
	orderItem, err := orderItemService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "OrderItem not found",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, orderItem)
}

func UpdateOrderItem(c *gin.Context) {
	id := c.Param("id")
	var orderItemService services.OrderItemService

	orderItem, err := orderItemService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "OrderItem not found",
			"error":   err.Error(),
		})
		return
	}
	var updatedOrderItem models.OrderItem
	if err := c.ShouldBindJSON(&updatedOrderItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	updatedOrderItem.ID = orderItem.ID
	updateItem, err := orderItemService.Update(updatedOrderItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "OrderItem updating filed",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, updateItem)
}

func DeleteOrderItem(c *gin.Context) {
	var orderItemService services.OrderItemService
	id := c.Param("id")
	err := orderItemService.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "OrderItem not found",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OrderItem deleted successfully"})
}
