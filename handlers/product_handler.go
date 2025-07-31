package handlers

import (
	"PracticalProject/models"
	"PracticalProject/services"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service services.ProductService
}

func NewProductHandler(service services.ProductService) *ProductHandler {
	return &ProductHandler{service}
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	products, err := h.service.GetAllProducts()

	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(200, products)
	//var products []models.Product
	//if err := config.DB.Find(&products).Error; err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	//	return
	//}
	//c.JSON(200, products)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var products models.Product
	if err := c.ShouldBindJSON(&products); err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}
	newProducts, err := h.service.CreateProduct(products)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error creating product",
			"Error":   err.Error(),
		})
	}
	c.JSON(200, newProducts)

	//var products models.Product
	//if err := c.ShouldBindJSON(&products); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	//	return
	//}

	//if err := config.DB.Create(&products).Error; err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	//	return
	//}

	//c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetById(c *gin.Context) {
	idParam := c.Param("id")
	product, err := h.service.GetByIdProduct(idParam)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Product not found",
			"Error":   err.Error(),
		})
	}
	c.JSON(200, gin.H{"This product": product})
	//
	//var products models.Product
	//idParam := c.Param("id")
	//
	//if err := config.DB.First(&products, idParam).Error; err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	//	return
	//}
	//
	//c.JSON(http.StatusOK, gin.H{"This product": products})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	var product models.Product
	idParam := c.Param("id")
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid JSON",
			"Error":   err.Error(),
		})
	}

	updated, err := h.service.UpdateProduct(idParam, product)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Product not found",
			"Error":   err.Error(),
		})
	}
	c.JSON(200, updated)

	//if err := config.DB.First(&products, idParam).Error; err != nil {
	//	c.JSON(http.StatusNotFound, gin.H{"Error": "There is no such id"})
	//	return
	//}
	//
	//var updateProduct models.Product
	//if err := c.ShouldBindJSON(&updateProduct); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"Error": "JSON error"})
	//	return
	//}
	//config.DB.Model(&products).Updates(updateProduct)
	//c.JSON(200, products)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	if err := h.service.DeleteProduct(idParam); err != nil {
		c.JSON(404, gin.H{
			"message": "Product not found",
			"Error":   err.Error(),
		})
	}

	//if err := config.DB.First(&products, idParam).Error; err != nil {
	//	c.JSON(404, gin.H{"Error": "There is no such id"})
	//	return
	//}
	//config.DB.Delete(&products)

	c.JSON(200, gin.H{"message": "Product deleted successfully"})
}
