package models

import "time"

// User represents the user model for API requests/responses
type User struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// Event represents the event model for API requests/responses
type Event struct {
	ID          int       `json:"id,omitempty"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	OwnerID     int       `json:"owner_id" binding:"required"`
}

// EventResponse represents the event model for API responses with owner info
type EventResponse struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Date        time.Time    `json:"date"`
	Location    string       `json:"location"`
	OwnerID     int          `json:"owner_id"`
	Owner       UserResponse `json:"owner,omitempty"`
}

// CreateUserRequest represents the request to create a user
type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// CreateEventRequest represents the request to create an event
type CreateEventRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Date        string `json:"date" binding:"required"` // String for parsing
	Location    string `json:"location" binding:"required"`
	OwnerID     int    `json:"owner_id" binding:"required"`
}
