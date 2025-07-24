package handlers

import (
	"PracticalProject/config"
	"PracticalProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPromocodes(c *gin.Context) {
	var promocodes []models.Promocode
	if err := config.DB.Find(&promocodes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, promocodes)
}

func CreatePromocode(c *gin.Context) {
	var promocode models.Promocode
	if err := c.ShouldBindJSON(&promocode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	if err := config.DB.Create(&promocode).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, promocode)
}

func GetPromocodeByID(c *gin.Context) {
	var promocode models.Promocode
	id := c.Param("id")
	if err := config.DB.First(&promocode, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Promocode not found"})
		return
	}
	c.JSON(http.StatusOK, promocode)
}

func UpdatePromocode(c *gin.Context) {
	var promocode models.Promocode
	id := c.Param("id")

	if err := config.DB.First(&promocode, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Promocode not found"})
		return
	}

	var updatedPromocode models.Promocode
	if err := c.ShouldBindJSON(&updatedPromocode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := config.DB.Model(&promocode).Updates(updatedPromocode).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, promocode)
}

func DeletePromocode(c *gin.Context) {
	var promocode models.Promocode
	id := c.Param("id")

	if err := config.DB.First(&promocode, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Promocode not found"})
		return
	}

	if err := config.DB.Delete(&promocode).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Promocode deleted successfully"})
}
