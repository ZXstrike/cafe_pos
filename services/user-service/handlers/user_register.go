package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zxstrike/cafe-pos/user-services/db"
	"github.com/zxstrike/cafe-pos/user-services/models"

	"golang.org/x/crypto/bcrypt"
)

func UserRegister(context *gin.Context) {

	db := db.Db

	requestBody := struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(400, gin.H{
			"error": "Invalid request body",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.DefaultCost)

	if err != nil {
		context.JSON(500, gin.H{
			"error": "Error hashing password",
		})

		return
	}

	var user models.User
	db.Where("username = ?", requestBody.Username).First(&user)

	if user.ID != 0 {
		context.JSON(400, gin.H{
			"error": "User already exists",
		})

		return
	}

	newUser := models.User{
		Username: requestBody.Username,
		Email:    requestBody.Email,
		Password: string(hash),
	}

	db.Create(&newUser)

	context.JSON(200, gin.H{
		"message": "User registration successful",
	})
}
