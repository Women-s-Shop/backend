package handlers

import (
	"PracticalProject/config"
	"PracticalProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPayments(c *gin.Context) {
	var payments []models.Payment
	if err := config.DB.Find(&payments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payments)
}

func CreatePayment(c *gin.Context) {
	var payment models.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	if err := config.DB.Create(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payment)
}

func GetPaymentByID(c *gin.Context) {
	var payment models.Payment
	id := c.Param("id")
	if err := config.DB.First(&payment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Payment not found"})
		return
	}
	c.JSON(http.StatusOK, payment)
}

func UpdatePayment(c *gin.Context) {
	var payment models.Payment
	id := c.Param("id")

	if err := config.DB.First(&payment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Payment not found"})
		return
	}

	var updatedPayment models.Payment
	if err := c.ShouldBindJSON(&updatedPayment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := config.DB.Model(&payment).Updates(updatedPayment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payment)
}

func DeletePayment(c *gin.Context) {
	var payment models.Payment
	id := c.Param("id")

	if err := config.DB.First(&payment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Payment not found"})
		return
	}

	if err := config.DB.Delete(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment deleted successfully"})
}
