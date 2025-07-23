package handlers

import (
	"PracticalProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var products = []models.Product{
	{Id: 1, Name: "Apple watch", Description: "its liked bisnes Women", Price: 185000, ImageUrl: "56", CategoryId: 1},
	{Id: 2, Name: "Smart watch", Description: "its liked Women", Price: 125000, ImageUrl: "6", CategoryId: 1},
}

func GetProduct(c *gin.Context) {
	c.JSON(200, products)
}

func CreateProduct(c *gin.Context) {
	var newProducts models.Product

	if err := c.ShouldBindJSON(&newProducts); err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}
	newProducts.Id = len(products) + 1
	products = append(products, newProducts)
	c.JSON(200, newProducts)
}

func GetById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Id not found"})
		return
	}

	for _, p := range products {
		if p.Id == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
}

func UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	var updateProduct models.Product
	if err := c.ShouldBindJSON(&updateProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "JSON error"})
		return
	}

	for i, p := range products {
		if p.Id == id {
			updateProduct.Id = id
			products[i] = updateProduct
			c.JSON(http.StatusOK, updateProduct)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"Error": "Not found"})
}

func DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Id error"})
		return
	}
	for i, p := range products {
		if p.Id == id {
			products = append(products[:i], products[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"massage": "Successfully deleted"})
			return
		}

	}
	c.JSON(404, gin.H{"Error": "Not found"})
}
