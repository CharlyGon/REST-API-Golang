package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	server := gin.Default()

	server.GET("/events", getEvents) //GET, POST, PUT, PATCH, DELETE
	server.POST("/events", createEvent)

	server.Run(":8080") //localhost:8080
}

func getEvents(context *gin.Context) {
	events := models.GetEvents()
	context.JSON(200, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBind(&event)

	if err != nil {
		context.JSON(400, gin.H{"message": "Could not parse request. Check your data and try again."})
		return
	}

	event.ID = len(models.GetEvents()) + 1
	event.UserID = 1

	event.Save()
	context.JSON(200, gin.H{"message": "Event was created successfully.", "event": event})
}
