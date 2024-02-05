package routes

import (
	"strconv"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(500, gin.H{"message": "Could not get events."})
		return

	}
	context.JSON(200, events)
}

const invalidEventIDMessage = "Invalid event ID."

const fetchEventErrorMessage = "Could not fetch event."

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"message": invalidEventIDMessage})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(500, gin.H{"message": fetchEventErrorMessage})
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

	userId := context.GetInt64("userID")
	event.UserID = userId

	err = event.Save()
	if err != nil {
		context.JSON(500, gin.H{"message": "Could not save event."})
		return
	}

	context.JSON(200, gin.H{"message": "Event was created successfully.", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"message": invalidEventIDMessage})
		return
	}

	userID := context.GetInt64("userID")
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(500, gin.H{"message": fetchEventErrorMessage})
		return
	}

	if event.UserID != userID {
		context.JSON(403, gin.H{"message": "You are not authorized to update this event."})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(400, gin.H{"message": "Could not parse request. Check your data and try again."})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(500, gin.H{"message": "Could not update event."})
		return
	}
	context.JSON(200, gin.H{"message": "Event was updated successfully.", "event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"message": invalidEventIDMessage})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(500, gin.H{"message": fetchEventErrorMessage})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(500, gin.H{"message": "Could not delete event."})
		return
	}

	context.JSON(200, gin.H{"message": "Event was deleted successfully."})
}
