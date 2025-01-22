package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zxstrike/cafe-pos/menu-services/db"
	"github.com/zxstrike/cafe-pos/menu-services/models"
	"gorm.io/gorm"
)

func CategoryGetAll(context *gin.Context) {

	db := db.Db

	page, err := strconv.Atoi(context.DefaultQuery("page", "0"))

	if err != nil {
		context.JSON(400, gin.H{
			"message": "Invalid page number",
		})
		return
	}

	pageSize, err := strconv.Atoi(context.DefaultQuery("page_size", "10"))

	if err != nil {
		context.JSON(400, gin.H{
			"message": "Invalid page size",
		})
		return
	}

	offset := (page - 1) * pageSize

	var categories []models.Category

	db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(pageSize)
	}).Find(&categories)

	context.JSON(200, gin.H{
		"message": "success",
		"data":    categories,
	})

}
