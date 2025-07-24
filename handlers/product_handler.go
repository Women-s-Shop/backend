package handlers

import (
	"PracticalProject/config"
	"PracticalProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetProduct(c *gin.Context) {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(200, products)
}

func CreateProduct(c *gin.Context) {
	var products models.Product
	if err := c.ShouldBindJSON(&products); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := config.DB.Create(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func GetById(c *gin.Context) {
	var products models.Product
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Id not found"})
		return
	}
	if err := config.DB.First(&products, &id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
}

func UpdateProduct(c *gin.Context) {
	var products models.Product
	idParam := c.Param("id")

	if err := config.DB.First(&products, idParam).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "There is no such id"})
		return
	}

	var updateProduct models.Product
	if err := c.ShouldBindJSON(&updateProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "JSON error"})
		return
	}
	config.DB.Model(&products).Updates(updateProduct)
	c.JSON(200, products)
}

func DeleteProduct(c *gin.Context) {
	var products models.Product
	idParam := c.Param("id")

	if err := config.DB.First(&products, idParam).Error; err != nil {
		c.JSON(404, gin.H{"Error": "There is no such id"})
		return
	}
	config.DB.Delete(&products)

	c.JSON(200, gin.H{"message": "Product deleted successfully"})
}
