package routes

import "github.com/gin-gonic/gin"

func RegistetRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvents)
	server.PUT("/events/:id", updateEvents)
	server.DELETE("/events/:id", deleteEvent)
}
