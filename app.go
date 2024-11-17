package main

import (
	"net/http"

	"bashiri.ir/booking_events_restfulAPI_golang/db"
	"bashiri.ir/booking_events_restfulAPI_golang/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	db.InitialDB()

	server.GET("/events", getEventsHandler)
	server.POST("/events", postEventsHandler)

	server.Run(":9090")

}

func getEventsHandler(context *gin.Context) {
	allEvents := models.GetAllEvents()
	context.JSON(http.StatusOK, allEvents)
}

func postEventsHandler(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request body!"})
		return
	}

	event.ID = 1
	event.UserID = 1
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "event created!", "event": event})
}
