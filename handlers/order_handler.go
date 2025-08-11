package handlers

import (
	"PracticalProject/models"
	"PracticalProject/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderHandler struct {
	service services.OrderService
}

func NewOrderHandler(service services.OrderService) OrderHandler {
	return OrderHandler{service}
}

func (h *OrderHandler) GetOrders(c *gin.Context) {

	orders, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	if err := h.service.CreateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Order creation failed",
			"Error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	order, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Order not found",
			"Error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	id := c.Param("id")

	var fields map[string]interface{}
	if err := c.ShouldBindJSON(&fields); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateOrder(id, fields); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})
}

func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteOrder(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Order not found",
			"Error":   err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
