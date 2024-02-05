package routes

import (
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(400, gin.H{"message": "Could not parse request. Check your data and try again."})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(500, gin.H{"message": "Could not save user."})
		return
	}

	context.JSON(201, gin.H{"message": "User was created successfully.", "user": user})
}
