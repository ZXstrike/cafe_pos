package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zxstrike/cafe-pos/menu-services/db"
	"github.com/zxstrike/cafe-pos/menu-services/models"
)

func CategoryPostAdd(context *gin.Context) {

	db := db.Db

	var category models.Category

	if err := context.BindJSON(&category); err != nil {
		context.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	db.Create(&category)

	context.JSON(200, gin.H{
		"message": "success",
	})

}
