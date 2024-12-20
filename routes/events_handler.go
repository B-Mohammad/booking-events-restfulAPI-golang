package routes

import (
	"net/http"
	"strconv"

	"bashiri.ir/booking_events_restfulAPI_golang/models"
	"github.com/gin-gonic/gin"
)

func getEventsHandler(context *gin.Context) {
	allEvents, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Data events could not fetch, Please try again later!"})
		return
	}
	context.JSON(http.StatusOK, allEvents)
}

func getSingleEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event ID!"})
		return
	}

	event, err := models.GetEvent(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "events not found!"})
		return
	}
	context.JSON(http.StatusOK, event)

}

func postEventsHandler(context *gin.Context) {

	userId := context.GetInt64("userId")
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request body!"})
		return
	}

	event.UserID = userId
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Data base not available!"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event created!", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event ID!"})
		return
	}

	event, err := models.GetEvent(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "event not found!"})
		return
	}

	var uEvent models.Event
	err = context.ShouldBindJSON(&uEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request body!"})
		return
	}

	userId := context.GetInt64("userId")
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not unauthorized!"})
		return
	}

	uEvent.ID = eventId
	err = uEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not update the event!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event successfully updated!"})

}

func deleteEventHandler(context *gin.Context) {
	eId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event ID!"})
		return
	}

	event, err := models.GetEvent(eId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "event not found!"})
		return
	}

	userId := context.GetInt64("userId")
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not unauthorized!"})
		return
	}

	err = event.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not Delete Events!"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event successfully deleted!"})
}
