package routes

import (
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"message": "Could not parse event ID."})
		return
	}

	event, err := models.GetEventById(eventID)
	if err != nil {
		context.JSON(500, gin.H{"message": "Could not fetch event."})
		return
	}

	err = event.Register(userID)
	if err != nil {
		context.JSON(500, gin.H{"message": "Could not register user for event."})
		return
	}

	context.JSON(200, gin.H{"message": "User registered for event."})
}

func cancelRegistration(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(400, gin.H{"message": "Could not parse event ID."})
		return
	}

	var event models.Event
	event.ID = eventID

	err = event.CancelRegistration(userID)
	if err != nil {
		context.JSON(500, gin.H{"message": "Could not cancel registration."})
		return
	}

	context.JSON(200, gin.H{"message": "Registration cancelled."})
}
