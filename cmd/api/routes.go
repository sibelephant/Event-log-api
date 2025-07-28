package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {
	g := gin.Default()

	v1 := g.Group("api/v1")
	{
		// User routes
		v1.POST("/users", app.CreateUser)
		v1.GET("/users", app.GetAllUsers)
		v1.GET("/users/:id", app.GetUser)

		// Event routes
		v1.POST("/events", app.CreateEvent)
		v1.GET("/events", app.GetAllEvents)
		// TODO: Implement these methods
		// v1.GET("/events/:id", app.getEvent)
		// v1.PUT("/events/:id", app.updateEvent)
		// v1.DELETE("/events/:id", app.deleteEvent)

		// Attendees routes
		v1.POST("/attendees", app.CreateAttendee)
		v1.GET("/attendees", app.GetAllAttendees)
		v1.DELETE("/attendees/:id", app.DeleteAttendee)

		// Event-specific attendee routes
		v1.GET("/events/:event_id/attendees", app.GetEventAttendees)

		// User-specific event routes
		v1.GET("/user/:user_id/attending-events", app.GetUserEvents)
		v1.DELETE("/user/:user_id/events/:event_id", app.DeleteUserFromEvent)
	}

	return g
}
