package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sibelephant/internal/models"
	"github.com/sibelephant/prisma/db"
)

// CreateAttendee adds a user to an event (user registers for an event)
func (app *application) CreateAttendee(c *gin.Context) {
	var req models.CreateAttendeeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user exists
	_, err := app.models.Users.FindUnique(
		db.Users.ID.Equals(req.UserID),
	).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// Check if event exists
	_, err = app.models.Events.FindUnique(
		db.Events.ID.Equals(req.EventID),
	).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Event not found"})
		return
	}

	// Check if user is already attending this event
	existingAttendee, _ := app.models.Attendees.FindFirst(
		db.Attendees.UserID.Equals(req.UserID),
		db.Attendees.EventID.Equals(req.EventID),
	).Exec(context.Background())
	if existingAttendee != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User is already attending this event"})
		return
	}

	// Create the attendee record
	createdAttendee, err := app.models.Attendees.CreateOne(
		db.Attendees.User.Link(
			db.Users.ID.Equals(req.UserID),
		),
		db.Attendees.Event.Link(
			db.Events.ID.Equals(req.EventID),
		),
	).Exec(context.Background())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully registered for event",
		"attendee": gin.H{
			"id":       createdAttendee.ID,
			"user_id":  createdAttendee.UserID,
			"event_id": createdAttendee.EventID,
		},
	})
}

// GetAllAttendees retrieves all attendee records
func (app *application) GetAllAttendees(c *gin.Context) {
	attendees, err := app.models.Attendees.FindMany().With(
		db.Attendees.User.Fetch(),
		db.Attendees.Event.Fetch(),
	).Exec(context.Background())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert to response format
	var response []models.AttendeeResponse
	for _, attendee := range attendees {
		response = append(response, models.AttendeeResponse{
			ID:      attendee.ID,
			UserID:  attendee.UserID,
			EventID: attendee.EventID,
			User: models.UserResponse{
				ID:    attendee.User().ID,
				Name:  attendee.User().Name,
				Email: attendee.User().Email,
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{"attendees": response})
}

// GetEventAttendees retrieves all attendees for a specific event
func (app *application) GetEventAttendees(c *gin.Context) {
	eventIDStr := c.Param("event_id")
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	// Check if event exists
	event, err := app.models.Events.FindUnique(
		db.Events.ID.Equals(eventID),
	).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	// Get attendees for this event
	attendees, err := app.models.Attendees.FindMany(
		db.Attendees.EventID.Equals(eventID),
	).With(
		db.Attendees.User.Fetch(),
	).Exec(context.Background())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert to response format
	var attendeesList []models.AttendeeResponse
	for _, attendee := range attendees {
		attendeesList = append(attendeesList, models.AttendeeResponse{
			ID:      attendee.ID,
			UserID:  attendee.UserID,
			EventID: attendee.EventID,
			User: models.UserResponse{
				ID:    attendee.User().ID,
				Name:  attendee.User().Name,
				Email: attendee.User().Email,
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"event": gin.H{
			"id":          event.ID,
			"name":        event.Name,
			"description": event.Description,
			"date":        event.Date,
			"location":    event.Location,
		},
		"attendees_count": len(attendeesList),
		"attendees":       attendeesList,
	})
}

// GetUserEvents retrieves all events a user is attending
func (app *application) GetUserEvents(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Check if user exists
	user, err := app.models.Users.FindUnique(
		db.Users.ID.Equals(userID),
	).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Get events the user is attending
	attendees, err := app.models.Attendees.FindMany(
		db.Attendees.UserID.Equals(userID),
	).With(
		db.Attendees.Event.Fetch(),
	).Exec(context.Background())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert to response format
	var eventsList []gin.H
	for _, attendee := range attendees {
		event := attendee.Event()
		eventsList = append(eventsList, gin.H{
			"id":          event.ID,
			"name":        event.Name,
			"description": event.Description,
			"date":        event.Date,
			"location":    event.Location,
			"owner_id":    event.OwnerID,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
		"events_count":     len(eventsList),
		"attending_events": eventsList,
	})
}

// DeleteAttendee removes a user from an event (unregister)
func (app *application) DeleteAttendee(c *gin.Context) {
	attendeeIDStr := c.Param("id")
	attendeeID, err := strconv.Atoi(attendeeIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid attendee ID"})
		return
	}

	// Check if attendee record exists
	attendee, err := app.models.Attendees.FindUnique(
		db.Attendees.ID.Equals(attendeeID),
	).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Attendee record not found"})
		return
	}

	// Delete the attendee record
	_, err = app.models.Attendees.FindUnique(
		db.Attendees.ID.Equals(attendeeID),
	).Delete().Exec(context.Background())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Successfully unregistered from event",
		"user_id":  attendee.UserID,
		"event_id": attendee.EventID,
	})
}

// DeleteUserFromEvent removes a specific user from a specific event
func (app *application) DeleteUserFromEvent(c *gin.Context) {
	userIDStr := c.Param("user_id")
	eventIDStr := c.Param("event_id")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	// Find the attendee record
	attendee, err := app.models.Attendees.FindFirst(
		db.Attendees.UserID.Equals(userID),
		db.Attendees.EventID.Equals(eventID),
	).Exec(context.Background())

	if err != nil || attendee == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User is not attending this event"})
		return
	}

	// Delete the attendee record
	_, err = app.models.Attendees.FindUnique(
		db.Attendees.ID.Equals(attendee.ID),
	).Delete().Exec(context.Background())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Successfully removed user from event",
		"user_id":  userID,
		"event_id": eventID,
	})
}
