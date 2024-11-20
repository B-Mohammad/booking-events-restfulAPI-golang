package middleware

import (
	"net/http"

	"bashiri.ir/booking_events_restfulAPI_golang/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "valid token is required!"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "valid token is required!", "error": err})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
