package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

const eventIDPath = "/events/:id"

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents) //GET, POST, PUT, PATCH, DELETE
	server.GET(eventIDPath, getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT(eventIDPath, updateEvent)
	authenticated.DELETE(eventIDPath, deleteEvent)
	authenticated.POST("events/:id/register", registerForEvent)
	authenticated.DELETE("events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
