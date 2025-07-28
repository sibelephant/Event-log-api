package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sibelephant/prisma/db"
	"github.com/sibelephant/internal/models"
)

func (app *application) CreateEvent(c *gin.Context) {
	var event models.CreateEventRequest

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse the date string to time.Time
	parsedDate, err := time.Parse(time.RFC3339, event.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use RFC3339 format (e.g., 2023-12-25T15:04:05Z)"})
		return
	}

	createdEvent, err := app.models.Events.CreateOne(
		db.Events.Name.Set(event.Name),
		db.Events.Description.Set(event.Description),
		db.Events.Date.Set(db.DateTime(parsedDate)),
		db.Events.Location.Set(event.Location),
		db.Events.Owner.Link(
			db.Users.ID.Equals(event.OwnerID),
		),
	).Exec(context.Background())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdEvent)
}

func (app *application) GetAllEvents(c *gin.Context) {
	events, err := app.models.Events.FindMany().Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, events)
}
