package routes

import "github.com/gin-gonic/gin"

const eventIDPath = "/events/:id"

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents) //GET, POST, PUT, PATCH, DELETE
	server.GET(eventIDPath, getEvent)
	server.POST("/events", createEvent)
	server.PUT(eventIDPath, updateEvent)
	server.DELETE(eventIDPath, deleteEvent)
}
