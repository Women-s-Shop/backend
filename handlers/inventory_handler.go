package handlers

import (
	"PracticalProject/config"
	"PracticalProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetInventories(c *gin.Context) {
	var inventories []models.Inventory
	if err := config.DB.Find(&inventories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, inventories)
}

func CreateInventory(c *gin.Context) {
	var inventory models.Inventory
	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	if err := config.DB.Create(&inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, inventory)
}

func GetInventoryByID(c *gin.Context) {
	var inventory models.Inventory
	id := c.Param("id")
	if err := config.DB.First(&inventory, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Inventory record not found"})
		return
	}
	c.JSON(http.StatusOK, inventory)
}

func UpdateInventory(c *gin.Context) {
	var inventory models.Inventory
	id := c.Param("id")

	if err := config.DB.First(&inventory, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Inventory record not found"})
		return
	}

	var updatedInventory models.Inventory
	if err := c.ShouldBindJSON(&updatedInventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := config.DB.Model(&inventory).Updates(updatedInventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, inventory)
}

func DeleteInventory(c *gin.Context) {
	var inventory models.Inventory
	id := c.Param("id")

	if err := config.DB.First(&inventory, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Inventory record not found"})
		return
	}

	if err := config.DB.Delete(&inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inventory record deleted successfully"})
}
