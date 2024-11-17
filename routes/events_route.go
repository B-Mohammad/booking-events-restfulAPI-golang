package routes

import "github.com/gin-gonic/gin"

func EventRoutes(server *gin.Engine) {

	server.GET("/events", getEventsHandler)
	server.GET("/events/:id", getSingleEvent)
	server.POST("/events", postEventsHandler)
	server.PUT("/events/:id", UpdateEvent)
}
