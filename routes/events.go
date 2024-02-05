package routes

import (
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(500, gin.H{"message": "Could not get events."})
		return

	}
	context.JSON(200, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"message": "Invalid event ID."})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(500, gin.H{"message": "Could not get event."})
		return
	}

	context.JSON(200, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(400, gin.H{"message": "Could not parse request. Check your data and try again."})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(500, gin.H{"message": "Could not save event."})
		return
	}

	context.JSON(200, gin.H{"message": "Event was created successfully.", "event": event})
}
