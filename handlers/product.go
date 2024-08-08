package handlers

import (
	"asdf/models"

	"github.com/gin-gonic/gin"
)

// CreateProduct handler
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var category models.Category
	if err := DB.First(&category, product.CategoryID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Category not found"})
		return
	}

	DB.Create(&product)
	c.JSON(201, product)
}

// GetProducts handler
func GetProducts(c *gin.Context) {
	var products []models.Product
	DB.Preload("Category").Find(&products)
	c.JSON(200, products)
}

// GetProduct handler
func GetProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := DB.Preload("Category").First(&product, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(200, product)
}

// UpdateProduct handler
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := DB.First(&product, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var category models.Category
	if err := DB.First(&category, product.CategoryID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Category not found"})
		return
	}

	DB.Save(&product)
	c.JSON(200, product)
}

// DeleteProduct handler
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(204, nil)
}
