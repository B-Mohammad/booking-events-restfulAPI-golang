package main

import (
	"net/http"

	"bashiri.ir/booking_events_restfulAPI_golang/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEventsHandler)

	server.Run(":9090")

}

func getEventsHandler(context *gin.Context) {
	allEvents := models.GetAllEvents()
	context.JSON(http.StatusOK, allEvents)
}
