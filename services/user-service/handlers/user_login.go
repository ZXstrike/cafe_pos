package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zxstrike/cafe-pos/user-services/db"
	"github.com/zxstrike/cafe-pos/user-services/models"
	"github.com/zxstrike/cafe-pos/user-services/utils/jwt"
	"golang.org/x/crypto/bcrypt"
)

func UserLogin(context *gin.Context) {

	db := db.Db

	requestBody := struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(400, gin.H{
			"error": "Invalid request body",
		})

		return
	}

	var user models.User

	db.Where("username = ?", requestBody.Username).First(&user)

	if user.ID == 0 {
		context.JSON(400, gin.H{
			"error": "User not found",
		})

		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password)); err != nil {
		context.JSON(400, gin.H{
			"error": "Invalid password",
		})

		return
	}

	access_token, err := jwt.GenerateAccessToken(user)

	if err != nil {
		context.JSON(500, gin.H{
			"error": "Error generating access token",
		})

		return
	}

	context.JSON(200, gin.H{
		"message":      "User login successful",
		"access_token": access_token,
	})
}
