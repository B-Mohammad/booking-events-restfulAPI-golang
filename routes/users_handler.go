package routes

import (
	"net/http"

	"bashiri.ir/booking_events_restfulAPI_golang/models"
	"bashiri.ir/booking_events_restfulAPI_golang/utils"
	"github.com/gin-gonic/gin"
)

func signUpHandler(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse query params!"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "user can could not save!"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "new user added!"})

}

func loginHandler(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse query params!"})
		return
	}

	err = user.CheckCredential()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "user could not authenticate!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "you logged in!", "token": token})

}
