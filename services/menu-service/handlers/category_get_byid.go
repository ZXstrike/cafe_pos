package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zxstrike/cafe-pos/menu-services/db"
	"github.com/zxstrike/cafe-pos/menu-services/models"
)

func CategoryGetById(context *gin.Context) {

	db := db.Db

	category := models.Category{}

	db.First(&category, context.Query("id"))

	context.JSON(200, gin.H{
		"message": "Get category by ID",
		"data":    category,
	})
}
