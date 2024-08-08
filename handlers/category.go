package handlers

import (
	"asdf/models"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	DB.Create(&category)
	c.JSON(201, category)
}

func GetCategories(c *gin.Context) {
	var categories []models.Category
	DB.Find(&categories)
	c.JSON(200, categories)
}

func GetCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	if err := DB.First(&category, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Category not found"})
		return
	}
	c.JSON(200, category)
}

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	if err := DB.First(&category, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Category not found"})
		return
	}
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	DB.Save(&category)
	c.JSON(200, category)
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if err := DB.Delete(&models.Category{}, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Category not found"})
		return
	}
	c.JSON(204, nil)
}
