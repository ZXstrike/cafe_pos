package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zxstrike/cafe-pos/order-services/db"
	"github.com/zxstrike/cafe-pos/order-services/models"
)

func AddNewOrder(context *gin.Context) {

	db := db.Db

	var body struct {
		TableID uint   `json:"table_id" binding:"required"`
		Name    string `json:"name" binding:"required"`
		Note    string `json:"note"`
		Items   []struct {
			ProductID uint `json:"product_id" binding:"required"`
			Quantity  uint `json:"quantity" binding:"required"`
		} `json:"items" binding:"required"`
	}

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	order := models.Order{
		TableID: body.TableID,
		Name:    body.Name,
		Note:    body.Note,
	}

	db.Create(&order)

	for _, item := range body.Items {
		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}

		db.Create(&orderItem)
	}

	context.JSON(200, gin.H{
		"message": "success",
	})
}
