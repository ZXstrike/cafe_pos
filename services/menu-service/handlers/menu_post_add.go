package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zxstrike/cafe-pos/menu-services/db"
	"github.com/zxstrike/cafe-pos/menu-services/models"
)

func MenuAdd(context *gin.Context) {
	// Use the existing DB instance
	db := db.Db

	// Define request body struct
	var Body struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description" binding:"required"`
		Price       int    `json:"price" binding:"required"`
		Quantity    int    `json:"quantity" binding:"required"`
		CategoryID  uint   `json:"category_id" binding:"required"` // Changed to uint for consistency
	}

	// Bind JSON input to the Body struct
	if err := context.BindJSON(&Body); err != nil {
		context.JSON(400, gin.H{
			"message": "Invalid request: " + err.Error(),
		})
		return
	}

	// Check if the category exists
	var category models.Category
	if err := db.First(&category, "id = ?", Body.CategoryID).Error; err != nil {
		context.JSON(400, gin.H{
			"message": "Category not found or invalid CategoryID",
		})
		return
	}

	// Create the Menu object
	menu := models.Menu{
		Name:        Body.Name,
		Description: Body.Description,
		Price:       Body.Price,
		Quantity:    Body.Quantity,
		CategoryID:  Body.CategoryID,
	}

	// Save the Menu to the database
	if err := db.Create(&menu).Error; err != nil {
		context.JSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}

	// Return success response
	context.JSON(201, gin.H{
		"message": "Menu added successfully",
		"data":    menu,
	})
}
