package main

import (
	"bashiri.ir/booking_events_restfulAPI_golang/db"
	"bashiri.ir/booking_events_restfulAPI_golang/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	db.InitialDB()

	routes.EventRoutes(server)
	server.Run(":9090")

}
