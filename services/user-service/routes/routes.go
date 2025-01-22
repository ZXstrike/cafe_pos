package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zxstrike/cafe-pos/user-services/handlers"
)

func InitRoutes(router *gin.Engine) {
	router.POST("/register", handlers.UserRegister)
	router.POST("/login", handlers.UserLogin)

}
