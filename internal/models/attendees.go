package models

// CreateAttendeeRequest represents the request to add a user to an event
type CreateAttendeeRequest struct {
	UserID  int `json:"user_id" binding:"required"`
	EventID int `json:"event_id" binding:"required"`
}

// AttendeeResponse represents the attendee model for API responses
type AttendeeResponse struct {
	ID      int          `json:"id"`
	UserID  int          `json:"user_id"`
	EventID int          `json:"event_id"`
	User    UserResponse `json:"user,omitempty"`
}

// EventWithAttendeesResponse represents an event with its attendees
type EventWithAttendeesResponse struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Date        string             `json:"date"`
	Location    string             `json:"location"`
	OwnerID     int                `json:"owner_id"`
	Owner       UserResponse       `json:"owner,omitempty"`
	Attendees   []AttendeeResponse `json:"attendees"`
}

// UserWithEventsResponse represents a user with events they're attending
type UserWithEventsResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Events []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Date        string `json:"date"`
		Location    string `json:"location"`
	} `json:"attending_events"`
}
