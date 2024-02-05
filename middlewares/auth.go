package middlewares

import (
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(401, gin.H{"message": "Authorization token is required."})
		return
	}

	userID, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(401, gin.H{"message": "Invalid token."})
		return
	}

	context.Set("userID", userID)
	context.Next()
}
