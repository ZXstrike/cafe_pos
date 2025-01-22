package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zxstrike/cafe-pos/menu-services/handlers"
)

func InitRoutes(router *gin.Engine) {
	router.POST("/menu-add", handlers.MenuAdd)
	router.GET("/menu-get-all", handlers.MenuGetAll)
	router.GET("/category-get-all", handlers.CategoryGetAll)
	router.GET("/category-get-byid", handlers.CategoryGetById)
	router.POST("/category-add", handlers.CategoryPostAdd)
}
