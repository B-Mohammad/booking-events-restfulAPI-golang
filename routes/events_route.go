package routes

import (
	"bashiri.ir/booking_events_restfulAPI_golang/middleware"
	"github.com/gin-gonic/gin"
)

func EventRoutes(server *gin.Engine) {

	auth := server.Group("/")
	auth.Use(middleware.Authenticate)

	auth.DELETE("/events/:id", deleteEventHandler)
	auth.PUT("/events/:id", updateEvent)
	auth.POST("/events", postEventsHandler)

	auth.POST("events/:id/register", postRegHandler)
	auth.DELETE("events/:id/register", deleteRegHandler)

	server.GET("/events", getEventsHandler)
	server.GET("/events/:id", getSingleEvent)

	server.POST("/signup", signUpHandler)
	server.POST("/login", loginHandler)
}
