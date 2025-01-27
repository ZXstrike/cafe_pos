package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zxstrike/cafe-pos/order-services/handlers"
)

func InitRoutes(router *gin.Engine) {
	router.POST("/order", handlers.AddNewOrder)
	router.GET("/order", handlers.GetOrders)
	// router.GET("/order/:id", handlers.GetOrder)
	// router.PUT("/order/:id", handlers.UpdateOrder)
	// router.DELETE("/order/:id", handlers.DeleteOrder)
	// router.GET("/order/:id/items", handlers.GetOrderItems)
	// router.POST("/order/:id/items", handlers.AddOrderItem)
	// router.DELETE("/order/:id/items/:item_id", handlers.DeleteOrderItem)
}
