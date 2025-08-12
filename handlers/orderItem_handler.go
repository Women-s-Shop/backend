package handlers

import (
	"PracticalProject/models"
	"PracticalProject/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderItemHandler struct {
	service *services.OrderItemService
}

func NewOrderItemHandler(s *services.OrderItemService) *OrderItemHandler {
	return &OrderItemHandler{service: s}
}

func (h *OrderItemHandler) GetOrderItems(c *gin.Context) {
	items, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get order items",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *OrderItemHandler) CreateOrderItem(c *gin.Context) {
	var item models.OrderItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := h.service.Create(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Order item creation failed",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func (h *OrderItemHandler) GetOrderItemByID(c *gin.Context) {
	id := c.Param("id")

	item, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Order item not found",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *OrderItemHandler) UpdateOrderItem(c *gin.Context) {
	id := c.Param("id")

	existing, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Order item not found",
			"error":   err.Error(),
		})
		return
	}

	var update models.OrderItem
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON error"})
		return
	}
	update.ID = existing.ID

	updated, err := h.service.Update(update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Order item update failed",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (h *OrderItemHandler) DeleteOrderItem(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Order item not found",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order item deleted successfully"})
}
