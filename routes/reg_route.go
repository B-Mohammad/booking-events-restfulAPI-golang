package routes

import (
	"net/http"
	"strconv"

	"bashiri.ir/booking_events_restfulAPI_golang/models"
	"github.com/gin-gonic/gin"
)

func postRegHandler(context *gin.Context) {
	eId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}

	userId := context.GetInt64("userId")

	event, err := models.GetEvent(eId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "event not available1", "err": err})
		return
	}

	err = event.RegisterToEvent(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "event not available2", "err": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "registered!"})

}

func deleteRegHandler(context *gin.Context) {
	eId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return

	}

	userId := context.GetInt64("userId")

	event, err := models.GetEvent(eId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "event not available"})
		return
	}

	err = event.DeleteRegister(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "event not available"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "erg deleted!"})
}
