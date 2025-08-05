package handlers

import (
	"PracticalProject/models"
	"PracticalProject/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CartHandler struct {
	service services.CartService
}

func NewCartHandler(service services.CartService) CartHandler {
	return CartHandler{service}
}

func (h *CartHandler) GetCarts(c *gin.Context) {
	carts, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(200, carts)
}

func (h *CartHandler) CreateCart(c *gin.Context) {
	var cart models.Cart
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	if err := h.service.CreateCart(&cart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cart creation failed",
			"Error":   err.Error()})
		return
	}
	c.JSON(http.StatusCreated, cart)
}

func (h *CartHandler) GetCartByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid cart Id",
			"Error":   err.Error(),
		})
		return
	}
	cart, err := h.service.GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Cart not found",
			"Error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, cart)
}

func (h *CartHandler) UpdateCart(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid cart ID",
			"Error":   err.Error(),
		})
		return
	}
	cart, err := h.service.GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Cart not found",
			"Error":   err.Error(),
		})
		return
	}
	if err := c.ShouldBindJSON(cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := h.service.UpdateCart(cart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cart)
}

func (h *CartHandler) DeleteCart(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid cart ID",
			"Error":   err.Error(),
		})
		return
	}
	cart, err := h.service.GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Cart not found",
			"Error":   err.Error(),
		})
		return
	}

	if err := h.service.DeleteCart(cart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart deleted successfully"})
}
