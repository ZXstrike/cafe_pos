package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zxstrike/cafe-pos/order-services/db"
	"github.com/zxstrike/cafe-pos/order-services/models"
)

func GetOrders(context *gin.Context) {
	db := db.Db

	id := context.Param("id")

	var orders []models.Order

	if id != "" {
		if err := db.Preload("Items").First(&orders, "id = ?", id).Error; err != nil {
			context.JSON(400, gin.H{
				"message": "Order not found",
			})
			return
		}
	} else {
		if err := db.Preload("Items").Find(&orders).Error; err != nil {
			context.JSON(500, gin.H{
				"message": "Internal server error",
				"error":   err.Error(),
			})
			return
		}
	}

	context.JSON(200, gin.H{
		"message": "success",
		"data":    orders,
	})
}
